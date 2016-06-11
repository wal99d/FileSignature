# File Signature Extractor and Verifier
This Go app will extract the file Signature from file passed as an argument, and then it validates it from in-memory list of file signatures  
like below command:
```
#bin/filesig -f FILE_GOES_HERE -o START_OFFSET -b NUMBER_OF_BYTES_TO_READ"
```
***for example***
*To extract only 4 Bytes from a PDF file indicated by "-b" starting from offset 0 indicated by "-o" you use the below command:*

```
#bin/filesig -f /location/SomeFile.pdf -o 0 -b 4
The Signature of this file is: 25 50 44 46
File idenfied and it's: PDF|FDF
```
currently the in-memory db contains 942 file signatures

## How-to

You have two options:

1. By running docker and downloading "wal99d/filesig:0.1" image and sharing folder from your host to docker container all in one command as below:

`#docker run -it -v /location/folder/files:/host wal99d/filesig:0.1 /home/filesig`

 and then run below command once you successfuly logged in docker image

 `#bin/filesig -f /host/SomeFile.pdf -o 0 -b 4`

 if everything works fine then you will get below result

 ```
 The Signature of this file is: 25 50 44 46
 File idenfied and it's: PDF|FDF
 ```
2. By installing [Go](https://golang.org/dl/), cloning, and exporting GOPATH as shown below:

```
#git clone https://github.com/wal99d/FileSignature.git
#cd FileSignature
#export GOPATH=`pwd`
#export PATH=$PATH:$GOPATH
```

 finally you need to comple Go project using below command:

`#go install filesig`


-
***inspired by [Gary C. Kessler.](http://www.garykessler.net/library/file_sigs.html)***
