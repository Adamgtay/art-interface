package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// function to read file from path argument
func readFile(filePath string) (string, error) {
	// Read the entire file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	return string(content), nil
}

// check argument is valid filepath
func isFilePath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {

	// toggles
	multiLine := false
	encodeMode := false
	decodeMode := false

	args := os.Args

	// define flags
	flag.Bool("e", false, "Enable \033[33mencoding\033[0m mode\n\n")
	flag.Bool("h", false, "\033[33mExample Usage:\033[0m\nsingle-line decode mode: \033[33m$ go run . \"[5 #]T\" -d\033[0m\nsingle line encode mode: \033[33m$ go run . \"#####T\" -e\033[0m\nmulti-line decode mode: \033[33m$ go run . ./filepath -d -m\033[0m\nmulti-line encode mode: \033[33m$ go run . ./filepath -e -m\033[0m\n")
	flag.Bool("d", false, "Enable \033[33mdecoding\033[0m mode\n")
	flag.Bool("m", false, "Enable \033[33mmulti-line\033[0m mode\n")

	flag.Parse()

	if (len(args) != 3 && len(args) != 4) || args[1] == "h" {
		if len(args) != 3 && len(args) != 4 {
			if len(args) < 3 {
				fmt.Println()
				fmt.Println("\033[31mError! Missing Argument:\033[0m Miniumum 3 arguments")
				fmt.Println()
				fmt.Println("Usage:")
				flag.PrintDefaults()
				os.Exit(0)

			} else {
				fmt.Println()
				fmt.Println("\033[31mError! Extra Argument:\033[0m", args[4:])
				fmt.Println()
				fmt.Println("Usage:")
				flag.PrintDefaults()
				os.Exit(0)
			}
		}
		fmt.Println()
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
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
			fmt.Println()
			fmt.Println("\033[31mError! Invalid Argument:\033[0m", args[1])
			fmt.Println()
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(0)
		}
	} else if len(args) == 3 { // multiline
		switch args[1] {
		case "-e":
			encodeMode = true
		case "-d":
			decodeMode = true
		default:
			fmt.Println()
			fmt.Println("\033[31mError! Invalid Argument:\033[0m", args[1])
			fmt.Println()
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(0)
		}
		switch args[2] {
		case "-m":
			multiLine = true
		default:
			fmt.Println()
			fmt.Println("\033[31mError! Invalid Argument:\033[0m", args[2])
			fmt.Println()
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(0)
		}
	}

	if multiLine {
		if !isFilePath(args[0]) {
			fmt.Println()
			fmt.Println("\033[31mError! Invalid Filepath:\033[0m", args[0], "(multi-line mode should not be enabled for string input)")
			fmt.Println()
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(0)

		}
		extractArtSequenceStringFromArgs, err = readFile(args[0]) // txt file read
		if err != nil {
			fmt.Println("file Error")
			os.Exit(0)
		}
	} else if !multiLine && isFilePath(args[0]) {
		fmt.Println()
		fmt.Println("\033[31mError! Missing Argument:\033[0m multi-line mode must be enabled with filepath input")
		fmt.Println()
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	} else {
		extractArtSequenceStringFromArgs = args[0]
	}

	if decodeMode {

		containsBrackets := true

		splitSequenceAtRightBracket := strings.SplitAfter(extractArtSequenceStringFromArgs, "]")
		splitSequenceAtLeftBracket := strings.SplitAfter(extractArtSequenceStringFromArgs, "[")

		// splitAfter returns slice of length 1 if string does not contain seperator
		if len(splitSequenceAtRightBracket) == 1 && len(splitSequenceAtLeftBracket) == 1 {
			containsBrackets = false
		}

		if containsBrackets {

			// if length of last slice is zero, remove it.
			if len(splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1]) == 0 {
				splitSequenceAtRightBracket = splitSequenceAtRightBracket[:len(splitSequenceAtRightBracket)-1]
			}

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

			// if last string is an unbracketed symbol(s) ie. of length greater than zero, append to end of splitIntoBracketDataAndSingleData slice
			if len(splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1]) > 0 && splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1][0] != '[' {
				splitIntoBracketDataAndSingleData = append(splitIntoBracketDataAndSingleData, splitSequenceAtRightBracket[len(splitSequenceAtRightBracket)-1])
			}

			// check for uneven square brackets
			for _, data := range splitIntoBracketDataAndSingleData {
				if data[0] == '[' || data[len(data)-1] == ']' {
					if data[0] == '[' {
						if data[len(data)-1] != ']' {
							fmt.Println("\033[31mError! Unbalanced brackets:\033[0m", data)
							os.Exit(0)
						}
					} else if data[len(data)-1] == ']' {
						if data[0] != '[' {
							fmt.Println("\033[31mError! Unbalanced brackets:\033[0m", data)
							os.Exit(0)
						}
					}
				}
			}

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
							fmt.Println("\033[31mError! Format Error:\033[0m", data, " -> check input line:", newLineCount)
							os.Exit(0)
						} else {
							fmt.Println("\033[31mError! Format Error:\033[0m", data)
							os.Exit(0)
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
						fmt.Println()
						fmt.Println("\033[31mError! Format Error:\033[0m", extractedDigits)
						fmt.Println()
						fmt.Println("Usage:")
						flag.PrintDefaults()
						os.Exit(0)
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
		} else {
			// here logic if no brackets in input
			fmt.Println(extractArtSequenceStringFromArgs)
		}
	}

	if encodeMode { // nearly done - try and get rid of last newline + do readme + push to gitea / github
		var finalArtEncoded string
		var currentSymbol string
		splitStringFromArgs := strings.Split(extractArtSequenceStringFromArgs, "\n")
		for lineNum, line := range splitStringFromArgs {
			matchSymbol := false
			matchCount := 0
			for i := 0; i <= len(line)-1; i++ {
				if i == len(line)-1 {
					if line[i] == line[i-1] {
						matchSymbol = true
					} else {
						matchSymbol = false
					}
				} else if line[i] == line[i+1] {
					matchSymbol = true
				} else {
					matchSymbol = false
				}
				if matchSymbol {
					if matchSymbol && i == len(line)-1 {
						if lineNum < len(splitStringFromArgs)-1 {
							matchCount += 1
							currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]\n")
							finalArtEncoded += currentSymbol
							matchSymbol = false
							matchCount = 0
						} else {
							matchCount += 1
							currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")
							finalArtEncoded += currentSymbol
							matchSymbol = false
							matchCount = 0

						}
					} else {
						matchCount += 1
					}
				} else if !matchSymbol && matchCount > 0 {
					matchCount += 1
					currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")
					finalArtEncoded += currentSymbol
					matchCount = 0
				} else if !matchSymbol && matchCount == 0 {
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
}
