package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() int {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0

	ar := strings.Split(string(content), "\n")
	for _, c := range ar {

		s := string(c)
		x := firstDigit(s, false)
		Reverse(&s)
		y := firstDigit(s, true)
		sum += (x * 10) + y

	}
	return sum
}

func part2() int {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0

	ar := strings.Split(string(content), "\n")
	for _, c := range ar {
		digits := []string{
			"zero",
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
		}

		s := string(c)
		for i := range digits {
			s = strings.Replace(s, digits[i], string(digits[i][0])+strconv.Itoa(i)+string(digits[i][len(digits[i])-1]), -1)
		}
		x := firstDigit(s, false)

		Reverse(&s)
		y := firstDigit(s, true)
		sum += (x * 10) + y

	}

	return sum
}

func firstDigit(s string, reverse bool) int {
	for _, c := range s {
		if strings.Contains("0123456789", string(c)) {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				// ... handle error
				panic(err)
			}
			return i
		}
	}
	return 0
}

func Reverse(s *string) (result string) {
	for _, v := range *s {
		result = string(v) + result
	}
	*s = result
	return
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
