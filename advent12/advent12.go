package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {

	lines, numbers := parse()

	sum := 0
	for i := range lines {

		nn := []int{}
		for _, num := range numbers[i] {
			nn = append(nn, num)
		}
		for _, num := range numbers[i] {
			nn = append(nn, num)
		}

		x := valid(lines[i], numbers[i])
		// y := valid(lines[i]+".", numbers[i])
		// w := valid("#"+lines[i], numbers[i])
		v := valid(lines[i]+"#", numbers[i])
		// u := valid("."+lines[i], numbers[i])
		println()
		// fmt.Println(x, y, w, v, u)
		fmt.Println("line ", i)
		if v == 0 {
			z := valid(lines[i]+"?"+lines[i], nn)
			// fmt.Println(startLine*endLine, z)
			fmt.Println(z, " z")
			if z == x*x {

				fmt.Println(x * x * x * x * x)
				sum += (x * x * x * x * x)
				println()
				continue
			}
		}
		startLine := valid(lines[i]+"?", numbers[i])
		endLine := valid("?"+lines[i], numbers[i])
		middleLine := valid("?"+lines[i]+"?", numbers[i])

		fmt.Println(startLine, middleLine, middleLine, middleLine, endLine)
		fmt.Println(startLine * middleLine * middleLine * middleLine * endLine)
		sum += (startLine * middleLine * middleLine * middleLine * endLine)
		// middleLine2 := valid("#" + lines[i] + "#",numbers[i])
		// middleLine3 := valid("." + lines[i] + "#",numbers[i])
		// middleLine4 := valid("#" + lines[i] + ".",numbers[i])

		println()
	}
	fmt.Println(sum)
}
func part1() {
	lines, numbers := parse()

	scores := []int{}
	for i, line := range lines {
		scores = append(scores, valid(line, numbers[i]))
	}

	sum := 0

	for _, score := range scores {
		sum += score
	}

	fmt.Println(sum)

}

func valid(str string, nums []int) int {

	if strings.Contains(str, "?") {

		index := strings.Index(str, "?")
		str1 := str[:index] + "." + str[index+1:]
		str2 := str[:index] + "#" + str[index+1:]

		return valid(str1, nums) + valid(str2, nums)
	} else {
		if !strings.Contains(str, "#") {
			return 0
		}
		line := strings.Trim(str, ".")
		for line[0] == '.' || line[len(line)-1] == '.' {

			line = strings.Trim(str, ".")
		}
		numIndex := 0

		sum := 0
		sumNum := 0
		for _, char := range line {
			if char == '#' {
				sum++
			}
		}

		for _, num := range nums {
			sumNum += num
		}

		if sumNum != sum {
			return 0
		}
		for len(line) > 0 {

			if numIndex == len(nums) && len(line) > 0 {
				return 0
			}
			if numIndex == len(nums) && line == "" {
				return 1
			}

			num := nums[numIndex]

			if numIndex > 0 {
				if line[0] == '#' {
					return 0
				}
			}
			for num > 0 {
				// fmt.Println(num)
				line = strings.Trim(line, ".")
				if len(line) == 0 {
					return 0
				}
				if line[0] == '#' {
					line = line[1:]
					num--
					if num > 0 && line[0] == '.' {
						return 0
					}

					if len(line) > 0 && num == 0 && numIndex == len(nums)-1 {
						return 0
					}
				} else {
					return 0
				}
			}

			numIndex++
		}
		return 1
	}

}

func parse() ([]string, [][]int) {
	// content, _ := os.ReadFile("input.txt")
	content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("oneline.txt")

	lines := strings.Split(string(content), "\n")
	numbers := [][]int{}

	lines = lines[:len(lines)-1]
	for i, line := range lines {
		nums := strings.Split(line[strings.Index(line, " ")+1:], ",")
		numInts := []int{}
		for _, num := range nums {
			number, _ := strconv.Atoi(num)
			numInts = append(numInts, number)
		}
		numbers = append(numbers, numInts)
		lines[i] = line[:strings.Index(line, " ")]
	}
	return lines, numbers
}

func main() {
	part1()
	// part2()
}
