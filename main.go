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

type Variable struct {
    container Container
    record Record
}

type Expression struct {
     value1 Record
     value2 Record
}


func main() {
  
}
