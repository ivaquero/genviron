package main

import (
	"encoding/json"

	"log"
	"os"
	"os/exec"
	"strings"
)

func executeScript(scriptPath string) {
	cmd := exec.Command("bash", "-c", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Define the GENVIRON environment variable, if not set it defaults to $HOME/genviron
	homeDir := os.Getenv("HOME")
	genPath := os.Getenv("GENVIRON")

	// Let's first read the `config.json` file
	content, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `items`
	var items interface{}
	err = json.Unmarshal(content, &items)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	if genPath == "" {
		defaultPath := homeDir + "/genviron"
		genPath = defaultPath
	}

	GenPlugins := map[string]string{
		"gpd":   genPath + "/defaults/wezterm.lua",
		"gpod":  genPath + "/plugins/os-debians.go",
		"gpom":  genPath + "/plugins/os-macos.go",
		"gpor":  genPath + "/plugins/os-rehats.go",
		"gpow":  genPath + "/plugins/os-windows.go",
		"gpcbw": genPath + "/plugins/cli-bitwarden.go",
		"gpces": genPath + "/plugins/cli-espanso.go",
		"gpcjr": genPath + "/plugins/cli-jupyter.go",
		"gpcol": genPath + "/plugins/cli-ollama.go",
		"gpcvs": genPath + "/plugins/cli-vscode.go",
		"gpljl": genPath + "/plugins/lang-julia.go",
		"gplrb": genPath + "/plugins/lang-ruby.go",
		"gplrs": genPath + "/plugins/lang-rust.go",
		"gppb":  genPath + "/plugins/pkg-brew.go",
		"gppc":  genPath + "/plugins/pkg-conda.go",
		"gppnj": genPath + "/plugins/pkg-npm.go",
		"gpppx": genPath + "/plugins/pkg-pixi.go",
		"gpps":  genPath + "/plugins/pkg-scoop.go",
		"gpptl": genPath + "/plugins/pkg-tlmgr.go",
		"gpuf":  genPath + "/plugins/utils-files.go",
		"gpufm": genPath + "/plugins/utils-formats.go",
		"gpunw": genPath + "/plugins/utils-networks.go",
		"gpxns": genPath + "/plugins/xtra-notes.go",
	}

	// Load the plugins
	uname, err := exec.Command("uname", "-a").Output()
	if err != nil {
		panic(err)
	}

	systemInfo := string(uname)
	switch {
	case strings.Contains(systemInfo, "Darwin"):
		executeScript(GenPlugins["gpom"])
	case strings.Contains(systemInfo, "Ubuntu") || strings.Contains(systemInfo, "Debian") || strings.Contains(systemInfo, "WSL"):
		executeScript(GenPlugins["gpod"])
	case strings.Contains(systemInfo, "NT"):
		executeScript(GenPlugins["gpow"])
	}

	// System configuration files
	GenSys := map[string]string{
		"gen": genPath + "/custom.go",
		"g":   homeDir + "/.gitconfig",
		"vi":  homeDir + "/.vimrc",
	}

	executeScript(GenSys["gen"])
}
