package utils

import (
	"log"
	"strings"
)

// CheckError: checks if an error ocurred and if so, it logs and "fails" (ends the program's execution)
func CheckError(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

// GetTrimmedString: Returns the same string trimmed
func GetTrimmedString(line string) string {
    return strings.Trim(line, " ")
}
