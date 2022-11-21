package main

import (
	"fmt"
	"os"
)

func main() {
	//var s,sep string
	for i, arg := range os.Args {
		fmt.Printf("%d_%v\n", i, arg)
	}

}
