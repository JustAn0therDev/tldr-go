package main

import (
    // "fmt"
    "os"
    "github.com/JustAn0therDev/tldr-go/mdparser"
    "github.com/JustAn0therDev/tldr-go/utils"
    // "github.com/fatih/color"
)

// TODO: make the program have a configurable path so that the user does not have to 
// type in the name of the file everytime they execute the command.

// For now, the only thing that will be accepted as a command-line argument is the name
// of a program. A directory inside the base program will have the same name and a bunch of
// ".md" files with command descriptions.

// TODO: Create a UI module that will have the responsibility for rendering the markdown content in different colors

func main() {
    path := os.Args[1]
    fileName := os.Args[2]

    pathExists, err := mdparser.PathExists(path)

    utils.CheckError(err)

    if !pathExists {
        panic("path does not exist")
    }

    fileContent, err := mdparser.GetFileContent(path, fileName)

    utils.CheckError(err)

    // these functions probably do not execute together because Go is asynchronous,
    // so these will execute later. With defer, we can wait for the function to end and execute the rest
    // DOING: Print this with color
    // fmt.Println(fileContent)
    mdparser.ParseMarkdownString(fileContent)
}
