package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strings"
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

func checkDo(list []byte, enabled bool) bool {
	if enabled {
		dont_string := string(list)
		return dont_string == "don't()"
	} else {
		do_string := string(list)
		return do_string == "do()"

	}
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

	// content := string(file)
	regex, err := regexp.Compile(`mul\([0-9]*,[0-9]*\)`)

	check(err)

	list := regex.FindAll(file, -1)
	value := 0

	for i := range list {
		first := 0
		second := 0
		first_finished := false
		for j := range list[i] {
			char := int(list[i][j])
			if 48 <= char && char <= 57 {
				if first_finished {
					second = (char - 48) + second*10
				} else {
					first = (char - 48) + first*10
				}
			} else if char == 44 {
				first_finished = true
			}

		}
		fmt.Println(first, second)
		value += first * second
	}
	fmt.Println(value)
}

func ex2() {
	file, err := os.ReadFile(filename)
	check(err)

	// content := string(file)
	regex, err := regexp.Compile(`mul\([0-9]*,[0-9]*\)|do\(\)|don\'t\(\)`)

	check(err)

	list := regex.FindAll(file, -1)
	value := 0
	enabled := true

	for i := range list {
		first := 0
		second := 0
		first_finished := false

		for j := range list[i] {
			char := int(list[i][j])
			if 48 <= char && char <= 57 {
				if first_finished {
					second = (char - 48) + second*10
				} else {
					first = (char - 48) + first*10
				}
			} else if char == 44 {
				first_finished = true
			} else if char == 100 {
				if checkDo(list[i], enabled) {
					enabled = !enabled
				}
			}
		}
		if enabled {
			value += first * second
		}
	}
	fmt.Println(value)
}

func main() {
	ex2()
}
