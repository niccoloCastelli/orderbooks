package utils

import (
	"compress/gzip"
	"github.com/niccoloCastelli/orderbooks/common/constants"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func GunzipWrite(w io.Writer, r io.Reader) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gr.Close()
	data, err := ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func ExtractFile(fs afero.Fs, absPath string, targetFs afero.Fs, targetPath string, exchangeName string, fileType constants.FileType, refTime time.Time) (string, error) {
	filePath, name := filepath.Split(absPath)
	ext := filepath.Ext(name)
	if targetPath != "" {
		name = strings.Join([]string{string(fileType), exchangeName, refTime.Format("2006-01-02"), name}, "_")
	} else {
		targetPath = filePath
	}
	newName := strings.TrimSuffix(name, ext)
	targetFile := path.Join(targetPath, newName)
	_, err := os.Stat(targetFile)
	if !os.IsNotExist(err) {
		return targetFile, nil
	}
	//path.Join(targetPath,newName)

	//data, err := afero.ReadFile(fs, absPath)
	fReader, err := fs.Open(absPath)
	if err != nil {
		return "", err
	}

	file, err := targetFs.Create(targetFile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return targetFile, GunzipWrite(file, fReader)
}
