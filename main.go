package main

import "fmt"

type Container struct{
    identifier int
    size int
}

type Record struct{
     identifier string
     typeIdentifier string
     block Block;
}
 

type Value struct {
     typeIdentifier string

}

type Block struct{
     identifier string
     method  string
     args []Value
}

type Map struct {
     key Value
     value Value
}


func main() {
    var record Record
    record.identifier = "module"
    record.block.identifier = "x"
    append(record.block.args, Value{"int"})
}
