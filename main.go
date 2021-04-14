package main

import "fmt"



type GraphVertice struct {
    vertice List
}

type GraphEdge struct {
    vertice List
    edge GraphVertice;
}

type Graph struct {
    vertices []GraphVertice
    edges []GraphEdge
}

type Unit struct {
    present Quantifier
    args List
}

type Function struct {
    sum Unit
    leaf Unit
}

type List struct {
    key Function
    value Function
}


func main() {
    
}
