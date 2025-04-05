package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tidwall/gjson"
)

var theOS = runtime.GOOS
var theArch = runtime.GOARCH
var homeDir = os.Getenv("HOME")

func SetSysEnvVar() {

	switch theOS {
	case "darwin":
		if theArch == "arm64" {
			os.Setenv("TERMINFO", "/usr/share/terminfo")
		} else {
			os.Setenv("TERMINFO", "/usr/local/share/terminfo")
		}
		os.Setenv("APPDATA", homeDir+"/Library/'Application Support'")
		os.Setenv("Cache", homeDir+"/Library/Cache")
	}
}

func ReadPathConfig(genvDir string) {
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

	appDataDir := os.Getenv("APPDATA")
	switch theOS {
	case "darwin":
		GENVSYS["lg"] = appDataDir + "/lazygit/config.yml"
		GENVSYS["wz"] = homeDir + "/.config/wezterm/wezterm.lua"
	case "linux":
		GENVSYS["lg"] = homeDir + "/.config/lazygit/config.yml"
		GENVSYS["wz"] = homeDir + "/.config/wezterm/wezterm.lua"
	case "mingw":
		GENVSYS["wz"] = homeDir + "/.wezterm.lua"
	}
}

func main() {
	genvDirDefault := homeDir + "/genviron"

	if os.Getenv("GENVIRON") == "" {
		os.Setenv("GENVIRON", genvDirDefault)
	}
	genvDir := os.Getenv("GENVIRON")

	SetSysEnvVar()
	ReadPathConfig(genvDir)
}
