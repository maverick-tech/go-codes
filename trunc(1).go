package main

import (
	"fmt"
)

func main() {

	var f float64


	fmt.Println("Enter a float value : ")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Printf("Wrong Value\n")
	} else {
		fmt.Printf("Truncated value is %d", int(f))
	}
}
