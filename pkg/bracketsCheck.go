package art

import "strings"


func unbalancedBracketsCheck(inputData string) {
	totalLeftBrackets := 0
	totalRightBrackets := 0

	for _, char := range inputData {
		if char == '[' {
			totalLeftBrackets += 1
		} else if char == ']' {
			totalRightBrackets += 1
		}
	}

	if totalLeftBrackets != totalRightBrackets {
		printError(UNBALANCED_BRACKETS)
	}
}

func containsBrackets(extractArtSequenceStringFromArgs string) ([]string, bool) {
	splitSequenceAtRightBracket := strings.SplitAfter(extractArtSequenceStringFromArgs, "]")
	splitSequenceAtLeftBracket := strings.SplitAfter(extractArtSequenceStringFromArgs, "[")

	// splitAfter returns slice of length 1 if string does not contain seperator
	if len(splitSequenceAtRightBracket) == 1 && len(splitSequenceAtLeftBracket) == 1 {
		return splitSequenceAtRightBracket, false
	} else {
		return splitSequenceAtRightBracket, true
	}

}
