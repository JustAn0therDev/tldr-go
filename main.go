package main

import (
    "fmt"
    "os"
    "io/ioutil"
)


func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func GetFileContent(path string, fileName string) ([]byte, error) {
    pathWithFileName := fmt.Sprintf("%v/%v", path, fileName)
    return ioutil.ReadFile(pathWithFileName)
}

// main is all lower-case because we don't want to export it to other files
// every property or function that starts in upper case is "importable" from other files
func main() {
    // the first argument is always the name of the program (in this case, "tldr-go")
    // For now, the only thing that will be accepted as a command-line argument is the name
    // of a program. A directory inside the base program will have the same name and a bunch of
    // ".md" files with command descriptions.
    pathArg := os.Args[1]
    fileName := os.Args[2]

    pathExists, err := PathExists(pathArg)

    checkError(err)

    if !pathExists {
        panic("The path does not exist")
    }

    buffer, err := GetFileContent(pathArg, fileName)

    checkError(err)

    fmt.Println(string(buffer))
}
