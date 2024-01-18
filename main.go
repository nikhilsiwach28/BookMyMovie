package main

import (
	"fmt"
	"github.com/nikhilsiwach28/BookMyMovie/Config/config"
)

func main() {
	fmt.Println("Hello World!")
	// GetConfigs
	// CreateServer(configs)
	config := GetConfigFromEnv()
	startServer()

}
