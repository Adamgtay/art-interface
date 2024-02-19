package art_interf

func DecodeInput(inputString string) (string, bool) {

	var output string

	unbalancedBracketsCheck(inputString)
	if !unbalancedBracketsCheck(inputString) {
		return inputString, true
	}
	splitSequenceAtRightBracket, containsBrackets := containsBrackets(inputString)

	if containsBrackets {

		// if length of last slice is zero, remove it.
		if len(splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1]) == 0 {
			splitSequenceAtRightBracket = splitSequenceAtRightBracket[:len(splitSequenceAtRightBracket)-1]
		}

		splitIntoBracketDataAndSingleData := sortBracketedAndNonBracketedStrings(splitSequenceAtRightBracket)

		// if last string is an unbracketed symbol(s) ie. of length greater than zero, append to end of splitIntoBracketDataAndSingleData slice
		if len(splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1]) > 0 && splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1][0] != '[' {
			splitIntoBracketDataAndSingleData = append(splitIntoBracketDataAndSingleData, splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1])
		}

		useRegExToValidateData(splitIntoBracketDataAndSingleData)

		output = readString(splitIntoBracketDataAndSingleData)

	} else {
		// here logic if no brackets in input
		output = inputString
	}

	return output, false

}
