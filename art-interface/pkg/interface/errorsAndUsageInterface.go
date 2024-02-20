package art_interf

import (
	"flag"
	"fmt"
	"os"
)

var (
	ENCODE_TEXT               = "Enable \033[33mencoding\033[0m mode\n\n"
	DECODE_TEXT               = "Enable \033[33mdecoding\033[0m mode\n"
	MULTILINE_TEXT            = "Enable \033[33mmulti-line\033[0m mode\n"
	USAGE_TEXT                = "\033[33mExample Usage:\033[0m\nsingle-line decode mode: \033[33m$ go run . \"[5 #]T\" -d\033[0m\nsingle line encode mode: \033[33m$ go run . \"#####T\" -e\033[0m\nmulti-line decode mode: \033[33m$ go run . ./filepath -d -m\033[0m\nmulti-line encode mode: \033[33m$ go run . ./filepath -e -m\033[0m\n"
	MISSING_ARG               = "\033[31mError! Missing Argument:\033[0m Minimum 3 arguments"
	EXTRA_ARG                 = "\033[31mError! Extra Argument:\033[0m"
	INVALID_ARG               = "\033[31mError! Invalid Argument:\033[0m"
	INVALID_FILEPATH          = "\033[31mError! Invalid Filepath\033[0m"
	ERROR_READFILE            = "\033[31mError! Error reading file\033[0m"
	MISSING_ARG_FOR_MULTILINE = "\033[31mError! Missing Argument:\033[0m multi-line mode must be enabled with filepath input"
	UNBALANCED_BRACKETS       = "\033[31mError! Unbalanced brackets\033[0m"
	FORMAT_ERROR              = "\033[31mError! Format Error:\033[0m"
)

// function to read file from path argument
func ReadFile(filePath string) (string, error) {
	// Read the entire file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	return string(content), nil
}

// check argument is valid filepath
func IsFilePath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PrintError(errorText string, errorData ...string) {
	fmt.Println()
	fmt.Println(errorText, errorData)
	fmt.Println()
	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(0)
}

func PrintUsage() {
	fmt.Println()
	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(0)
}
