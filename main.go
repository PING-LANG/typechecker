package main

import "fmt"

type Container struct{
    identifier int
    size int
}

type Record struct{
    identifier int
    integer int
    float float32
    double float64
    str    string
}

type Case struct{
     record Record
     token Token;
}

func main() {
    fmt.Println("hello");  
}
