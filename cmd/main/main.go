package main

import (
	"FleetManagerAPI/cmd/server"
	"FleetManagerAPI/config"
)

func main() {
	config.LoadAllConfigs(".env")
	server.Serve()
}