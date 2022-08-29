# goi
Go joi port - validation package

[![Status](https://github.com/geek/goi/actions/workflows/go.yml/badge.svg)](https://github.com/geek/goi/actions/workflows/go.yml)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg?style=flat-square)

## Install

```shell
go get github.com/geek/goi
```

## Usage Example

```go
s := goi.String("myField").Required().Valid("foo", "bar")
if err := s.Validate(str); err != nil {
    return err
}

n := goi.Number[float64]("test").Invalid(1.1, 1.2).Min(.5).Max(3.0)
if err := n.Validate(num); err != nil {
    return err
}
```

## API

[API](./api.md)