.PHONY: start
start:
	@GO111MODULE=off go run cmd/gws/main.go -port=8080 -path=/postreceive -command='echo "hello world"' -secret=mysecret

.PHONY: curl
curl:
	@curl "http://localhost:8080/postreceive" -H 'X-Hub-Signature: sha1=33f9d709782f62b8b4a0178586c65ab098a39fe2'

.PHONY: release
release:
	@rm -rf dist
	@goreleaser

.PHONY: test
test:
	@echo 'todo'
