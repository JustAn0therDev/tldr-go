package main

import (
	"log"
	"os"

	"github.com/JustAn0therDev/tldr-go/mdparser"
	"github.com/JustAn0therDev/tldr-go/utils"
)

// init: executes as soon as the program starts.
func init() {
    log.SetFlags(0)
}

// main: program entrypoint.
func main() {
	if len(os.Args) == 1 {
		log.Fatalln("a directory argument must be passed in to the program!");
	}

    path := os.Args[1]

    pathExists, err := mdparser.PathExists(path)

    utils.CheckError(err)

    if !pathExists {
        log.Fatalln("path not found!")
    }

    fileNames := mdparser.GetAllMdFileNamesInPath(path)
	instructions := mdparser.MakeInstructionByMarkupMap()

    for _, fileName := range fileNames {
        fileContent, err := mdparser.GetFileContent(path, fileName)

        utils.CheckError(err)

        mdparser.PrintParsedMarkdownString(fileContent, instructions)
    }
}
