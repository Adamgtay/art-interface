package art_interf

import (
	"strings"
	"unicode/utf8"
)

func EncodeInput(inputString string) string {
	// Replace non-breaking spaces with regular spaces
	inputString = strings.ReplaceAll(inputString, "\u00A0", " ")

	var finalArtEncoded string
	var currentSymbol string
	splitStringFromArgs := strings.Split(inputString, "\n")
	for lineNum, line := range splitStringFromArgs {
		duplicateSymbol := false
		matchCount := 0
		for i := 0; i < utf8.RuneCountInString(line); i++ {
			duplicateSymbol = isDuplicateSymbol(i, line)
			if duplicateSymbol {
				currentSymbol, matchCount = ifDuplicateSymbol(i, lineNum, matchCount, line, splitStringFromArgs)
				finalArtEncoded += currentSymbol
				duplicateSymbol = false
			} else if !duplicateSymbol && matchCount > 0 {
				currentSymbol = endOfDuplicateSymbols(i, matchCount, line)
				finalArtEncoded += currentSymbol
				matchCount = 0
			} else if !duplicateSymbol && matchCount == 0 {
				currentSymbol = ifSingleSymbol(i, lineNum, line, splitStringFromArgs)
				finalArtEncoded += currentSymbol
			}
		}
	}
	return finalArtEncoded
}
