package main

import (
	"fmt"
	"opal/internal/config"
	"opal/internal/httpserver"
	"opal/internal/librarymgmt"
	"opal/internal/usermgmt"
	"os"
)

// TODO: use linker version numbers for releases
var version_number = "dev"

func main() {
	fmt.Println("Opal media server starting \nVersion:", version_number)

	//TODO: make configurable
	dbDir := "./db"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		fmt.Println("Error: failed to mkdir ./db")
		os.Exit(1)
	}

	config.Init()
	usermgmt.Init()
	librarymgmt.Init()

	httpserver.Start()
}
