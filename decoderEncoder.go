package main

import (
	"flag"
	"os"
	"strings"
)

func main() {

	// toggles
	multiLine := false
	encodeMode := false
	decodeMode := false

	args := os.Args

	// define flags
	flag.Bool("e", false, ENCODE_TEXT)
	flag.Bool("h", false, USAGE_TEXT)
	flag.Bool("d", false, DECODE_TEXT)
	flag.Bool("m", false, MULTILINE_TEXT)

	flag.Parse()

	if (len(args) != 3 && len(args) != 4) || args[1] == "h" {
		if len(args) != 3 && len(args) != 4 {
			if len(args) < 3 {
				printError(MISSING_ARG)

			} else {
				errorData := strings.Join(args[4:], " ")
				printError(EXTRA_ARG, errorData)
			}
		}
		printUsage()
	}

	args = args[1:]

	var extractArtSequenceStringFromArgs string
	var err error

	if len(args) == 2 { // single line
		switch args[1] {
		case "-e":
			encodeMode = true
		case "-d":
			decodeMode = true
		default:
			printError(INVALID_ARG, args[1])
		}
	} else if len(args) == 3 { // multiline
		switch args[1] {
		case "-e":
			encodeMode = true
		case "-d":
			decodeMode = true
		default:
			printError(INVALID_ARG, args[1])
		}
		switch args[2] {
		case "-m":
			multiLine = true
		default:
			printError(INVALID_ARG, args[2])
		}
	}

	if multiLine {
		if !isFilePath(args[0]) {
			printError(INVALID_FILEPATH, args[0])
		}
		extractArtSequenceStringFromArgs, err = readFile(args[0]) // txt file read
		if err != nil {
			printError(ERROR_READFILE, args[0])
		}
	} else if !multiLine && isFilePath(args[0]) {
		printError(MISSING_ARG_FOR_MULTILINE)
	} else {
		extractArtSequenceStringFromArgs = args[0]
	}

	if decodeMode {
		decodeInput(extractArtSequenceStringFromArgs)

	}

	if encodeMode {
		encodeInput(extractArtSequenceStringFromArgs)
	}
}
