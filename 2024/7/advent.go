package main

import (
	"fmt"
	"os"
	"strings"
	// "regexp"
	"strconv"
	// "strings"
)

var filename string = "content.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func string_to_int(el string) int {
	i, err := strconv.Atoi(el)
	if err != nil {
		panic(err)
	}
	return i
}

func extractPage(filename string) string {
	file, err := os.ReadFile(filename)
	check(err)
	content := string(file)
	return content
}

func checkIfValid(target int, current int, nextInts []int) bool {
	if current > target {
		return false
	}
	if len(nextInts) == 0 {
		return current == target
	}
	nextPlus := checkIfValid(target, current+nextInts[0], nextInts[1:])
	nextMult := checkIfValid(target, current*nextInts[0], nextInts[1:])
	return nextMult || nextPlus
}

func checkIfValid2(target int, current int, nextInts []int) bool {
	if current > target {
		return false
	}
	if len(nextInts) == 0 {
		return current == target
	}
	nextPlus := checkIfValid2(target, current+nextInts[0], nextInts[1:])
	nextMult := checkIfValid2(target, current*nextInts[0], nextInts[1:])
	nextConcat := checkIfValid2(target, concatInt(current, nextInts[0]), nextInts[1:])
	return nextMult || nextPlus || nextConcat
}

func concatInt(first int, second int) int {
	string_first := strconv.Itoa(first)
	string_second := strconv.Itoa(second)
	third := string_first + string_second
	third_int := string_to_int(third)
	return third_int
}

func ex1() {
	content := extractPage(filename)
	lines := strings.Split(content, "\n")
	count := 0
	for _, line := range lines {
		blocks := strings.Split(line, ": ")
		target := string_to_int(blocks[0])
		ints_string := strings.Split(blocks[1], " ")
		ints := make([]int, 0)
		for _, el := range ints_string {
			currentInt := string_to_int(el)
			ints = append(ints, currentInt)
		}
		isValid := checkIfValid(target, ints[0], ints[1:])
		if isValid {
			count += target
		}
	}
	fmt.Println(count)
}

func ex2() {
	content := extractPage(filename)
	lines := strings.Split(content, "\n")
	count := 0
	for _, line := range lines {
		blocks := strings.Split(line, ": ")
		target := string_to_int(blocks[0])
		ints_string := strings.Split(blocks[1], " ")
		ints := make([]int, 0)
		for _, el := range ints_string {
			currentInt := string_to_int(el)
			ints = append(ints, currentInt)
		}
		isValid := checkIfValid2(target, ints[0], ints[1:])
		if isValid {
			count += target
			fmt.Println(target)
		}
	}
	fmt.Println(count)
}

func main() {
	ex2()
}
