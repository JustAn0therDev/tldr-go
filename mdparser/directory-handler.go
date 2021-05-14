package mdparser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/JustAn0therDev/tldr-go/utils"
)

// TODO: Create a struct that will have these functions as methods and maintain state with
// path, so the GetFileContent function and parsing of the ".md" file will be its responsibility

// TODO: Actually... I'm thinking about making this more functional. State management will be something way more
// complex to maintain along the way. I'd rather manage a rolling thingy from where I can manage the data with
// step than using state determining how and when it will be made every step of the way

// TODO: Make the functions testable. I must make them parse the file string by receiving a string, something
// that be manipulated in runtime, not only by a file. This would have me obligated to create a test file
// and use it in every unit test and have it versioned.

// checks if a given path exists in the current application directory.
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func GetAllMdFileNamesInPath(path string) []string {
    fileNames := []string{}
    var fileName string

    dir, err := os.Open(path)

    utils.CheckError(err)

    defer dir.Close()

    files, err := dir.ReadDir(-1)

    utils.CheckError(err)

    for _, file := range files {
        fileName = file.Name()
        if file.Type().IsRegular() && filepath.Ext(fileName) == ".md" {
            fileNames = append(fileNames, file.Name())
        }
    }

    return fileNames
}

// gets the file content by path and filename
func GetFileContent(path string, fileName string) (string, error) {
    pathWithFileName := fmt.Sprintf("%v/%v", path, fileName)
    buffer, err := ioutil.ReadFile(pathWithFileName)
    utils.CheckError(err)
    return string(buffer), nil
}
