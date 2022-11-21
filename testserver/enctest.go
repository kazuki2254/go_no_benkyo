package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]int)
	m["hoge"] = 4
	m["fuga"] = 5

	i := 4
	ii := 5
	o := "+"

	hoge := strconv.Itoa(i) + o + strconv.Itoa(ii)
	m[hoge] = 6
	j, _ := json.Marshal(m)
	fmt.Println(m)
	fmt.Printf("%s\t%T", j, j)
}
