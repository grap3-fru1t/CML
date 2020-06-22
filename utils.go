// Tool used to build the package
package main

import (
    "fmt"
    "encoding/json"
    "os"
)

type Configuration struct {
    Port              int
    Base_url          string
    Token             string
}

// Extrct the necessary information from the config
func read_config() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("Error when reading config:", err)
	}
	return configuration
}