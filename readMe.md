# <p style="text-align: center;">Art->Text ENCODER / DECODER Text->Art</p>

    

<p style="text-align: center;">Program will handle single and multiline text-art encoding or decoding.</p> 



## Usage:

### $ go run . -h

    -d    Enable decoding mode
    
    -e    Enable encoding mode

    -h    Example Usage:
          single-line decode mode: $ go run . "[5 #]T" -d
          single line encode mode: $ go run . "#####T" -e
          multi-line decode mode: $ go run . ./filepath -d -m
          multi-line encode mode: $ go run . ./filepath -e -m
    
    -m    Enable multi-line mode

## Test Examples:

### Input: 
      $ go run . "[5 #]T" -d  

### Output:
      #####T

### Input:
      $ go run . ./lion.encoded.txt -d -m

### Output: 

  ![Reference Image](screenshots/Screenshot%202024-02-15%20at%2010.30.00.png)

### Input:
      $ go run . ./lion.art.txt -e -m

### Output: 

  ![Reference Image](screenshots/Screenshot%202024-02-15%20at%2010.37.25.png)        