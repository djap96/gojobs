package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/djap96/gojobs/utils"
)

func getSubString(str, character string) string {
	lastIndex := strings.LastIndex(str, character)

	if lastIndex == -1 {
		return ""
	}

	return str[lastIndex+1:]
}

func getCustomOption(optionIdentifier string, args []string) string {
	for i := 3; i < len(args); i++ {
		argument := args[i]

		if strings.EqualFold(argument[0:2], optionIdentifier+"=") {
			return getSubString(argument, "=")
		}
	}
	return ""
}

func getDashedString(originalString string) string {

	var dashedString string

	for _, c := range originalString {
		switch c {
		case '_':
			dashedString += "__"
		case '-':
			dashedString += "--"
		case ' ':
			dashedString += "_"
		default:
			dashedString += fmt.Sprintf("%c", c)
		}
	}

	return dashedString
}

func main() {

	args := os.Args
	n := len(args)

	binaryPath := args[0]
	binaryName := getSubString(binaryPath, "/")

	if n < 3 {
		fmt.Println("\nUsage:")

		fmt.Println("-> You need at lest the name and the value of the badge!")
		fmt.Print("-> Please use the following syntaxis: ")
		fmt.Print(binaryName + " <name> <value> [c=lightgray] [f=name-value-color.svg]\n")
		fmt.Println("\nExamples:")
		fmt.Println("----> " + binaryName + " name value")
		fmt.Println("----> " + binaryName + " name value c=\"orange\"")
		fmt.Println("----> " + binaryName + " \"My name\" value c=yellow")
		fmt.Println("----> " + binaryName + " name \"My value\" c=ff0000 f=image.svg")
		fmt.Println("----> " + binaryName + " \"Go version\" 1.6 f=\"My image.svg\"")
		return
	}

	name := args[1]
	value := args[2]

	name = getDashedString(name)
	value = getDashedString(value)

	color := "lightgray"

	customColor := getCustomOption("c", args)
	customName := getCustomOption("f", args)

	if customColor != "" {
		if hexColor, ok := webcolors.GetHexValue(customColor); ok {
			color = hexColor
		} else {
			return
		}
	}

	fileName := name + "-" + value + "-" + color + ".svg"

	if customName != "" {
		if !strings.Contains(customName, ".") {
			customName += ".svg"
		}
		fileName = customName
	}

	apiRequestString := name + "-" + value + "-" + color + ".svg"

	urlPetition := "https://img.shields.io/badge/" + apiRequestString + "?style=flat-square"

	fmt.Println("Fetching from: " + urlPetition)

	response, _ := http.Get(urlPetition)
	body := response.Body
	defer body.Close()

	if !strings.Contains(response.Header.Get("Content-Type"), "image/svg+xml") {
		fmt.Println("img.shields.io is not responding right now, try later..")
		return
	}

	image, _ := os.Create(fileName)
	defer image.Close()

	io.Copy(image, body)
	fmt.Println(fileName + " saved!!!")
}
