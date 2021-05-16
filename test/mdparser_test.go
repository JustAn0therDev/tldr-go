package test

import (
	"testing"
	
	"github.com/JustAn0therDev/tldr-go/mdparser"
)

func TestIsBiggestHeaderMarkup(t *testing.T) {
	biggestHeaderMarkupString := "#header"
	if !mdparser.IsBiggestHeaderMarkup(biggestHeaderMarkupString) {
		t.Errorf("The result for %v was should have been true", biggestHeaderMarkupString)
	}
}

func TestIsBiggestHeaderMarkupShouldReturnFalse(t *testing.T) {
	notABiggestHeaderMarkupString := "header"
	if mdparser.IsBiggestHeaderMarkup(notABiggestHeaderMarkupString) {
		t.Errorf("The result for %v was should have been false", notABiggestHeaderMarkupString)
	}
}

func TestIsHeaderMarkup(t *testing.T) {
	headerMarkupString := "#header#"
	if !mdparser.IsHeaderMarkup(headerMarkupString) {
		t.Errorf("The result for %v was should have been true", headerMarkupString)
	}
}

func TestIsHeaderMarkupShouldReturnFalse(t *testing.T) {
	notAHeaderMarkupString := "#test"
	if mdparser.IsHeaderMarkup(notAHeaderMarkupString) {
		t.Errorf("The result for %v was should have been false", notAHeaderMarkupString)
	}
}

func TestIsCodeMarkup(t *testing.T) {
	codeMarkupString := "`code string`"
	if !mdparser.IsCodeMarkup(codeMarkupString) {
		t.Errorf("The result for %v was should have been true", codeMarkupString)
	}
}

func TestIsCodeMarkupShouldReturnFalse(t *testing.T) {
	notACodeMarkupString := "code string"
	if mdparser.IsCodeMarkup(notACodeMarkupString) {
		t.Errorf("The result for %v was should have been false", notACodeMarkupString)
	}
}
