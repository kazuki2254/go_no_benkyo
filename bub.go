//bubble sort

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Stal interface {
	sort() string
}

type SortB struct {
	Data []int
	Cntl int
}

type SortI struct {
	Data []int
	Cntl int
}

func sortingalgo(s Stal) string {
	return s.sort()
}

func (s SortI) sort() string {
	return fmt.Sprintf("\nSortI.sort()\nIt is a dummy\nresult:\t%v\n", s)
}

func (s SortB) sort() string {
	/*
		c:=true
		for c {
			c=false
			for n:=len(s.Data)-1;n>0;n--{
				if(s.Data[n-1])>s.Data[n] {
					s.Data[n-1],s.Data[n]=s.Data[n],s.Data[n-1]
					c=true
				}
				s.Cntl++
			}
		}
		return fmt.Sprintf("\n\n--methot srot()--\nresult:\t%v\nLoop:\t%d\n\n",s.Data,s.Cntl)
	*/

	return fmt.Sprintf("\nSortB.sort()--\n\n%v", sortBub(s.Data, s.Cntl))
}

func sortBub(o []int, cc int) string {
	c := true
	for c {
		c = false
		for n := len(o) - 1; n > 0; n-- {
			if o[n-1] > o[n] {
				o[n-1], o[n] = o[n], o[n-1]
				c = true
				//			fmt.Println(o)
			}
			cc++
		}

	}
	return fmt.Sprintf("--func SortBub()--\nresult:\t%v\nLoop:%d\n", o, cc)

}

func main() {
	o := []int{}
	c := 0
	for _, n := range os.Args[1:] {
		v, _ := strconv.Atoi(n)
		o = append(o, v)
	}
	var s SortB
	var d SortI
	d.Data = o
	s.Data = o

	fmt.Println("moto:", "\t", o)

	fmt.Println(sortingalgo(d))

	fmt.Println(sortingalgo(s))

	fmt.Println("main()")
	fmt.Println(sortBub(o, c))

}
