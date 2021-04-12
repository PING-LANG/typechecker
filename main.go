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
 

type Case struct {
     identifier string
     typeIdentifier string

}

type Block struct{
     identifier string
     method  string
     args []Case
}



func main() {
    record Record
    record.identifier = "module"
    record.Block.identifier = "x"
    append(record.Block.args, Case{})
}
