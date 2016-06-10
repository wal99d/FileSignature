# File Signature Extractor and Verifier
This Go app will extract the file Signature from file passed as an argument, and then it validates it from in-memory list of file signatures  
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
currently the in-memory db contains 942 file signatures

## How-to Start

You have two options:

1. By running docker and downloading "wal99d/filesig:0.1" image in one command as below:

`#docker run -it -v /location/folder/files:/host wal99d/filesig:0.1 /home/filesig`

 and then run below command once you successfuly logged in docker image

 `#bin/filesig -f /host/SomeFile.pdf -b 4`

 if everything works fine then you will get below result

 ```
 The Signature of this file is: 25 50 44 46
 File idenfied and it's: PDF|FDF
 ```

2. By installing Go and exporting GOPATH then compling Go project using below command:

`#go install filesig`


## Todo

* Need to implement reading bytes from some value to end value

-
***inspired by [Gary C. Kessler.](http://www.garykessler.net/library/file_sigs.html)***
