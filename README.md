# requests

Helper library for interacting with https and http endpoints with a python requests similarity

# Quickstart Instructions

Install the library
```go
go get github.com/dougkirkley/requests
```

# Usage

Without https

```go
package main

import (
    "log"
    "github.com/dougkirkley/requests"
    "github.com/dougkirkley/requests/https"
)

func main() {
    client := requests.NewClient(&https.CertInput{})
    resp, err := client.Get("https://api.helium.io/v1/blocks/height")
    if err != nil {
        log.Println(err)
    }
    log.Println(resp)
}
```

With https

```go
package main

import (
    "log"
    "github.com/dougkirkley/requests"
    "github.com/dougkirkley/requests/https"
)

func main() {
    client := requests.NewClient(&https.CertInput{
        CaCert: "cacerts.pem",
        KeyFile: "key.pem",
        CertFile: "cert.pem",
    })
    resp, err := client.Get("https://api.helium.io/v1/blocks/height")
    if err != nil {
        log.Println(err)
    }
    log.Println(resp)
}
```