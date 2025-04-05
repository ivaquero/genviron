package main

import (
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	// Define the GENVIRON environment variable, if not set it defaults to $HOME/genviron
	homeDir := os.Getenv("HOME")
	genPath := os.Getenv("GENVIRON")

	value := gjson.Get("custom.json")

}
