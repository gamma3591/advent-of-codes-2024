package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var filename string = "content.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func spreadNumbers(numbers []string) ([]int, []int) {
	length := len(numbers) / 2
	left := make([]int, length)
	right := make([]int, length)

	for i := 0; i < length; i++ {
		leftNum, err := strconv.Atoi(numbers[2*i])
		check(err)
		rightNum, err := strconv.Atoi(numbers[2*i+1])
		check(err)

		left[i] = leftNum
		right[i] = rightNum
	}
	return left, right
}

func printUniqueValue(arr []int) map[int]int {
	//Create a   dictionary of values for each element
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func ex1() {
	file, err := os.ReadFile(filename)
	check(err)

	content := string(file)
	numbers := strings.Fields(content)
	left, right := spreadNumbers(numbers)
	sort.Ints(left)
	sort.Ints(right)
	value := 0

	fmt.Println(left)
	fmt.Println(right)

	for i := range left {
		value += abs(left[i], right[i])
		fmt.Println(abs(left[i], right[i]))
	}

	fmt.Println(value)
}

func ex2() {
	file, err := os.ReadFile(filename)
	check(err)

	content := string(file)
	numbers := strings.Fields(content)
	length := len(numbers) / 2
	left, right := spreadNumbers(numbers)

	sort.Ints(right)

	value := 0

	dict := printUniqueValue(right)
	for i := 0; i < length; i++ {
		if dict[left[i]] > 0 {
			value += left[i] * dict[left[i]]
		}
	}
	fmt.Println(value)

}

func main() {
	ex1()
}
