package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	m["zero"] = 0
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	list := strings.Split(string(data), "\n")
	// fmt.Println(list[1])
	for _, item := range list {
		fmt.Println(item)
		numbers := []int{}
		num_strings := ""
		if item == "" {
			continue
		}
		for _, s := range item {

			n, err := strconv.Atoi(string(s))
			if err != nil {
				num_strings += string(s)
				// iterate over the maps keys
				for key, value := range m {
					// if the value in the map is equal to the number string
					is_in := strings.Contains(num_strings, key)
					if is_in {
						fmt.Println(is_in)
						numbers = append(numbers, value)
						num_strings = string(s)
					}
				}
				continue
			}
			numbers = append(numbers, n)

		}
		fmt.Println(numbers)
		combi := strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1])
		fmt.Println(combi)
		int_combi, _ := strconv.Atoi(combi)
		result += int_combi
	}
	fmt.Println(result)
}
