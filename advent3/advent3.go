package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSurroundingNumber(str string, index int) int {
	digits := "1234567890"
	if !strings.Contains(digits, string(str[index])) {
		return -1
	} else {
		numberString := string(str[index])
		i := 0
		for i = index + 1; strings.Contains(digits, string(str[i])) && i < len(str)-1; i++ {
			numberString += string(str[i])
		}
		for i = index - 1; strings.Contains(digits, string(str[i])) && i > 0; i-- {
			numberString = string(str[i]) + numberString
		}
		if i == 0 {
			numberString = string(str[0]) + numberString
		}
		num, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("err")
		}
		return num
	}

}

func part2() {

	mainSpokes := [][]int{}
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

	symbols := "/!@?[]{}#$%&*+-="
	digits := "1234567890"
	ar := strings.Split(string(content), "\n")
	symbolCoordinates := []string{} // "lineNumber,charIndex"
	for il, line := range ar {
		for ic, char := range line {
			if strings.ContainsRune(symbols, char) {
				coordString := strconv.Itoa(il) + "," + strconv.Itoa(ic)
				symbolCoordinates = append(symbolCoordinates, coordString)
			}
		}
	}

	fmt.Println(symbolCoordinates)

	for _, str := range symbolCoordinates {
		lineNumber, err := strconv.Atoi(str[0:strings.Index(str, ",")])
		if err != nil {

		}
		charIndex, err := strconv.Atoi(str[strings.Index(str, ",")+1:])
		if err != nil {

		}
		spokes := []int{}
		start := false
		end := false
		if charIndex == 0 {
			start = true
		}

		if charIndex == len(ar[0])-1 {
			end = true
		}

		//get spokes from above
		if !start && !end {
			if lineNumber > 0 {
				lineAbove := ar[lineNumber-1]

				//above line
				if strings.Contains(digits, string(lineAbove[charIndex-1])) {

					if strings.Contains(digits, string(lineAbove[charIndex])) {

						if strings.Contains(digits, string(lineAbove[charIndex+1])) {
							// above line is 111
							number := getSurroundingNumber(lineAbove, charIndex)
							spokes = append(spokes, number)

						} else {
							//above line is 11x
							number := getSurroundingNumber(lineAbove, charIndex)
							spokes = append(spokes, number)

						}

					} else {
						if strings.Contains(digits, string(lineAbove[charIndex+1])) {
							// above line has 1x1
							//aka has two spokes
							number := getSurroundingNumber(lineAbove, charIndex-1)
							spokes = append(spokes, number)
							number = getSurroundingNumber(lineAbove, charIndex+1)
							spokes = append(spokes, number)

						} else {
							//above line has 1xx
							number := getSurroundingNumber(lineAbove, charIndex-1)
							spokes = append(spokes, number)
						}

					}
				} else {
					if strings.Contains(digits, string(lineAbove[charIndex])) {

						if strings.Contains(digits, string(lineAbove[charIndex+1])) {
							//above line has x11
							number := getSurroundingNumber(lineAbove, charIndex)
							spokes = append(spokes, number)

						} else {
							//above line has x1x
							number := getSurroundingNumber(lineAbove, charIndex)
							spokes = append(spokes, number)
						}
					} else {
						if strings.Contains(digits, string(lineAbove[charIndex+1])) {
							//above line has xx1
							number := getSurroundingNumber(lineAbove, charIndex+1)
							spokes = append(spokes, number)
						} else {
							//above line has xxx
						}
					}
				}
			}

			if lineNumber < len(ar)-2 {
				lineBelow := ar[lineNumber+1]

				if strings.Contains(digits, string(lineBelow[charIndex-1])) {

					if strings.Contains(digits, string(lineBelow[charIndex])) {

						if strings.Contains(digits, string(lineBelow[charIndex+1])) {
							// above line is 111
							number := getSurroundingNumber(lineBelow, charIndex)
							spokes = append(spokes, number)

						} else {
							//above line is 11x
							number := getSurroundingNumber(lineBelow, charIndex)
							spokes = append(spokes, number)

						}

					} else {
						if strings.Contains(digits, string(lineBelow[charIndex+1])) {
							// above line has 1x1
							//aka has two spokes
							number := getSurroundingNumber(lineBelow, charIndex-1)
							spokes = append(spokes, number)
							number = getSurroundingNumber(lineBelow, charIndex+1)
							spokes = append(spokes, number)

						} else {
							//above line has 1xx
							number := getSurroundingNumber(lineBelow, charIndex-1)
							spokes = append(spokes, number)
						}

					}
				} else {
					if strings.Contains(digits, string(lineBelow[charIndex])) {

						if strings.Contains(digits, string(lineBelow[charIndex+1])) {
							//above line has x11
							number := getSurroundingNumber(lineBelow, charIndex)
							spokes = append(spokes, number)

						} else {
							//above line has x1x
							number := getSurroundingNumber(lineBelow, charIndex)
							spokes = append(spokes, number)
						}
					} else {
						if strings.Contains(digits, string(lineBelow[charIndex+1])) {
							//above line has xx1
							number := getSurroundingNumber(lineBelow, charIndex+1)
							spokes = append(spokes, number)
						} else {
							//above line has xxx
						}
					}
				}
			}

			curLine := ar[lineNumber]
			if strings.Contains(digits, string(curLine[charIndex-1])) {
				number := getSurroundingNumber(curLine, charIndex-1)
				spokes = append(spokes, number)
			}
			if strings.Contains(digits, string(curLine[charIndex+1])) {
				number := getSurroundingNumber(curLine, charIndex+1)
				spokes = append(spokes, number)
			}
		} else if start {
			//symbol is at the start of the line

			//above line
			//11

			//1x
			//x1

			//number on the right

			//below line

			//1x
			//x1

		} else if end {
			//symbol is at the end of the line

			//above line
			//1x
			//x1

			//number on the right

			//below line
			//1x
			//x1
		}
		if len(spokes) > 1 {
			mainSpokes = append(mainSpokes, spokes)
		}

	}

	sum := 0
	for _, ar := range mainSpokes {
		num := 1
		for _, number := range ar {
			fmt.Println(number)

			num *= number
		}
		sum += num
	}
	fmt.Println(sum)
}

