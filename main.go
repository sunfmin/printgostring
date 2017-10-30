package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err, os.Args[1])
		return
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#+v\n", buf.String())
}
