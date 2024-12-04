package main

import (
	"fmt"
	"os"
	"strings"
	// "regexp"
	// "strconv"
	// "strings"
)

var filename string = "content.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkIfXmas(page []string, line int, column int) int {
	value := 0
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			sens := sens{i, j}
			if checkBySens(page, line, column, sens) {
				value += 1
			}
		}
	}
	return value
}

func checkIfAmas(page []string, line int, column int) int {
	if checkAmasDiag(page, line, column, 1) && checkAmasDiag(page, line, column, -1) {
		return 1
	}
	return 0
}

func checkAmasDiag(page []string, line int, column int, sens int) bool {
	first_char := string(page[line+sens][column+1])
	second_char := string(page[line-sens][column-1])
	if (first_char == "M" && second_char == "S") || (first_char == "S" && second_char == "M") {
		return true
	}
	// return (first_char == "M" && second_char == "S") || (first_char == "S" && second_char == "M")
	return false
}

func checkBySens(page []string, line int, column int, sens sens) bool {
	totalLine := len(page)
	totlaColumn := len(page[0])
	if totalLine > line+3*sens.height && line+3*sens.height > -1 && totlaColumn > column+3*sens.width && column+3*sens.width > -1 {
		return string(page[line+sens.height][column+sens.width]) == "M" && string(page[line+2*sens.height][column+2*sens.width]) == "A" && string(page[line+3*sens.height][column+3*sens.width]) == "S"
	}
	return false
}

func extractPage(filename string) string {
	file, err := os.ReadFile(filename)
	check(err)
	content := string(file)
	return content
}

type sens struct {
	height int
	width  int
}

func ex1() {
	content := extractPage(filename)
	number := 0
	lines := strings.Split(content, "\n")

	for numberLine, line := range lines {
		for numberLettre, lettre := range line {
			char := string(lettre)
			if char == "X" {
				number += checkIfXmas(lines, numberLine, numberLettre)
			}
		}
	}
	fmt.Println(number)
}

func ex2() {
	content := extractPage(filename)
	number := 0
	lines := strings.Split(content, "\n")
	linesMiddle := lines[1 : len(lines)-1]

	for numberLine, line := range linesMiddle {
		lineMiddle := line[1 : len(line)-1]

		for numberLettre, lettre := range lineMiddle {
			char := string(lettre)
			if char == "A" {
				number += checkIfAmas(lines, numberLine+1, numberLettre+1)
			}
		}
	}
	fmt.Println(number)
}

func main() {
	ex2()
}
