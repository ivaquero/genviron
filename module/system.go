package module

import (
	"fmt"
	"os/exec"
)

func update(theOS string) {
	println("Installing needed updates")

	if theOS == "darwin" {
		cmd := exec.Command("softwareupdate", "-i", "-a")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running softwareupdate:", err)
		}
	}
}
