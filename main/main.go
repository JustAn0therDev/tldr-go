package main

import (
    "fmt"
    "os"
    "github.com/JustAn0therDev/tldr-go/mdparser"
    "github.com/JustAn0therDev/tldr-go/utils"
)

func main() {
    // the first argument is always the name of the program (in this case, "tldr-go")
    // For now, the only thing that will be accepted as a command-line argument is the name
    // of a program. A directory inside the base program will have the same name and a bunch of
    // ".md" files with command descriptions.
    pathArg := os.Args[1]
    fileName := os.Args[2]

    pathExists, err := mdparser.PathExists(pathArg)

    utils.CheckError(err)

    if !pathExists {
        panic("The path does not exist")
    }

    buffer, err := mdparser.GetFileContent(pathArg, fileName)

    utils.CheckError(err)

    fmt.Println(string(buffer))
}
