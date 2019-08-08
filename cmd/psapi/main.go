package main

import (
	"flag"
	"log"
	"time"

	"github.com/RTradeLtd/PinningServiceAPI/api"
	"github.com/RTradeLtd/rtfs/v2"
)

var (
	ipfsURL = flag.String("ipfs.url", "127.0.0.1:5001", "ipfs http api endpoint")
)

func init() {
	flag.Parse()
}

func main() {
	manager, err := rtfs.NewManager(*ipfsURL, "", time.Hour)
	if err != nil {
		log.Fatal(err)
	}
	router := api.NewAPI(manager)
	router.Run(":8080")
}
