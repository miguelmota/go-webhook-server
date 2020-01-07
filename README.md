# go-webhook-server

> A simple server to receive webhooks and execute commands

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/go-webhook-server/master/LICENSE)
[![Build Status](https://travis-ci.org/miguelmota/go-webhook-server.svg?branch=master)](https://travis-ci.org/miguelmota/go-webhook-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/go-webhook-server?)](https://goreportcard.com/report/github.com/miguelmota/go-webhook-server)
[![GoDoc](https://godoc.org/github.com/miguelmota/go-webhook-server?status.svg)](https://godoc.org/github.com/miguelmota/go-webhook-server)

## Install

Installing from Go:

```bash
$ go get -u github.com/miguelmota/go-webhook-server/cmd/gws
```

Installing pre-compiled binary:

```bash
$ wget https://github.com/miguelmota/go-webhook-server/releases/download/v0.0.7/gws_0.0.7_Linux_x86_64.tar.gz
$ tar -xvzf gws_0.0.7_Linux_x86_64.tar.gz gws
$ chmod +x gws
$ sudo mv gws /usr/local/bin/gws
```

## Usage

```bash
$ gws --help
```

## Getting started

Example of setting up a [Github](https://developer.github.com/webhooks/creating/) webhook:

```bash
$ export SECRET_TOKEN=mysecret
$ gws -port=8080 -path=/postreceive -command='echo "hello world"'

Registered path /postreceive to run command "echo \"hello world\""
Listening on port 8080
```

```bash
$ curl "http://localhost:8080/postreceive" -X 'X-Hub-Signature: sha1=33f9d709782f62b8b4a0178586c65ab098a39fe2'
hello world
```

Example of how to use bash script as the command:

```bash
$ cat > command.sh
#!/bin/bash
echo 'hello world'
^C

$ chmod +x command.sh

$ gws -method=POST -path=/postreceive -command=$(pwd)/command.sh
```

## License

[MIT](LICENSE)
