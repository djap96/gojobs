//Pick a random pic inside {wallpapersDir} and change the **Gnome-Desktop** with it

package main

import (
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

const wallpapersDir = "/home/djap96/Pictures/wall/"

func newImage() string {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	files, _ := ioutil.ReadDir(wallpapersDir)

	randInt := random.Intn(len(files))
	fileName := files[randInt].Name()

	if strings.Contains(fileName, ".jpg") {
		return fileName
	}
	if strings.Contains(fileName, ".png") {
		return fileName
	}

	return newImage()
}

func main() {

	filePwd := "'file://" + wallpapersDir + newImage() + "'"
	exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", filePwd).Run()
}
