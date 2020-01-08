package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/miguelmota/go-webhook-server/server"
)

func main() {
	var port uint = 8080
	var command string
	var secret string
	var method = "GET"
	var path = "/"

	if os.Getenv("PORT") != "" {
		portU64, err := strconv.ParseUint(os.Getenv("PORT"), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		port = uint(portU64)
	}

	if os.Getenv("SECRET") != "" {
		secret = os.Getenv("SECRET")
	}
	if os.Getenv("SECRET_TOKEN") != "" {
		secret = os.Getenv("SECRET_TOKEN")
	}

	flag.UintVar(&port, "port", port, "Port")
	flag.StringVar(&command, "command", command, "Command")
	flag.StringVar(&path, "path", path, "Path")
	flag.StringVar(&secret, "secret", secret, "Secret")
	flag.StringVar(&method, "method", method, "Method")
	flag.Parse()

	svr := server.NewServer(&server.Config{
		Port:    port,
		Path:    path,
		Command: command,
		Secret:  secret,
		Method:  method,
	})

	if err := svr.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}
