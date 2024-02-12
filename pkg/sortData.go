package art

import (
	"fmt"
	"regexp"
	"strconv"
)

func sortBracketedAndNonBracketedStrings(splitSequenceAtRightBracket []string) []string {

	var splitIntoBracketDataAndSingleData []string // sort into bracketed and non-bracketed strings

	for _, section := range splitSequenceAtRightBracket {
		for i, char := range section {
			if i > 0 && char == '[' {
				splitIntoBracketDataAndSingleData = append(splitIntoBracketDataAndSingleData, section[:i])
				splitIntoBracketDataAndSingleData = append(splitIntoBracketDataAndSingleData, section[i:])
				break

			} else if char == '[' {
				splitIntoBracketDataAndSingleData = append(splitIntoBracketDataAndSingleData, section[i:])
				break
			}
		}
	}

	return splitIntoBracketDataAndSingleData

}

func useRegExToValidateData(splitIntoBracketDataAndSingleData []string) {
	// analyse each string in splitIntoBracketDataAndSingleData to validate structure
	// [5 #]  <-- regexp to match this where # can one or more of any character (including a space) and 5 can be one or more digits
	squareBracketRegExpPattern := `\[\d+\s.+?\]`

	squareBracketRegExpPatternCompile := regexp.MustCompile(squareBracketRegExpPattern)
	newLineCount := 1

	for _, data := range splitIntoBracketDataAndSingleData {
		if data[0] == '[' {
			validDataStructure := squareBracketRegExpPatternCompile.FindAllStringSubmatch(data, -1)
			if validDataStructure == nil {
				if len(splitIntoBracketDataAndSingleData) > 1 { // is multiline
					errorData := data + "check line:" + strconv.Itoa(newLineCount)
					printError(FORMAT_ERROR, errorData)
				} else {
					printError(FORMAT_ERROR, data)
				}
			} else {
				continue
			}
		} else if data == "\n" {
			newLineCount += 1
		} else {
			continue
		}
	}

}

func readStringAndPrint(splitIntoBracketDataAndSingleData []string) {
	// read each string and print output
	for _, data := range splitIntoBracketDataAndSingleData {
		if data[0] == '[' {
			// bracketed data
			var extractedDigits string
			var extractedSymbols string
			mandatorySpaceCount := false
			for _, char := range data {
				if char == '[' || char == ']' { // if brackets -> ignore
					continue
				} else if char >= '0' && char <= '9' && !mandatorySpaceCount {
					extractedDigits += string(char)
				} else if char >= '0' && char <= '9' && mandatorySpaceCount {
					extractedSymbols += string(char)
				} else if char == ' ' && !mandatorySpaceCount { // mandatory space -> ignore
					mandatorySpaceCount = true
					continue
				} else if char == ' ' && mandatorySpaceCount { // printed space
					extractedSymbols += " "
				} else {
					extractedSymbols += string(char)
				}
			}
			mandatorySpaceCount = false
			// method to convert extractedDigits into single integer
			extractedDigitsInteger, err := strconv.Atoi(extractedDigits)
			if err != nil {
				printError(FORMAT_ERROR, extractedDigits)
			} else {
				for x := 0; x < extractedDigitsInteger; x++ {
					fmt.Print(extractedSymbols)
				}
			}
		} else {
			fmt.Print(data) // print unbracketed data
		}
	}
	fmt.Println()

}
