package main

import (
	art "art/pkg"
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
	flag.Bool("e", false, art.ENCODE_TEXT)
	flag.Bool("h", false, art.USAGE_TEXT)
	flag.Bool("d", false, art.DECODE_TEXT)
	flag.Bool("m", false, art.MULTILINE_TEXT)

	flag.Parse()

	if (len(args) != 3 && len(args) != 4) || args[1] == "-h" {
		if args[1] == "-h" {
			art.PrintUsage()
		} else if len(args) != 3 && len(args) != 4 {
			if len(args) < 3 {
				art.PrintError(art.MISSING_ARG)
			} else {
				errorData := strings.Join(args[4:], " ")
				art.PrintError(art.EXTRA_ARG, errorData)
			}
		}
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
			art.PrintError(art.INVALID_ARG, args[1])
		}
	} else if len(args) == 3 { // multiline
		switch args[1] {
		case "-e":
			encodeMode = true
		case "-d":
			decodeMode = true
		default:
			art.PrintError(art.INVALID_ARG, args[1])
		}
		switch args[2] {
		case "-m":
			multiLine = true
		default:
			art.PrintError(art.INVALID_ARG, args[2])
		}
	}

	if multiLine {
		if !art.IsFilePath(args[0]) {
			art.PrintError(art.INVALID_FILEPATH, args[0])
		}
		extractArtSequenceStringFromArgs, err = art.ReadFile(args[0]) // txt file read
		if err != nil {
			art.PrintError(art.ERROR_READFILE, args[0])
		}
	} else if !multiLine && art.IsFilePath(args[0]) {
		art.PrintError(art.MISSING_ARG_FOR_MULTILINE)
	} else {
		extractArtSequenceStringFromArgs = args[0]
	}

	if decodeMode {
		art.DecodeInput(extractArtSequenceStringFromArgs)

	}

	if encodeMode {
		art.EncodeInput(extractArtSequenceStringFromArgs)
	}
}
