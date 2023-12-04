package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func listOfNumbers(str string) []int {

	numberList := []int{}
	curNumber := ""
	digits := "1234567890"

	for _, c := range str {
		if strings.Contains(digits, string(c)) {
			curNumber += string(c)
		} else {
			if curNumber != "" {
				x, err := strconv.Atoi(curNumber)
				if err != nil {
					fmt.Println("err")
				}
				numberList = append(numberList, x)
				curNumber = ""
			}
		}
	}
	if curNumber != "" {
		x, err := strconv.Atoi(curNumber)
		if err != nil {
			fmt.Println("err")
		}
		numberList = append(numberList, x)
	}
	return numberList
}

func part2() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0
	mini_sum := 0

	ar := strings.Split(string(content), "\n")
	scores := []int{}

	for _, c := range ar {
		if c == "" {
			continue
		}
		split := strings.Split(string(c), "|")
		winString := split[0]
		myNumbers := listOfNumbers(split[1])

		for _, i := range myNumbers {

			if strings.Contains(winString, " "+strconv.Itoa(i)+" ") {
				mini_sum++
			}
		}
		sum += mini_sum
		scores = append(scores, mini_sum)
		mini_sum = 0

	}

	amounts := []int{}
	for range scores {
		amounts = append(amounts, 1)
	}
	for i, e := range scores {
		n := i + e + 1
		m := i + 1
		for i2 := m; i2 < n; i2++ {
			amounts[i2] = amounts[i2] + (amounts[i])
		}
	}
	sum = 0
	for _, e := range amounts {
		sum += e
	}

	fmt.Println(sum)
}

func part1() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0
	mini_sum := 0

	ar := strings.Split(string(content), "\n")

	for _, c := range ar {
		if c == "" {
			break
		}
		split := strings.Split(string(c), "|")
		winString := split[0]

		myNumbers := listOfNumbers(split[1])

		for _, i := range myNumbers {

			if strings.Contains(winString, " "+strconv.Itoa(i)+" ") {
				if mini_sum == 0 {
					mini_sum = 1
				} else {
					mini_sum *= 2
				}
			}
		}
		sum += mini_sum
		mini_sum = 0

	}

	fmt.Println(sum)
}
func main() {
	part1()
	part2()
}
