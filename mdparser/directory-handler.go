package mdparser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/JustAn0therDev/tldr-go/utils"
)

// PathExists: checks if a given path exists in the current application directory.
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

// GetAllMdFileNamesInPath: gets all filenames in the current directory that have the ".md" extension
func GetAllMdFileNamesInPath(path string) []string {
    fileNames := []string{}
    var fileName string

	pathExists, err := PathExists(path)

	if !pathExists {
		return make([]string, 0)
	}

	utils.CheckError(err)

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

// GetFileContent: gets the file content by path and filename
func GetFileContent(path string, fileName string) (string, error) {
    pathWithFileName := fmt.Sprintf("%v/%v", path, fileName)
    buffer, err := ioutil.ReadFile(pathWithFileName)
    utils.CheckError(err)
    return string(buffer), nil
}
