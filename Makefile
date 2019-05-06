.PHONY: start
start:
	@GO111MODULE=off go run cmd/gws/main.go -port=8080 -path=/postreceive -command='echo "hello world"' -secret=mysecret

.PHONY: curl
curl:
	@curl "http://localhost:8080/postreceive" -H 'X-Hub-Signature: mysecret'
.PHONY: release
release:
	@rm -rf dist
	@goreleaser
