package art

import (
	"fmt"
	"strconv"
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
				if i == len(line)-1 { // if end of current
					if lineNum < len(splitStringFromArgs)-1 {
						matchCount += 1
						currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]\n")
						finalArtEncoded += currentSymbol
						duplicateSymbol = false
						matchCount = 0
					} else { // if last line of input (exclude new line)
						matchCount += 1
						currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")
						finalArtEncoded += currentSymbol
						duplicateSymbol = false
						matchCount = 0

					}
				} else {
					matchCount += 1
				}
			} else if !duplicateSymbol && matchCount > 0 {
				matchCount += 1
				currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")
				finalArtEncoded += currentSymbol
				matchCount = 0
			} else if !duplicateSymbol && matchCount == 0 {
				if i == len(line)-1 && lineNum < len(splitStringFromArgs)-1 {
					currentSymbol = string(line[i]) + "\n"
					finalArtEncoded += currentSymbol
				} else {
					currentSymbol = string(line[i])
					finalArtEncoded += currentSymbol
				}
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
