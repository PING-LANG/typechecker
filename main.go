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

type Unit struct {
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

//module type
//abstract module
//design module
//code module

type Module struct {
    moduleType int   
}

type Result struct {
     call result
     value Value
}

type Variant struct {
    result Result
    ttype Type
}

type Proof struct {
    sum Unit
    leaf Unit
}

type Type struct {
    dependentType Proof
    result proof
}

type List struct {
     key Type
     value Type
}

type Graph struct {
    edge Proof;
    vertice List;
}

func main() {
    
}
