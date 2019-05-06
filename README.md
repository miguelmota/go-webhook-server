# go-webhook-server

> A simple server to receive webhooks and execute commands

## Install

```bash
go get -u github.com/miguelmota/go-webhook-server/cmd/gws
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
