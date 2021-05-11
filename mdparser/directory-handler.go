package mdparser

import (
    "fmt"
    "os"
    "io/ioutil"
)


// TODO: Create a struct that will have these functions as methods and maintain state with
// path, so the GetFileContent function and parsing of the ".md" file will be its responsibility

// checks if a given path exists in the current application directory.
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

// gets the file content by path and filename
func GetFileContent(path string, fileName string) ([]byte, error) {
    pathWithFileName := fmt.Sprintf("%v/%v", path, fileName)
    return ioutil.ReadFile(pathWithFileName)
}
