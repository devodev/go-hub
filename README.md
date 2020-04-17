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

Run a container using the previously built image while mounting the CWD
```
$ docker run \
    --rm \
    --volume="$(pwd -W):/srv/src/github.com/devodev/go-hub" \
    --tty \
    --interactive \
    go-hub
$ root@03e67598a37f:/srv/src/github.com/devodev/go-hub# ll
total 4
drwxrwxrwx 1 root root 4096 Apr 17 16:31 ./
drwxr-xr-x 1 root root 4096 Apr 17 16:36 ../
-rwxr-xr-x 1 root root  184 Apr 17 15:21 .editorconfig*
drwxrwxrwx 1 root root 4096 Apr 17 15:25 .git/
-rwxr-xr-x 1 root root  301 Apr 17 15:21 .gitignore*
-rwxr-xr-x 1 root root  734 Apr 17 16:32 Dockerfile*
-rwxr-xr-x 1 root root 1094 Apr 17 15:14 LICENSE.txt*
-rwxr-xr-x 1 root root 2023 Apr 17 16:41 README.md*
-rwxr-xr-x 1 root root   23 Apr 17 15:24 go.mod*
-rwxr-xr-x 1 root root   76 Apr 17 16:07 main.go*
```

Start deving
```
$ go run ./main.go
hello world!
```

## Contributing
This repository is under heavy development and is subject to change in the near future.</br>
Versioning will be locked and a proper contributing section will be created in a timely manner, when code is stabilized.</br>

## License
`go-hub` is released under the MIT license. See [LICENSE.txt](LICENSE.txt)
