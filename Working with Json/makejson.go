package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	mymap := make(map[string]string)
	var name string
	var address string
	
	fmt.Printf("Enter Name : \n")
	_, err := fmt.Scan(&name)
 
	if err != nil {
	  fmt.Printf("Something Went Wrong!\n")
	} else {
		mymap["name"]=name
	}
	
	fmt.Printf("Enter address : \n")
	_, err2 := fmt.Scan(&address)
 
	if err2 != nil {
	  fmt.Printf("Something Went Wrong!\n")
	} else {
		mymap["address"]=address
	}
	
	//creating json
	myjson,_ := json.Marshal(mymap)

	//printing json
	fmt.Printf("created json is\n"+string(myjson))
}
