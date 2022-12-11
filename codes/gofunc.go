package codes

import (
	"bytes"
	"fmt"
)

func GoFunctions() {
	// contains
	data1 := []byte{102, 79, 99, 111, 110} // falcon
	data2 := []byte{111, 110}              //on

	if bytes.Contains(data1, data2) {
		fmt.Println("Contains")
	} else {
		fmt.Println("Does not contains")
	}

	// equal
	if bytes.Equal([]byte("falcon"), []byte("falcon")) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	data3 := []byte{111, 119, 108, 9, 99, 97, 116, 32, 32, 32, 32, 100, 111,
		103, 32, 112, 105, 103, 32, 32, 32, 32, 98, 101, 97, 114}

	fields := bytes.Fields(data3)

	fmt.Println(fields)
	for _, e := range fields {
		fmt.Printf("%s ", string(e))
	}
}
