package main

import (
	"github.com/RTradeLtd/PinningServiceAPI/api"
)

func main() {
	router := api.NewAPI()
	router.Run(":8080")
}
