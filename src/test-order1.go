package main

import "fmt"

type IntSlice []int
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	src := IntSlice{12, 2, 29, 0x28, 2934, 11, -2}

	fmt.Println(src)

	length := src.Len()
	for i:=0; i< length; i++{
		for j:=0; j< length; j++{

			if( src.Less(i, j)){
				src.Swap(i,j)
			}

		}


	}

	fmt.Println(src)


}
