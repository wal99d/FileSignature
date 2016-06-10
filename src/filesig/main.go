package main

import (
        "fmt"
        "os"
        "flag"
)

func main() {
        //Here we take file from 1st argument
        var file string
        var lastByte int
        flag.StringVar(&file, "f","","filesig -f FILE_NAME")
        flag.IntVar(&lastByte, "b", 0, "filesig -b NUMBER_OF_BYTES_TO_READ")
        flag.Parse()
        //then we open this file
        f, _ := os.Open(file)
        //we tell GO to wait tell this func "main" finishes execution then close the file
        defer f.Close()
        //we must read file state in order to get the size of the file
        fileStat, _ := os.Stat(file)
        //then we take the file size
        fileSize := fileStat.Size()
        //here we create byte array to store all file contents in it
        bytes := make([]byte, int(fileSize))
        //we create offset varaible to use it to track how bytes we read from our file
        var offset int64
        offset = 0
        //we use Seek func from os library very usfel to seek/jump within file contents
        _, _ = f.Seek(offset, 0)
        //this is the main Read func to store all file content into bytes array, if there is an error print it out
        if _, err := f.Read(bytes); err != nil {
                fmt.Println(err)
        }
        //then we print final 4 bytes using very handy tool which is slice
        fmt.Printf("% x\n", bytes[0:lastByte])
}
