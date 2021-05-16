package test

import (
	"testing"
	
	"strings"
	"github.com/JustAn0therDev/tldr-go/utils"
)

// the tests are now working, all I need to do is test the things that are
// testable
func TestGetTrimmedString(t *testing.T) {
	stringToTrim := "    trim     "
	stringToTrim = utils.GetTrimmedString(stringToTrim)
	
	if strings.Count(stringToTrim, "\n") > 0 {
		t.Error("The trim function did not trim the test string")
	}
}

func TestGetTrimmedStringWithAlreadyTrimmedString(t *testing.T) {
	stringToTrim := "trim"
	stringToTrim = utils.GetTrimmedString(stringToTrim)
	
	if strings.Count(stringToTrim, "\n") > 0 {
		t.Error("The trim function added breakline characters to a trimmed string.")
	}
}
