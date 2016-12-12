# gohttp

[![Golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![Build Status](https://travis-ci.org/philchia/gohttp.svg?branch=master)](https://travis-ci.org/philchia/gohttp)
[![Coverage Status](https://coveralls.io/repos/github/philchia/gohttp/badge.svg?branch=master)](https://coveralls.io/github/philchia/gohttp?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/philchia/gohttp)](https://goreportcard.com/report/github.com/philchia/gohttp)
[![GoDoc](https://godoc.org/github.com/philchia/gohttp?status.svg)](https://godoc.org/github.com/philchia/gohttp)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://opensource.org/licenses/MIT)

gohttp is a lightweight and elegant http network library for Go.

## Usage

### Get

```go
import  "github.com/philchia/gohttp"

gohttp.Request(Get, "https://www.google.com").
           ResponseString(
               func(code int, header http.Header, body string, err error) {
                   log.Println(code)
                   log.Println(header)
                   log.Println(body)
            })
```

### Post

```go
import  "github.com/philchia/gohttp"

parameters := map[string]interface{}{
    "page": 1,
    "size": 20,
    }

gohttp.Request(Post, "https://www.google.com", parameters).
           ResponseString(
               func(code int, header http.Header, body string, err error) {
                   log.Println(code)
                   log.Println(header)
                   log.Println(body)
            })
```

## Todo

* Add header for every single reuqest
* Error retry handler
* Redirect handler
* Body encode customize
* Async response handler
* Basic auth support
* Mock for test

## License

gohttp code published under MIT license