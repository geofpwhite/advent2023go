package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	inputs := parse()
	// fmt.Println(inputs)
	sum := 0
	for _, input := range inputs {
		sum += history(input)
	}
	fmt.Println(sum)
}
func part2() {
	inputs := parse()
	// fmt.Println(inputs)
	sum := 0
	for _, input := range inputs {
		sum += history2(input)
	}
	fmt.Println(sum)
}

func allZeroes(line []int) bool {
	for i := range line {
		if line[i] != 0 {
			return false
		}
	}
	return true
}
func history2(line []int) int {
	curLine := line
	history := [][]int{
		curLine,
	}
	for !allZeroes(curLine) {
		newLine := []int{}

		for i := range curLine[:len(curLine)-1] {
			val := curLine[i+1] - curLine[i]
			newLine = append(newLine, val)
		}
		history = append(history, newLine)
		curLine = newLine

	}
	curNum := 0
	for i := len(history) - 1; i > 0; i-- {

		if len(history[i]) == 0 {
			curNum += 0
		}
		curNum = history[i-1][0] - curNum
	}
	return curNum
}

func history(line []int) int {
	curLine := line
	history := [][]int{
		curLine,
	}
	for !allZeroes(curLine) {
		newLine := []int{}

		for i := range curLine[:len(curLine)-1] {
			abs := curLine[i+1] - curLine[i]
			// if abs < 0 {
			// 	abs *= -1
			// }
			// fmt.Println(curLine[i+1]-curLine[i], abs)
			newLine = append(newLine, abs)
			// fmt.Println(abs, int(abs))
		}
		history = append(history, newLine)
		curLine = newLine
		// fmt.Println(curLine)

	}
	curNum := 0

	for i := 0; i < len(history)-1; i++ {
		if len(history[i]) == 0 {
			curNum += 0
		}
		curNum += history[i][len(history[i])-1]
	}
	for i := 0; i < len(history); i++ {
		fmt.Println(history[i])
	}
	return curNum
}
func parse() [][]int {
	content, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(content), "\n")
	inputs := [][]int{}
	for _, line := range lines {
		nummies := []int{}
		nums := strings.Split(line, " ")
		if len(nums) == 0 {
			continue
		}
		for i := range nums {
			if nums[i] == "" {
				continue
			}
			number := nums[i]
			if !strings.Contains("1234567890", string(number[len(number)-1])) {
				number = number[:len(number)-1]
			}
			num, _ := strconv.Atoi(strings.ReplaceAll(number, " ", ""))
			nummies = append(nummies, num)
		}
		inputs = append(inputs, nummies)
	}
	return inputs[:len(inputs)-1]
}

func main() {
	part1()
	part2()
}
