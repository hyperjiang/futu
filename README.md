# Futu open api golang client

[![GoDoc](https://pkg.go.dev/badge/github.com/hyperjiang/futu)](https://pkg.go.dev/github.com/hyperjiang/futu)
[![CI](https://github.com/hyperjiang/futu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/hyperjiang/futu/actions/workflows/ci.yml)
[![](https://goreportcard.com/badge/github.com/hyperjiang/futu)](https://goreportcard.com/report/github.com/hyperjiang/futu)
[![codecov](https://codecov.io/gh/hyperjiang/futu/branch/main/graph/badge.svg)](https://codecov.io/gh/hyperjiang/futu)
[![Release](https://img.shields.io/github/release/hyperjiang/futu.svg)](https://github.com/hyperjiang/futu/releases)

Futu open api golang client. Require go version >= 1.21.

API doc: https://openapi.futunn.com/futu-api-doc/

## Usage

```bash
go get -u github.com/hyperjiang/futu
```

```go
import "github.com/hyperjiang/futu"

client, err := futu.NewClient()
if err != nil {
    log.Fatal(err)
}

res, err := client.GetGlobalState()
fmt.Println(res)
```
