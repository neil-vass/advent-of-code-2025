package main

import _ "embed"

//go:embed input.txt
var puzzleData string

func main() {
	println("Hello, World!")
	println(puzzleData)
}
