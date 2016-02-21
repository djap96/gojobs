//Pick a random pic inside {wallpapersDir} and change the **Gnome-Desktop** with it

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type folder struct {
	alias, path string
}

var folders = []folder{
	{"art", "/home/djap96/Pictures/wall/art/"},
	{"chicks", "/home/djap96/Pictures/wall/chicks/"},
	{"movies", "/home/djap96/Pictures/wall/movies/"},
	{"universe", "/home/djap96/Pictures/wall/universe/"},
	{"world", "/home/djap96/Pictures/wall/world/"},
}

var picsExtensions = []string{".png", ".jpg"}

func newImage() string {

	seed := time.Now().Unix()
	random := rand.New(rand.NewSource(seed))

	var randInt int
	var randFolder string

	if len(os.Args) > 1 {
		var founded bool

		for _, f := range folders {
			if os.Args[1] == f.alias {
				randFolder = f.path
				founded = true
			}
		}

		if !founded {
			fmt.Println("Given alias doesn't exists!!")
			return ""
		}
	} else {
		randInt = random.Intn(len(folders))
		randFolder = folders[randInt].path
	}

	files, _ := ioutil.ReadDir(randFolder)
	picsNumber := len(files)

	if picsNumber == 0 {
		fmt.Println("Cool images not found!!")
		return ""
	}

	randInt = random.Intn(picsNumber)

	fileName := files[randInt].Name()

	for _, ext := range picsExtensions {
		if strings.Contains(fileName, ext) {
			return randFolder + fileName
		}
	}

	return newImage()
}

func main() {

	imagePath := newImage()

	if imagePath == "" {
		return
	}

	filePwd := "'file://" + newImage() + "'"
	exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", filePwd).Run()
}
