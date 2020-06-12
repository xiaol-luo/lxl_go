package main

// import 方式
import "flag"

import (
	"os"
)

const (
	Space = " "
	NewLine = "\n"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

func main() {
	flag.Parse()
	var outStr string = ""
	for i := 0; i < flag.NArg(); i ++ {
		// fmt.Printf("flag %d is %s\n", i , flag.Arg(i))
		if i > 0 {
			outStr += Space
		}
		outStr += flag.Arg(i)
	}
	if !*omitNewline {
		outStr += NewLine
	}
	os.Stdout.WriteString("os print" + outStr)
	// fmt.Print("fmt print" + testStr)
}
