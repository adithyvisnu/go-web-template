package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/adithyavisnu/go-web-template/configurations"
	"github.com/adithyavisnu/go-web-template/data"
	"github.com/adithyavisnu/go-web-template/server"
)

var (
	ds         *data.DataStore
	configfile = flag.String("config", "/etc/smtpd.conf", "Path to the configuration file")
)

func main() {
	fmt.Println("Loading configuration files...")
	err := configurations.LoadConfig(*configfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connecting to databases...")
	ds = data.NewDataStore()

	fmt.Println("Starting Application")
	server.Start()
	server.Initialize(configurations.GetWebConfig(), ds)
}
