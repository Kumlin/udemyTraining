package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Andrew Kumlin"

	tpl := `
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>Hello World!</title>
  </head>
  <body>
  <h1>` + name + `</h1>
  </body>
  </html>
  `
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer newFile.Close()

	io.Copy(newFile, strings.NewReader(tpl))

}