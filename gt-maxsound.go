package main

import (
	"os"
	"os/exec"
)

func main() {
	volume := "150%"

	if len(os.Args) > 1 {
		volume = os.Args[1] + "%"
	}

	exec.Command("pactl", "set-sink-volume", "0", volume).Run()
}
