package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
)

func cli(reader io.Reader, writer io.Writer) bool {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if string(scanner.Bytes()) == ".exit" {
		return false
	}
	writer.Write([]byte("Unrecognized command: "))
	writer.Write(scanner.Bytes())
	writer.Write([]byte("\n"))
	return true
}

func main() {
	for {
		fmt.Print("db>")
		if !cli(os.Stdin, os.Stdout) {
			return
		}
	}
}