func part1() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0

	ar := strings.Split(string(content), "\n")
	validNums := []int{}
	for i := range ar {

		// s := string(c)
		// substrings := strings.Split(s, ",")

		numbers := []int{}
		if i == 0 {
			validNums = valid(numbers, "", ar[0], ar[1])
		} else if i == len(ar)-1 {
			validNums = valid(numbers, ar[i-1], ar[i], "")
		} else {
			validNums = valid(numbers, ar[i-1], ar[i], ar[i+1])
		}

		for _, num := range validNums {
			sum += num
			fmt.Println("sum + ", num)
		}
		fmt.Println(validNums)

	}
	fmt.Println(sum)
}

func valid(numbers []int, lineAbove string, curLine string, lineBelow string) []int {

	validNumbers := []int{}

	curNumber := ""

	indices := []int{}
	for i, c := range curLine {
		if strings.Contains("1234567890", string(c)) {
			if curNumber == "" {
				indices = append(indices, i)
			}
			curNumber += string(c)
		} else {
			if curNumber != "" {
				x, err := strconv.Atoi(curNumber)
				if err != nil {
					fmt.Println("err")
				} else {
					numbers = append(numbers, x)
				}
				curNumber = ""

			}
		}
	}
	if curNumber != "" {
		x, err := strconv.Atoi(curNumber)
		if err != nil {
			fmt.Println("err")
		}
		numbers = append(numbers, x)
	}

	symbols := "/!@?[]{}#$%&*+-="
	fmt.Println(numbers)
	for i, num := range numbers {
		valid := false
		numberString := strconv.Itoa(num)
		firstIndex := indices[i]
		lastIndex := firstIndex + len(numberString)
		if firstIndex == 0 {
			firstIndex = 1
		}
		if lastIndex == len(curLine) {
			lastIndex -= 1
		}

		// if strings.ContainsRune(symbols, rune(curLine[firstIndex-2])) {
		// 	valid = true
		// }
		for _, cha := range curLine[firstIndex-1 : lastIndex+1] {
			if strings.ContainsRune(symbols, (cha)) {
				fmt.Println(string(cha))
				valid = true
			}
		}
		if !valid && lineAbove != "" {
			for _, cha := range lineAbove[firstIndex-1 : lastIndex+1] {
				if strings.ContainsRune(symbols, (cha)) {
					valid = true
				}
			}
		}
		if lineBelow != "" && !valid {
			for _, cha := range lineBelow[firstIndex-1 : lastIndex+1] {
				if strings.ContainsRune(symbols, (cha)) {
					valid = true
				}
			}

		}

		if valid {
			validNumbers = append(validNumbers, num)
		}

	}

	return validNumbers
}

func main() {
	part1()
	part2()
}
