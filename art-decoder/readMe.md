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

        [8  ]@|\[2 @]
        [7  ]-[2  ][4 @]
        [6  ]/7[3  ][4 @]
        [5  ]/[4  ][6 @]
        [5  ]\-' [8 @]`-[15 _]
        [6  ]-[9 @][13  ]/[4  ]\
         [7 _]/[4  ]/_[7  ][6 _]/[6  ]|[10 _]-
        /,[10 _]/[2  ]`-.[3 _]/,[13 _][10 -]_)