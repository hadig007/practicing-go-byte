package codes

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFromTxtFile() {
	// read .txt file to byte
	content, err := ioutil.ReadFile("./files/name.txt")
	if err != nil {
		log.Fatal(err)
	}
	// print as bytes
	fmt.Println(content)
	println("===============")
	// print as string text
	fmt.Println(string(content))
}
