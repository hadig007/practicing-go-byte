package codes

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func BinToBytes() {
	f, err := os.Open("files/ethereum_evm_illustrated.pdf")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Printf("%s", hex.Dump(buf))
	}

}
