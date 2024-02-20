package art_interf

import (
	"fmt"
	"strings"
)

// split into smaller functions
func EncodeInput(inputString string) {
	var finalArtEncoded string
	var currentSymbol string
	splitStringFromArgs := strings.Split(inputString, "\n")
	for lineNum, line := range splitStringFromArgs {
		duplicateSymbol := false
		matchCount := 0
		for i := 0; i <= len(line)-1; i++ {
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
	if finalArtEncoded[len(finalArtEncoded)-1] == '\n' {
		finalArtEncoded = finalArtEncoded[:len(finalArtEncoded)-1]
		fmt.Println(finalArtEncoded)
	} else {
		fmt.Println(finalArtEncoded)
	}

}
