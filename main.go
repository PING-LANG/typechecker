package main

import "fmt"

type Container struct{
    buffer int
    size int
}

type Value struct {
     reference Container
     typeIdentifier string
}

type Error struct {
     reason Value
     solution Value
}

type Declaration struct {
    typeIdentifier Value
    name Value
}

type MethodCall struct {
    identifier Value
    method Value
    args List
}

type Case struct {
    literal Value
     
}

type Loop struct {
    literal Result
}

type Module struct {
   
}

func main() {
    
}
