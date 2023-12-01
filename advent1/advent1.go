package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	// fmt.Println(string(content))
	sum := 0

	ar := strings.Split(string(content), "\n")
	for _, c := range ar {
		digits := [10]string{
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
		fmt.Println(s)
		x := firstDigit(s, false)

		fmt.Println(x, " x")
		s = Reverse(s)
		y := firstDigit(s, true)
		// fmt.Print(s)
		// fmt.Println(x)
		// fmt.Print(y)
		sum += (x * 10) + y

	}

	fmt.Println(sum)
}

func firstDigit(s string, reverse bool) int {

	for _, c := range s {
		// fmt.Printf("%d %s\n", i, c)
		// fmt.Printf(string(c) + " c")

		// if len(stringNumberIndices) > 0 {
		// 	if index > stringNumberIndices[0] {
		// 		return stringNumberValues[0]
		// 	}
		// }
		if strings.Contains("0123456789", string(c)) {
			fmt.Println(string(c) + "  c")
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
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
