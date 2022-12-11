package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// read .txt file to byte
	content, err := ioutil.ReadFile("name.txt")
	if err != nil {
		log.Fatal(err)
	}
	// print as bytes
	fmt.Println(content)
	println("===============")
	// print as string text
	fmt.Println(string(content))
}
