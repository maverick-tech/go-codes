package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func main() {

	var slice = make([]int,0,3)
	var number string
	for {
		fmt.Printf("Enter an integer : \n")
		_, err := fmt.Scan(&number)
	 
		if err != nil {
		  break
		}
		if strings.Compare(number,"X")==0 || strings.Compare(number,"x")==0 {
			break
		} else {
			num,_ :=strconv.Atoi(number)
			slice = append(slice,num)
			sort.Ints(slice)
			fmt.Printf("Sorted Slice : %v\n",slice)
		}
	}
}
