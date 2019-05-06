# go-webhook-server

> A simple server to receive webhooks and execute commands

## Install

Installing from Go:

```bash
go get -u github.com/miguelmota/go-webhook-server/cmd/gws
```

Installing pre-compiled binary:

```bash
$ wget https://github.com/miguelmota/go-webhook-server/releases/download/v0.0.3/gws_0.0.4_Linux_x86_64.tar.gz
$ tar -xvzf gws_0.0.4_Linux_x86_64.tar.gz gws
$ gws --help
```

## Get started

```bash
$ export SECRET_TOKEN=mysecret
$ gws -port=8080 -path=/postreceive -command='echo "hello world"'

Registered path /postreceive to run command "echo \"hello world\""
Listening on port 8080
```

```bash
$ curl "http://localhost:8080/postreceive" -X 'X-Hub-Signature: mysecret'
hello world
```

## License

[MIT](LICENSE)
