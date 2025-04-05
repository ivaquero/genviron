package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tidwall/gjson"
)

func ReadPathConfig(homeDir string, genvDir string) {
	GENVBUILTIN := make(map[string]string)
	GENVSYS := make(map[string]string)

	path_file := genvDir + "/config/path.json"
	path_data, err := os.ReadFile(path_file)
	if err != nil {
		fmt.Println("Error reading config.json:", err)
		return
	}

	builtin_config_files := gjson.Get(string(path_data), "builtin_config_files")
	// println("builtin_config_files: ", builtin_config_files.String())
	for _, pair := range builtin_config_files.Array() {
		key := pair.Get("alias").String()
		value := pair.Get("path").String()
		GENVBUILTIN[key] = homeDir + value
	}

	sys_config_files := gjson.Get(string(path_data), "sys_config_files")
	// println("sys_config_files: ", sys_config_files.String())
	for _, pair := range sys_config_files.Array() {
		key := pair.Get("alias").String()
		value := pair.Get("path").String()
		GENVSYS[key] = homeDir + value
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		GENVSYS["lg"] = "/Library/Application Support/lazygit/config.yml"
		GENVSYS["wz"] = homeDir + "/.config/wezterm/wezterm.lua"
	case "linux":
		GENVSYS["lg"] = homeDir + "/.config/lazygit/config.yml"
		GENVSYS["wz"] = homeDir + "/.config/wezterm/wezterm.lua"
	case "mingw":
		GENVSYS["wz"] = homeDir + "/.wezterm.lua"
	}
}

func main() {
	homeDir := os.Getenv("HOME")
	genvDirDefault := homeDir + "/genviron"

	if os.Getenv("GENVIRON") == "" {
		os.Setenv("GENVIRON", genvDirDefault)
	}
	genvDir := os.Getenv("GENVIRON")

	ReadPathConfig(homeDir, genvDir)
}
