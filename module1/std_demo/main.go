package main

import (
	"fmt"
	"github.com/vivekprm/algo-in-go/module1"
)
// Many problem solve sites have you reading from stdin and stdout or files.
// stdin (standard input) is input that is provided by a user, often by typing from their keyboard, but this can also be piped in in most terminals.
// stdout (standard output) is where most programs will write output. In Go you write here by default if you do something like fmt.Println("hi").
// Piping files to stdin and stdout is also pretty easy:

// $ go run main.go < input.txt > output.txt

// This is all meant to just make it easier if you want to go practice on sites like Google Code Jam and others that tend to use stdin 
// and stdout for input/output of a program.

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module1.GCD(a, b)
		fmt.Println(gcd)
	}
}
