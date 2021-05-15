package main

import (
	"log"
	"os"

	"github.com/JustAn0therDev/tldr-go/mdparser"
	"github.com/JustAn0therDev/tldr-go/utils"
)

func init() {
    log.SetFlags(0)
}

func main() {
    path := os.Args[1]

    pathExists, err := mdparser.PathExists(path)

    utils.CheckError(err)

    if !pathExists {
        log.Fatalln("path not found")
    }

    fileNames := mdparser.GetAllMdFileNamesInPath(path)

    for _, fileName := range fileNames {
        fileContent, err := mdparser.GetFileContent(path, fileName)

        utils.CheckError(err)

        mdparser.PrintParsedMarkdownString(fileContent)
    }
}
