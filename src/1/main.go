package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
}

func getNumberFromLine(line string) int {

	if len([]rune(line)) < 1 {
		return 0
	}

	var result int = 0
	var digits = "0123456789"
	var firstIndex = strings.IndexAny(line, digits)
	var lastIndex = strings.LastIndexAny(line, digits)

	if firstIndex < 0 || lastIndex < 0 {
		return 0
	}

	var stringResult = string(line[firstIndex]) + string(line[lastIndex])
	result, err := strconv.Atoi(stringResult)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
