package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

// TODO: proper tests

func main() {
	data, err := ioutil.ReadFile("test/payload.min.json")
	if err != nil {
		panic(err)
	}

	secret := []byte("authereum")
	message := []byte(strings.TrimSpace(string(data)))

	hash := hmac.New(sha1.New, secret)
	hash.Write(message)

	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
	expected := "de05ddb9b629dabf1617f99fc6bcc74ca188a0d3"
	fmt.Println(expected)
	// "X-Hub-Signature: sha1=de05ddb9b629dabf1617f99fc6bcc74ca188a0d3"
}
