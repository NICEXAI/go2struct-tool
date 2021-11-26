# go2struct-tool
Use the command to convert arbitrary formats to Go Struct (including json, toml, yaml, etc.)

### Installation

Run the following command under your project:

> go get -u github.com/NICEXAI/go2struct-tool

### Basic Usage

#### convert a yml file to a go struct

> go2struct-tool -i ./setting_default.yml -o config.go

### auto-conversion
If you want to enable automatic conversion, you can use the `-w` flag.

> go2struct-tool -w -i ./setting_default.yml -o config.go
