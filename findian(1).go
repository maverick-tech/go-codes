package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	consoleReader := bufio.NewReader(os.Stdin)
    
	fmt.Println("Enter a String : ")
    str, _ := consoleReader.ReadString('\n')

	str = strings.ToLower(str) 
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str,"a"){

		fmt.Printf("Found!\n")

	} else {

		fmt.Printf("Not Found!\n")

	}
}
