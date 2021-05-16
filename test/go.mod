module github.com/JustAn0therDev/tldr-go/test

go 1.16

replace github.com/JustAn0therDev/tldr-go/mdparser => ../mdparser

require github.com/JustAn0therDev/tldr-go/mdparser v0.0.0-00010101000000-000000000000

replace github.com/JustAn0therDev/tldr-go/utils => ../utils

require (
	github.com/JustAn0therDev/tldr-go/utils v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.11.0 // indirect
)
