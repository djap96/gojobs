package main

import (
	"os"
	"os/exec"
)

func main() {

	volume := "150%"
	sinkEntity := "0"

	if len(os.Args) > 1 {
		volume = os.Args[1] + "%"
	}

	if len(os.Args) > 2 {
		sinkEntity = os.Args[2]
	}

	exec.Command("pactl", "set-sink-volume", sinkEntity, volume).Run()
}
