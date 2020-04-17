# go-hub
`go-hub` is an attempt at providing a simple, easy to use and expandable websocket client-server communication bus.

[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/devodev/go-hub?sort=semver)](https://github.com/devodev/go-hub/tags)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/devodev/go-hub)](https://github.com/golang/go/wiki/Modules)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/mod/github.com/devodev/go-hub)
[![Go Report Card](https://goreportcard.com/badge/github.com/devodev/go-hub)](https://goreportcard.com/report/github.com/devodev/go-hub)
[![GitHub license](https://img.shields.io/github/license/devodev/go-hub?style=flat)](https://github.com/devodev/go-hub/blob/master/LICENSE.txt)

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Overview
Currently, **`go-hub` requires Go version 1.13 or greater**.

## Installation
```
go get github.com/devodev/go-hub
```

## Development
You can use the provided `Dockerfile` to build an image that will provide a clean environment for development purposes.</br>
Instructions that follow assumes you are running `Windows`, have `Docker Desktop` installed and its daemon is running.

Clone this repository and build the image
```
$ git clone https://github.com/devodev/go-hub
$ cd ./go-hub
$ docker build --tag=go-hub .
```

Run a container using the previously built image while mounting the CWD and exposing port 8080
```
$ docker run \
    --rm \
    --volume="$(pwd -W):/srv/src/github.com/devodev/go-hub" \
    --tty \
    --interactive \
    -p 8080:8080 \
    go-hub
$ root@03e67598a37f:/srv/src/github.com/devodev/go-hub#
```

Start deving
```
$ go run .
go: downloading github.com/gorilla/websocket v1.4.2
go: extracting github.com/gorilla/websocket v1.4.2
go: finding github.com/gorilla/websocket v1.4.2
2020/04/17 18:10:34 ListenAndServe: ":8080"
```

## Contributing
This repository is under heavy development and is subject to change in the near future.</br>
Versioning will be locked and a proper contributing section will be created in a timely manner, when code is stabilized.</br>

## License
`go-hub` is released under the MIT license. See [LICENSE.txt](LICENSE.txt)
