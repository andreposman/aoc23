package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := readInput()
	fmt.Printf("%v", input)
}

func readInput() ([]string, error) {
	fn := "input.txt"
	readFile, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}

	return fileLines, err
}
