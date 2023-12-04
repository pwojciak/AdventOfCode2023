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
	var result int = 0
	var digits = "0123456789"
	var firstIndex = strings.IndexAny(line, digits)
	var lastInde = strings.LastIndexAny(line, digits)

	var stringResult = string(line[firstIndex]) + string(line[lastInde])
	result, err := strconv.Atoi(stringResult)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
