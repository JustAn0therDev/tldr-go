package test

import (
	"testing"

	"github.com/JustAn0therDev/tldr-go/mdparser"
)

func TestPathExists(t *testing.T) {
	exists, err := mdparser.PathExists("../main/dotnet")

	if err != nil {
		t.Errorf("An error occurred in the PathExists func: %v", err)
	}
	
	if !exists {
		t.Error("The PathExists function should have returned true")
	}
}

func TestPathExistsShouldReturnFalse(t *testing.T) {
	exists, err := mdparser.PathExists("../main/sfds")

	if err != nil {
		t.Errorf("An error occurred in the PathExists func: %v", err)
	}
	
	if exists {
		t.Error("The PathExists function should have returned false")
	}
}

func TestGetAllMdFileNamesInPath(t *testing.T) {
	path := "../main/dotnet"
	fileNames := mdparser.GetAllMdFileNamesInPath(path)

	if len(fileNames) == 0 {
		t.Error("The file name slice returned empty")
	}
}

func TestGetAllMdFileNamesInPathShouldReturnEmpty(t *testing.T) {
	path := "dotnet"
	fileNames := mdparser.GetAllMdFileNamesInPath(path)

	if len(fileNames) > 0 {
		t.Error("The file name slice should have returned empty")
	}
}

func TestGetFileContent(t *testing.T) {
	contentString, err := mdparser.GetFileContent("../main/dotnet", "dotnet.md")

	if err != nil {
		t.Errorf("An error occurred in the GetFileContent func: %v", err)
	}
	
	if contentString == "" {
		t.Error("The content string shouldn't have returned empty")
	}
}

