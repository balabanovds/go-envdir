package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func collect(dir string) []string {
	var envs []string

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist", dir)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var b strings.Builder
		b.WriteString(file.Name())
		file.Name()
	}

	return envs
}

// readDir find files recursively
func readFile(file string, envs *[]string) {

	return nil
}
