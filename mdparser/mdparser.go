package mdparser

import (
    "strings"
    "github.com/fatih/color"
)

// TODO: Refactor this function
// the package I'm using does not return errors anywhere. I feel like this is bad, it should fail at some point and it's probably 
// calling panic and just cutting the execution of the main program altogether
func ParseMarkdownString(mdContent string) {
    currentLine := ""
    iterator := 0
    colorInstance := color.New()

    for iterator < len(mdContent) {
        currentRune := mdContent[iterator]

        // I still don't know if this is bad code just because its not using some design pattern
        // or because it has a bunch of if statements in it. This is something that I definitely need to think about

        // Thinking about the use of "if statements" is simply over engineering or just being cautious so the code does not because a bigger mess in the future.
        if currentRune == '\n' {

            if isBiggestHeaderMarkup(currentLine) {
                colorInstance.Add(color.FgRed)
                colorInstance.Add(color.Bold)

                currentLine = getStringTrimmedAndWithNoHeaderMarkups(currentLine)
            } else if isHeaderMarkup(currentLine) {
                colorInstance.Add(color.FgWhite)
                colorInstance.Add(color.Bold)

                currentLine = getStringTrimmedAndWithNoHeaderMarkups(currentLine)
            } else if isCodeMarkup(currentLine) {
                currentLine = getStringTrimmedAndWithNoCodeMarkups(currentLine)
                colorInstance.Add(color.FgCyan)
            } else {
               colorInstance.Add(color.FgWhite)
            }

            colorInstance.Println(currentLine)

            currentLine = ""
        } else {
            currentLine += string(currentRune)
        }

        iterator++
    }
}

func isBiggestHeaderMarkup(currentLine string) bool {
    return strings.Count(currentLine, "#") == 1
}

func isHeaderMarkup(currentLine string) bool {
    return strings.Count(currentLine, "#") >= 2
}

func isCodeMarkup(currentLine string) bool {
   return strings.Count(currentLine, "`") >= 1
}

func getStringTrimmedAndWithNoHeaderMarkups(currentLine string) string {
    replacedString := strings.ReplaceAll(currentLine, "#", "")
    return getTrimmedString(replacedString)
}

func getStringTrimmedAndWithNoCodeMarkups(currentLine string) string {
    replacedString := strings.ReplaceAll(currentLine, "`", "")
    return getTrimmedString(replacedString)
}

func getTrimmedString(currentLine string) string {
    return strings.Trim(currentLine, " ")
}
