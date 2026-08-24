package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		read, e := io.ReadAll(os.Stdin)
		if e != nil {
			fmt.Print("ERROR: ", e)
			return
		}
		fmt.Print(string(read))
	}

	if len(args) >= 2 {
		first_file := os.Args[1:]
		for _, file := range first_file {
			content, e := os.Open(file)
			if e != nil {
				fmt.Print("ERROR: ", e)
				continue
			}

			content_bytes, e := io.ReadAll(content)

			if e != nil {
				fmt.Print("ERROR: ", e)
				continue
			}
			content.Close()
			fmt.Print(string(content_bytes))
		}

	}
}
