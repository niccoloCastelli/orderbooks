# Orderbooks

__Crypto order books scraper__

Keep a limit order book copy updated in real time for various crypto pairs

__Supported exchanges__
- [x] bitmex _(perpetual)_
- [x] bitstamp _(spot)_
- [x] coinbase pro _(spot)_
- [x] kraken _(spot)_
- [x] okex _(spot)_


## Installation

### From binary
1. Download the binary
2. Run `chmod +x orderbooks`
3. Copy the file in a _$PATH_ directory (eg. `/usr/bin`)

### From source
1. Clone the repo
   
   `git clone ...`

2. Install dependencies
   
   `go get ./...`

3. Customize configuration
   
   Copy `config.json` and edit it according to your needs:
   - `pairs`: pairs to scrape (if an exchange doesn't support a pair it is ignored for that specific exchange with a warning on startup)
   - `snapshot_interval`: interval between full snapshots (in seconds)
   - `storage_path`: where snapshot and events file are saved
   - `exchanges`: exchanges to scrape 
   - `save_Events`: save events files

4. _(optional)_ Install systemd service (tested on Ubuntu 18.04)
   ```bash
   go build -o scripts/orderbooks
   ./scripts/install.sh
   ```
   ⚠️ **Systemd service is enabled on startup by default**
   


## Usage

### Help
Run `orderbooks --help` for command list and `orderbooks [COMMAND] --help` for detailed command help

### Run scraper
Run order books scraper and keep save data in csv files.

`orderbooks run --log_level 1 --config your_config_file.json`

⚠️**Uncompressed files for all the exchanges and supported pairs may take more than 5GB a day, use the builtin `compress` command to save 80% space**


#### Storage format

- `[FILE_TYPE]` (events|snapshots)
  - exchange
    - year
      - month
        - day
          - `[PAIR].csv`

### Compress files

Compress stored files up to today-n day (default n = 5)

`orderbooks compress [n]`

### Exctact files
Exctact compressed files for `[EXCHANGE]` and `[PAIR]` in `[TARGET PATH]`

`orderbooks extract --extract_events=[0|1] [EXCHANGE] [PAIR] [TARGET PATH]`

Optional parameters:
- `--extract_events` _(default 0)_ Exctract snapshots files if 0 or event files if 1

### Remove stored files
Remove old stored files up to today-n day (default n = 5)

`orderbooks clean [n]`


### Upload to Elasticsearch
Upload data stored in files to elasticsearch

`orderbooks to_elasticsearch --storage_path [STORAGE_PATH] --interval=10 --snapshot_mode=1 [PAIR] [EXCHANGE] [FROM_DATE] [TO_DATE]`

All positional arguments are required.

Dates are in `yyyy-mm-dd` format

Optional parameters:
- `--storage_path` _(default 0)_ 
- `--interval`
- `--snapshot_mode`

### Live feed

Publish real time snapshots on a NATS channel and optionally save snapshots to elasticsearch

`data_feed --interval 1 [PAIR] [EXCHANGE]`

### gRPC server

Run a grpc server to query stored files and/or stream live data (check `server/proto/orderbooks.proto` for available endpoints and API docs)

`orderbooks server --storage_path [STORAGE_PATH]`


## Contributing

The project was meant as a learning project, so it is not regularly maintained but any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
