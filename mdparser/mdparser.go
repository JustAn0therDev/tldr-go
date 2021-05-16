package mdparser

import (
	"strings"

	"github.com/fatih/color"
	"github.com/JustAn0therDev/tldr-go/utils"
)

func PrintParsedMarkdownString(mdContent string) {
    mdContentSliceSeparatedByBreakLines := strings.Split(mdContent, "\n")
    instructionToPrintFunctionMap := MakeInstructionByMarkupMap()

    for _, line := range mdContentSliceSeparatedByBreakLines {

        if IsBiggestHeaderMarkup(line) {
            instructionToPrintFunctionMap["PRINT_BIGGEST_HEADER_MARKUP"](line)
        } else if IsHeaderMarkup(line) {
            instructionToPrintFunctionMap["PRINT_HEADER_MARKUP"](line)
        } else if IsCodeMarkup(line) {
            instructionToPrintFunctionMap["PRINT_CODE_MARKUP"](line)
        } else {
            instructionToPrintFunctionMap["PRINT_DEFAULT_MARKUP"](line)
        }
    }
}

func MakeInstructionByMarkupMap() map[string]func(string) {
    mapStringToFunc := make(map[string]func(string))

    mapStringToFunc["PRINT_BIGGEST_HEADER_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgRed)
        currentColor.Add(color.Bold)

        currentColor.Println(GetStringTrimmedAndWithNoHeaderMarkups(s))
    }

    mapStringToFunc["PRINT_HEADER_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgWhite)
        currentColor.Add(color.Bold)

        currentColor.Println(GetStringTrimmedAndWithNoHeaderMarkups(s))
    }

    mapStringToFunc["PRINT_CODE_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgCyan)

        currentColor.Println(GetStringTrimmedAndWithNoCodeMarkups(s))
    }

    mapStringToFunc["PRINT_DEFAULT_MARKUP"] = func(s string) {
        currentColor := color.New()
        currentColor.Add(color.FgWhite)
        currentColor.Add(color.Bold)

        currentColor.Println(utils.GetTrimmedString(s))
    }

    return mapStringToFunc
}

func IsBiggestHeaderMarkup(line string) bool {
    return strings.Count(line, "#") == 1
}

func IsHeaderMarkup(line string) bool {
    return strings.Count(line, "#") >= 2
}

func IsCodeMarkup(line string) bool {
   return strings.Count(line, "`") >= 1
}

func GetStringTrimmedAndWithNoHeaderMarkups(line string) string {
    replacedString := strings.ReplaceAll(line, "#", "")
    return utils.GetTrimmedString(replacedString)
}

func GetStringTrimmedAndWithNoCodeMarkups(line string) string {
    replacedString := strings.ReplaceAll(line, "`", "")
    return utils.GetTrimmedString(replacedString)
}
