package main

import (
        "fmt"
        "os"
)

func main() {
        file := os.Args[1]
        f, _ := os.Open(file)
        defer f.Close()
        fileStat, _ := os.Stat(file)
        fileSize := fileStat.Size()
        bytes := make([]byte, int(fileSize))
        var offset int64
        offset = 0
        _, _ = f.Seek(offset, 0)
        if _, err := f.Read(bytes); err != nil {
                fmt.Println(err)
        }
        fmt.Printf("% x\n", bytes[0:4])
}
