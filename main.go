package main

import (
	"fmt"

	"github.com/adithyavisnu/go-web-template/data"
	"github.com/adithyavisnu/go-web-template/server"
)

func main() {
	fmt.Println("Starting Application")
	server.Start()
	data.newDataStore()
}
