package utils

import (
	"log"
	"strings"
)

func CheckError(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

func GetTrimmedString(line string) string {
    return strings.Trim(line, " ")
}
