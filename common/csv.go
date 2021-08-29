package common

type CsvRow interface {
	Row() []string
	Headers() []string
}

type CsvTable interface {
	Rows() []CsvRow
	Headers() []string
}
