package main

import (
	"os"
)

func main() {
	// Define the OXIDIZER environment variable, if not set it defaults to $HOME/oxidizer
	OxidizerPath := os.Getenv("HOME") + "/oxidizer"

	if os.Getenv("OXIDIZER") == "" {
		os.Setenv("OXIDIZER", OxidizerPath)
		print("Path of OXIDIZER: ", OxidizerPath, "\n")
	} else {
		print("Path of OXIDIZER exists: ", os.Getenv("OXIDIZER"), "\n")
	}
}
