package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	// "regexp"
	"strconv"
	// "strings"
	"slices"
)

var filename string = "content.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractPage(filename string) string {
	file, err := os.ReadFile(filename)
	check(err)
	content := string(file)
	return content
}

func string_to_int(el string) int {
	fmt.Println(el)
	// string to int
	i, err := strconv.Atoi(el)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

type condition struct {
	before string
	after  string
}

func ex1() {
	content := extractPage(filename)
	blocks := strings.Split(content, "\n\n")
	conditions := blocks[0]
	pages := blocks[1]
	conditions_string_list := strings.Split(conditions, "\n")
	condition_num := len(conditions_string_list)
	condition_list := make([]condition, condition_num)
	for rank, condition_string := range conditions_string_list {
		split := strings.Split(condition_string, "|")
		condition_list[rank] = condition{split[0], split[1]}
	}
	count := 0
	pages_list := strings.Split(pages, "\n")
	for _, pages := range pages_list {
		numbers_strings := strings.Split(pages, ",")
		page_map := make(map[string]int)
		for rank, page := range numbers_strings {
			page_map[page] = rank
		}
		isValid := true
		for _, condition := range condition_list {
			rank_before, is_after := page_map[condition.before]
			rank_after, is_before := page_map[condition.after]
			if is_after && is_before {
				isValid = isValid && rank_after >= rank_before
			}
		}
		if isValid {
			len := len(numbers_strings)
			middle := string_to_int(numbers_strings[len/2])
			count += middle
		}
	}
	fmt.Println(count)
}

func order_disorderly(list []string, ordered_list []string) []string {
	fmt.Println(list, ordered_list)
	output := make([]string, len(list))
	rank := 0
	for _, value := range ordered_list {
		if slices.Contains(list, value) {
			output[rank] = value
			rank += 1
		}
	}
	return output
}

func ex2() {
	content := extractPage(filename)
	blocks := strings.Split(content, "\n\n")
	conditions := blocks[0]
	pages := blocks[1]
	conditions_string_list := strings.Split(conditions, "\n")
	condition_num := len(conditions_string_list)
	condition_list := make([]condition, condition_num)
	for rank, condition_string := range conditions_string_list {
		split := strings.Split(condition_string, "|")
		condition_list[rank] = condition{split[0], split[1]}
	}
	count := 0
	pages_list := strings.Split(pages, "\n")

	for _, pages := range pages_list {
		numbers_strings := strings.Split(pages, ",")
		page_map := make(map[string]int)
		for rank, page := range numbers_strings {
			page_map[page] = rank
		}
		isValid := true
		for _, condition := range condition_list {
			rank_before, is_after := page_map[condition.before]
			rank_after, is_before := page_map[condition.after]
			if is_after && is_before {
				isValid = isValid && rank_after >= rank_before
			}
		}
		if !isValid {
			occurences := make(map[string]int)
			number_of_number := int(math.Round(math.Sqrt(float64(len(condition_list))))) + 2
			ordered_list := make([]string, number_of_number)
			for _, condition := range condition_list {
				after := condition.after
				before := condition.before
				if slices.Contains(numbers_strings, before) && slices.Contains(numbers_strings, after) {
					_, ok := occurences[after]
					if ok {
						occurences[after] += 1
					} else {
						occurences[after] = 1
					}
					_, ok_before := occurences[before]
					if !ok_before {
						occurences[before] = 0
					}
				}
			}
			fmt.Println(occurences)
			fmt.Println(len(occurences))
			for number, occurence := range occurences {
				ordered_list[occurence] = number
			}
			len := len(numbers_strings)
			ordered := order_disorderly(numbers_strings, ordered_list)
			fmt.Println(ordered)
			middle := string_to_int(ordered[len/2])
			count += middle
		}
	}

	fmt.Println(count)
}

func main() {
	ex2()
}
