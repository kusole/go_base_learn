package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// panic defer recover

func first() {
	fmt.Println("first func run")
}

func second() {
	fmt.Println("second func run")
}

func main() {
	defer second()
	first()
	fname := "test.txt"
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		os.Exit(1)
	}
	bReader := bufio.NewReader(f)
	for {
		line, ok := bReader.ReadString('\n')
		if ok != nil {
			break
		}
		fmt.Println(strings.Trim(line, "\r\n"))
	}

	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	fmt.Println("i am warking and singing...")
	panic("it starts to rain cats and dogs")
}
