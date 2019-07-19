package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must specify at least env_dir and command")
		os.Exit(1)
	}
	path := os.Args[1]
	cmd := os.Args[2:]

	fmt.Println(cmd)

	envSlice := collect(path)
	fmt.Println(envSlice)
}
