package gohello

import "fmt"

func HelloWorld() string {
	return "gohello: Hello World"
}

func PrintWorldLn(txt string) {
	fmt.Printf("gohello: %s\n", txt)
}
