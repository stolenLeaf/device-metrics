package main

import (
	"github.com/stolenleaf/device-metrics/internal/bootstrap"
)

func main() {

	server := bootstrap.SetupServer()
	server.Run(":3001")
}
