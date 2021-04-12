package main

import "fmt"

type Container struct{
    buffer int
    size int
}

type Record struct{
     identifier Value
     typeIdentifier Value
     block Block;
}
 

type Value struct {
     reference Container
     typeIdentifier string
}


type Block struct{
     identifier Value
     method  Value
     args Map
}

type Map struct {
     key Value
     value Value
}


func main() {
    var record Record
    record.identifier = Value{Container{},"module"}
    record.block.identifier = Value{Container{},"x"}
    record.block.method = Value{Container{},"Integer"}
    record.block.args.key = Value{Container{0,1},"0"}
    record.block.args.value = Value{Container{0,1},"0"}
}
