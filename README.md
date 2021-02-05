# MrAndreID / Go Null

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/gonull.svg)](https://pkg.go.dev/github.com/MrAndreID/gonull) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoNull` package is a collection of functions in the go language for handle null when json encode and json decode with `database/sql` package.

---

## Table of Contents

* [Install](#install)
* [Usage](#usage)
* [Versioning](#versioning)
* [Authors](#authors)
* [License](#license)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [More](#more)

---

## Install

To use The `MrAndreID/GoNull` package, you must follow the steps below:

```sh
go get -u github.com/MrAndreID/gonull
```

## Usage

To use The `MrAndreID/GoNull` package, you must combine it with The `database/sql` package.

### Null String

```go
import "database/sql"

name := gonull.NullString{sql.NullString{"Andrea Adam", true}}
```

### Null Int32

```go
import "database/sql"

row := gonull.NullInt32{sql.NullInt32{1, true}}
```

### Null Int64

```go
import "database/sql"

row := gonull.NullInt64{sql.NullInt64{int64(1), true}}
```

### Null Bool

```go
import "database/sql"

active := gonull.NullBool{sql.NullBool{true, true}}
```

### Null Time

```go
import (
    "time"
	"database/sql"
)

date, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
createdAt = gonull.NullTime{sql.NullTime{date, true}}
```

## Versioning

I use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## License

MIT licensed. See the LICENSE file for details.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Language website](https://golang.org/doc/).

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/gonull).
