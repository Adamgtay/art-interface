package art

import (
	"fmt"
	"regexp"
	"strconv"
)

func sortBracketedAndNonBracketedStrings(inputSliceString []string) []string {

	var outputSliceString []string // sort into bracketed and non-bracketed strings

	for _, section := range inputSliceString {
		for i, char := range section {
			if i > 0 && char == '[' {
				outputSliceString = append(outputSliceString, section[:i])
				outputSliceString = append(outputSliceString, section[i:])
				break

			} else if char == '[' {
				outputSliceString = append(outputSliceString, section[i:])
				break
			}
		}
	}

	return outputSliceString

}

func useRegExToValidateData(outputSliceString []string) {
	// analyse each string in outputSliceString to validate structure
	// [5 #]  <-- regexp to match this where # can one or more of any character (including a space) and 5 can be one or more digits
	squareBracketRegExpPattern := `\[\d+\s.+?\]`

	squareBracketRegExpPatternCompile := regexp.MustCompile(squareBracketRegExpPattern)
	newLineCount := 1

	for _, data := range outputSliceString {
		if data[0] == '[' {
			validDataStructure := squareBracketRegExpPatternCompile.FindAllStringSubmatch(data, -1)
			if validDataStructure == nil {
				if len(outputSliceString) > 1 { // is multiline
					errorData := data + "check line:" + strconv.Itoa(newLineCount)
					PrintError(FORMAT_ERROR, errorData)
				} else {
					PrintError(FORMAT_ERROR, data)
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

func readStringAndPrint(outputSliceString []string) {
	// read each string and print output
	for _, data := range outputSliceString {
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
				PrintError(FORMAT_ERROR, extractedDigits)
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

func isDuplicateSymbol(i int, line string) bool {
	if i == len(line)-1 {
		return line[i] == line[i-1]
	}
	return line[i] == line[i+1]
}

func ifDuplicateSymbol(i, lineNum, matchCount int, line string, splitStringFromArgs []string) (currentSymbol string, newMatchCount int) {
	currentSymbol = ""
	if i == len(line)-1 { // if end of current line
		if lineNum < len(splitStringFromArgs)-1 {
			matchCount += 1
			currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]\n")

			matchCount = 0
		} else { // if last line of input (exclude new line)
			matchCount += 1
			currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")

			matchCount = 0

		}
	} else {
		matchCount += 1
	}
	newMatchCount = matchCount

	return currentSymbol, newMatchCount
}

func endOfDuplicateSymbols(i, matchCount int, line string) string {
	currentSymbol := ""
	matchCount += 1
	currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")

	return currentSymbol
}

func ifSingleSymbol(i, lineNum int, line string, splitStringFromArgs []string) string {
	currentSymbol := ""
	if i == len(line)-1 && lineNum < len(splitStringFromArgs)-1 { // if end of line and not last line of input
		currentSymbol = string(line[i]) + "\n" // <-- add newline
	} else {
		currentSymbol = string(line[i])
	}
	return currentSymbol
}
