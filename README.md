# TOKENFEW

⚡ Fewer tokens. Smarter complexity. ⚡

## License

[![License:Apache2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/Apache2.0)

## Basic Requirements

Your device should meet the following basic requirements.

```shell
Distributor ID: Ubuntu
Description:    Ubuntu 24.04.4 LTS
Release:        24.04
Codename:       noble
```

```shell
Golang Version: 1.24.5+
```

## Development

```shell
go mod download
```

```shell
go run ./cmd/main.go
```

## Build

```shell
export GO111MODULE=on
export GOPROXY=https://goproxy.io,direct
go mod tidy && cd ./cmd
go build -o ../release/realtime main.go
```

## Team

TokenFew https://www.tokenfew.com
