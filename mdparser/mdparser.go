package mdparser

import (
	"strings"

	"github.com/fatih/color"
)

func PrintParsedMarkdownString(mdContent string) {
    mdContentSliceSeparatedByBreakLines := strings.Split(mdContent, "\n")
    instructionMap := makeInstructionByMarkupMap()

    for _, line := range mdContentSliceSeparatedByBreakLines {

        if isBiggestHeaderMarkup(line) {
            instructionMap["PRINT_BIGGEST_HEADER_MARKUP"](line)
        } else if isHeaderMarkup(line) {
            instructionMap["PRINT_HEADER_MARKUP"](line)
        } else if isCodeMarkup(line) {
            instructionMap["PRINT_CODE_MARKUP"](line)
        } else {
            instructionMap["PRINT_DEFAULT_MARKUP"](line)
        }
    }
}

func makeInstructionByMarkupMap() map[string]func(string) {
    mapStringToFunc := make(map[string]func(string))

    mapStringToFunc["PRINT_BIGGEST_HEADER_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgRed)
        currentColor.Add(color.Bold)

        currentColor.Println(getStringTrimmedAndWithNoHeaderMarkups(s))
    }

    mapStringToFunc["PRINT_HEADER_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgWhite)
        currentColor.Add(color.Bold)

        currentColor.Println(getStringTrimmedAndWithNoHeaderMarkups(s))
    }

    mapStringToFunc["PRINT_CODE_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgCyan)

        currentColor.Println(getStringTrimmedAndWithNoCodeMarkups(s))
    }

    mapStringToFunc["PRINT_DEFAULT_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgWhite)
        currentColor.Add(color.Bold)

        currentColor.Println(getTrimmedString(s))
    }

    return mapStringToFunc
}

func isBiggestHeaderMarkup(line string) bool {
    return strings.Count(line, "#") == 1
}

func isHeaderMarkup(line string) bool {
    return strings.Count(line, "#") >= 2
}

func isCodeMarkup(line string) bool {
   return strings.Count(line, "`") >= 1
}

func getStringTrimmedAndWithNoHeaderMarkups(line string) string {
    replacedString := strings.ReplaceAll(line, "#", "")
    return getTrimmedString(replacedString)
}

func getStringTrimmedAndWithNoCodeMarkups(line string) string {
    replacedString := strings.ReplaceAll(line, "`", "")
    return getTrimmedString(replacedString)
}

func getTrimmedString(line string) string {
    return strings.Trim(line, " ")
}
