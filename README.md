# resolvy

<div align="center">
  <a href="https://golang.org/">
    <img
      src="https://img.shields.io/badge/MADE%20WITH-GO-%23EF4041?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://pkg.go.dev/github.com/daspoet/resolvy">
    <img
      src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://goreportcard.com/report/github.com/daspoet/gowinkey">
    <img
      src="https://goreportcard.com/badge/github.com/daspoet/resolvy?style=for-the-badge"
      height="30"
    />
  </a>
</div>

## Contents

- [resolvy](#resolvy)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Getting started](#getting-started)

## Installation

To use `resolvy`, you need to have [Go](https://golang.org/) installed and set up.

Now, you can get `resolvy` by running

```shell
$ go get -u github.com/daspoet/resolvy
```

and import it in your code:

```go
import "github.com/daspoet/resolvy"
```

## Getting started

To understand how to register custom resolvers on a type with `resolvy`, let's look at the following example:

Suppose we have a struct

```go
import "time"

type Message struct {
    Timestamp time.Time `json:"timestamp,omitempty"`
    Author    string    `json:"author,omitempty"`
}
```

that we would like to marshal into JSON while also formatting its `Timestamp` field. While the standard library allows us to make Message implement the `json.Marshaler` interface

```go
import (
    "encoding/json",
    "time"
)

type Message struct {
    Timestamp time.Time
    Author    string    
}

func (msg Message) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct {
        Timestamp string `json:"timestamp,omitempty"`
        Author    string `json:"author,omitempty"`
    }{
        Timestamp: msg.Timestamp.Format(time.RFC850),
        Author:    msg.Author,
    })
}
```

we can immediately see that this is problematic, because we have to list all the fields we *don't* want to alter the marshalled representation of alongside the `Timestamp` field. This obscures our original intent to anyone trying to understand our code.

To solve this problem, `resolvy` allows us to register custom marshalers on specific fields of a struct type. Using `resolvy`, we can simplify our code:

```go
import (
    "encoding/json",
    "github.com/daspoet/resolvy",
    "time"
)

type Message struct {
    Timestamp time.Time `resolvy:"timestamp,omitempty"`
    Author    string    `resolvy:"author,omitempty"`
}

func (msg Message) MarshalJSON() ([]byte, error) {
    return resolvy.MarshalJSON(msg, MarshalConfig{
        Marshalers: map[string]FieldMarshaler{
            "timestamp": func() (interface{}, error) {
                return msg.Timestamp.Format(time.RFC850), nil
            },
        },
    })
}
```
