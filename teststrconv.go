//%v

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "3"
	b := "3"
	c, _ := strconv.ParseFloat(a, 64)
	d, _ := strconv.ParseFloat(b, 64)
	fmt.Printf("%v\n", c+d)
	fmt.Printf("%v\n", d)
}
