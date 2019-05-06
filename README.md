# go-webhook-server

> A simple server to receive webhooks and execute commands

## Install

Installing from Go:

```bash
$ go get -u github.com/miguelmota/go-webhook-server/cmd/gws
```

Installing pre-compiled binary:

```bash
$ wget https://github.com/miguelmota/go-webhook-server/releases/download/v0.0.6/gws_0.0.6_Linux_x86_64.tar.gz
$ tar -xvzf gws_0.0.6_Linux_x86_64.tar.gz gws
$ chmod +x gws
$ sudo mv gws /usr/local/bin/gws
```

## Usage

```bash
$ gws --help
```

## Getting started

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

## License

[MIT](LICENSE)
