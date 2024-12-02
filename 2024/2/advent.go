package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename string = "content.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func remove(s []int, index int) []int {
	sprime := make([]int, len(s))
	copy(sprime, s)
	sprime = append(sprime[:index], sprime[index+1:]...)
	return sprime
}

func stringToIntegers(strings []string) []int {
	var tmp = []int{}
	for _, i := range strings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, j)
	}
	return tmp
}

func checkValidity(ints []int) bool {
	sens := 0
	if ints[0] > ints[1] {
		sens = -1
	} else if ints[0] < ints[1] {
		sens = 1
	}

	valid := true
	if sens == 1 {
		for i := range len(ints) - 1 {
			size := ints[i+1] - ints[i]
			valid = valid && ints[i] < ints[i+1] && 1 <= size && size <= 3
		}
		if valid {
			return valid
		}
	} else if sens == -1 {
		for i := range len(ints) - 1 {
			size := ints[i] - ints[i+1]
			valid = valid && ints[i] > ints[i+1] && 1 <= size && size <= 3
		}
		if valid {
			return valid
		}
	}
	return false
}

func ex1() {
	file, err := os.ReadFile(filename)
	check(err)

	content := string(file)
	lines := strings.Split(content, "\n")
	value := 0

	for _, line := range lines {
		elements := strings.Split(line, " ")
		ints := stringToIntegers(elements)
		if checkValidity(ints) {
			value += 1
		}
	}
	fmt.Println(value)
}

func ex2() {
	file, err := os.ReadFile(filename)
	check(err)

	content := string(file)
	lines := strings.Split(content, "\n")
	value := 0

	for _, line := range lines {
		elements := strings.Split(line, " ")
		ints := stringToIntegers(elements)
		valid := checkValidity(ints)
		for j := range ints {
			test := checkValidity(remove(ints, j))
			valid = valid || test

		}
		if valid {
			value += 1
		}
	}
	fmt.Println(value)
}

func main() {
	ex2()
}
