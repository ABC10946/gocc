package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "引数の個数が正しくありません\n")
		os.Exit(1)
	}

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")
	fmt.Println("  mov rax,", os.Args[1])
	fmt.Println("  ret")
}