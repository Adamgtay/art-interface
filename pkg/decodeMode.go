package art

import "fmt"

func decodeInput(extractArtSequenceStringFromArgs string) {

	unbalancedBracketsCheck(extractArtSequenceStringFromArgs)
	splitSequenceAtRightBracket, containsBrackets := containsBrackets(extractArtSequenceStringFromArgs)

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

		readStringAndPrint(splitIntoBracketDataAndSingleData)

	} else {
		// here logic if no brackets in input
		fmt.Println(extractArtSequenceStringFromArgs)
	}

}
