#File Signature Extractor (Written in Go)
####This Go app will extract the file Signature from file passed as an argument, and then it validates it from in-memory list of file signatures  
like below command:
```
#bin/filesig -f FILE_GOES_HERE -b NUMBER_OF_BYTES_TO_READ"
```
***for example***
*To extract only 4 Bytes from a PDF file you use the below command:*

```
#bin/filesig -f /location/SomeFile.pdf -b 4
The Signature of this file is: 25 50 44 46
File idenfied and it's: PDF|FDF
```


#Todo
---
* Need to implement reading bytes from some value to end value
