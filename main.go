package main

import "fmt"
import "os"
import "strings"

func main() {
    file, _ := os.Open("test.map")
    b := make([]byte, 100)
    file.Read(b)
    fmt.Println(cap(b))
    fmt.Println(len(b))
    c := string(b)
    fmt.Println(c)
    fmt.Println(strings.Split(c, " "))
}
