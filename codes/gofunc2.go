package codes

import (
	"bytes"
	"fmt"
)

func GoFunctions2() {
	data := [][]byte{[]byte("an"), []byte("old"), []byte("wolf")}
	joined := bytes.Join(data, []byte(" "))

	fmt.Println(data)
	fmt.Println(joined)
	fmt.Println(string(joined))

	fmt.Println("--------------------------")

	data2 := []byte{102, 97, 108, 99, 111, 110, 32}
	data3 := bytes.Repeat(data2, 3)

	fmt.Println(data3)
	fmt.Println(string(data3))

	fmt.Println("--------------------------")

	data4 := []byte{32, 32, 102, 97, 108, 99, 111, 110, 32, 32, 32}
	data5 := bytes.Trim(data4, " ")

	fmt.Println(data5)
	fmt.Println(string(data5))
}
