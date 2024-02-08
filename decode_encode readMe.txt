Art->Text ENCODER / DECODER Text->Art 

Program will handle single and multiline art encoding or decoding.

Single line inputs must be a string within double quotes -> "string example".
Only three argument are required; . , string input, enable encoding/decoding mode.

Multi-line inputs must have four arguments; also enabling multi-line mode -> -m

Usage:
  -d    Enable decoding mode
    
  -e    Enable encoding mode
    
    
  -h    Example Usage:
        single-line decode mode: $ go run . "[5 #]T" -d
        single line encode mode: go run . "#####T" -e
        multi-line decode mode: $ go run . ./filepath -d -m
        multi-line encode mode: $ go run . ./filepath -e -m
    
  -m    Enable multi-line mode


Extra features:

Error Handling -
The program will handle all types of argument errors and return useful
error messages to assist the user, expanding on standard "Error" message.

Usage - 
Usage flags initiated to assist user