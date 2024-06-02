# envar
[![test](https://github.com/goccha/envar/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/goccha/envar/actions/workflows/test.yml)
A simple library that handles environment variables

## 使い方

```go
package example

var host = envar.String("DB_HOST")
var port = envar.Int("DB_PORT")
var debug = envar.Bool("DEBUG")
```

## default value

```go
var host = envar.Get("DB_HOST").String("localhost")
var port = envar.Get("DB_PORT").Int(3306)
var debug = envar.Get("DEBUG").Bool(true)
```

## multiple environment variables
Specify multiple environment variables and adopt the one with the value set

```go
var primaryHost = envar.Get("PRIMARY_HOST", "DB_HOST").String("localhost")
var primaryPort = envar.Get("PRIMARY_PORT", "DB_PORT").Int(3306)
var replicaHost = envar.Get("REPLICA_HOST", "DB_HOST").String("localhost")
var replicaPort = envar.Get("REPLICA_PORT", "DB_PORT").Int(3306)
```

## Struct Tag Support

```go
type Example struct {
	Host string `envar:"PRIMARY_HOST,DB_HOST;default=localhost"`
	Port int `envar:"PRIMARY_PORT,DB_PORT;default=3306"`
	Debug bool `envar:"DEBUG"`
}
obj := &Example{}
envar.Bind(&obj)
```

### Common prefix environment variable name

```go
type Example struct {
    Host string `envar:"default=localhost"`
    Port int `envar:"default=3306"`
    Debug bool
}
obj := &Example{}
envar.Bind(&obj, envar.WithPrefix("APP_"))
```