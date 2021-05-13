package mdparser

import (
    "strings"
    "github.com/fatih/color"
)

// TODO: Complete this function
func ParseMarkdownString(mdContent string) {
    color.Set(color.FgRed, color.Bold)
    currentString := ""
    iterator := 0

    for iterator < len(mdContent) {
        currentRune := mdContent[iterator]

        // I could just clean up the whole string with a regular expression
        // But I want different colors for different occurrences inside the markup
        if currentRune == '\n' {
            isHeaderMarkup := isHeaderMarkup(currentString)
            isCodeMarkup := isCodeMarkup(currentString)

            if isHeaderMarkup {
                c := color.New(color.FgRed)
                c.Add(color.Bold)

                currentString = getStringTrimmedAndWithNoHeaderMarkups(currentString)
                c.Println(currentString)
            }

            if isCodeMarkup {
                currentString = getStringTrimmedAndWithNoCodeMarkups(currentString)
                color.Cyan(currentString)
            }

            currentString = ""
        } else {
            currentString += string(currentRune)
        }

        iterator++
    }
}

func isHeaderMarkup(currentString string) bool {
    return strings.Count(currentString, "#") > 0
}

func isCodeMarkup(currentString string) bool {
   return strings.Count(currentString, "`") >= 1
}

func getStringTrimmedAndWithNoHeaderMarkups(currentString string) string {
    replacedString := strings.ReplaceAll(currentString, "#", "")
    return getTrimmedString(replacedString)
}

func getStringTrimmedAndWithNoCodeMarkups(currentString string) string {
    replacedString := strings.ReplaceAll(currentString, "`", "")
    return getTrimmedString(replacedString)
}

func getTrimmedString(currentString string) string {
    return strings.Trim(currentString, " ")
}
