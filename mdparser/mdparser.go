package mdparser

import (
	"strings"

	"github.com/fatih/color"
	"github.com/JustAn0therDev/tldr-go/utils"
)

// PrintParsedMarkdownString: prints the formatted markdown strings using the instruction map provided by its depencency: MakeInstructionByMarkupMap.
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

// MakeInstructionByMarkupMap: returns a map of strings mapped to a function that receives a string.
// The keys are "instructions" and the functions print out the formatted values to STDOUT.
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

// IsBiggestHeaderMarkup: returns a boolean value indicating if the line contains only one "#" chars (hashtag)
func IsBiggestHeaderMarkup(line string) bool {
    return strings.Count(line, "#") == 1
}

// IsHeaderMarkup: returns a boolean vlaue indicating if the line contains two or more "#" chars (hashtag)
func IsHeaderMarkup(line string) bool {
    return strings.Count(line, "#") >= 2
}

// IsCodeMarkup: returns a boolean value indication if the line contains two or more "`" chars (backtick)
func IsCodeMarkup(line string) bool {
   return strings.Count(line, "`") >= 1
}

// GetStringTrimmedAndWithNoHeaderMarkups: returns the same string trimmed and all hashtags replaced by an empty string
func GetStringTrimmedAndWithNoHeaderMarkups(line string) string {
    replacedString := strings.ReplaceAll(line, "#", "")
    return utils.GetTrimmedString(replacedString)
}

// GetStringTrimmedAndWithNoCodeMarkups: returns the same string trimmed and all backticks replaced by an empty string
func GetStringTrimmedAndWithNoCodeMarkups(line string) string {
    replacedString := strings.ReplaceAll(line, "`", "")
    return utils.GetTrimmedString(replacedString)
}
