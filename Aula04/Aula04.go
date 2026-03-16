package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//s := "élite"
	//
	//fmt.Printf("%8T %[1]v %d\n", s, len(s))
	//fmt.Printf("%8T %[1]v\n", []rune(s))
	//
	//b := []byte(s)
	//fmt.Printf("%8T %[1]v %d\n", b, len(b))
	//
	//a := "a string"
	//a = strings.ToUpper(a)
	//fmt.Printf("%8T %[1]v %d\n", a, len(a))

	if len(os.Args) < 3 {
		fmt.Println(os.Stderr, "not enough arguments")
		os.Exit(-1)
	}

	old, nova := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, nova)

		fmt.Println(t)
	}
}
