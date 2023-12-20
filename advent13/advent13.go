package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
	// f()
}
func parse() [][]string {
	content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
	// content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	valley := make([][]string, 0)
	curary := make([]string, 0)

	for _, line := range lines {
		if line == "" || line == "\n" {
			if len(curary) > 0 {
				valley = append(valley, curary)
				curary = make([]string, 0)
			}
		} else {
			curary = append(curary, line)
		}
	}
	return valley
}

func part2() {
	rowsAbove, columnsLeft := 0, 0

	valley := parse()
	for _, input := range valley {
		x, y := 0, 0
		sizev1 := 0
		sizeh1 := 0

		for i := range input {
			mirror, size := mirrorOnIndex(input, false, i)
			if mirror && size > sizeh1 {
				x = i + 1
				sizeh1 = size
			}
		}

		for i := range input[0] {

			mirror, size := mirrorOnIndex(input, true, i)

			if mirror && size > sizev1 {
				y = i + 1
				sizev1 = size
			}
		}

		var og string
		if y > x {
			og = strconv.Itoa(y) + " v"
		} else {
			og = strconv.Itoa(x) + " h"
		}

		// sizev := 0
		// sizeh := 0

		s, p := 0, 0
		for i := range input {
			for j := range input[i] {
				// sizev = 0
				// sizeh = 0
				var input2 []string = make([]string, len(input))
				for q, line := range input {
					input2[q] = line
				}
				if input2[i][j] == '.' {
					input2[i] = input[i][:j] + "#" + input[i][j+1:]
				} else {
					input2[i] = input[i][:j] + "." + input[i][j+1:]
				}

				for k := range input2 {
					mirror, size := mirrorOnIndex(input2, false, k)
					if mirror {

					}
					if mirror && strconv.Itoa(k+1)+" h" != og {

						s = k + 1
						fmt.Println("v=false mirrored on", k, "with size", size)
						// sizeh = size
					}
				}
				for k := range input2[0] {

					mirror, _ := mirrorOnIndex(input2, true, k)
					if mirror {

					}

					if mirror && strconv.Itoa(k+1)+" v" != og {
						p = k + 1
						// sizev = size
					}
				}

			}
		}

		if p > s {

			columnsLeft += p
			if p > 0 {
				fmt.Println(p)
				fmt.Println(columnsLeft-p, "+", p, "=", columnsLeft)
				continue
			}
		} else {
			rowsAbove += s
			if s > 0 {
				fmt.Println(s)
				fmt.Println(rowsAbove-s, "+", s, "=", rowsAbove)
				continue
			}
		}
		// rowsAbove+=s

		// columnsLeft += i

	}
	fmt.Println(rowsAbove, columnsLeft, (rowsAbove*100)+columnsLeft)

}

func part1() {
	rowsAbove, columnsLeft := 0, 0

	valley := parse()
	for _, input := range valley {
		s, p := 0, 0
		sizev := 0
		sizeh := 0
		for i := range input {
			mirror, size := mirrorOnIndex(input, false, i)
			if mirror && size > sizeh {
				s = i + 1
				sizeh = size
			}
		}

		for i := range input[0] {

			mirror, size := mirrorOnIndex(input, true, i)

			if mirror && size > sizev {
				p = i + 1
				sizev = size
			}
		}

		if p > s {
			columnsLeft += p
			fmt.Println("v=true mirrored on", p)
		} else {
			fmt.Println("v=false mirrored on", s)
			rowsAbove += s
		}
		fmt.Println(s, p)
		fmt.Println(sizev, sizeh)

	}
	fmt.Println(rowsAbove, columnsLeft, (rowsAbove*100)+columnsLeft)

}

func mirrorOnIndex(input []string, vertical bool, index int) (bool, int) {
	if vertical {

		if index == 0 {
			return columnsMatch(input, 0, 1), 1
		}
		if index == len(input[0])-1 {
			return columnsMatch(input, len(input[0])-2, len(input[0])-1), 1
		}
		i, j := 0, 0
		for i, j = index, index+1; i >= 0 && j < len(input[0]); i-- {
			if !columnsMatch(input, i, j) {
				return false, 0
			}

			j++
		}
		return true, j - i
	} else {
		if index == 0 {
			return rowsMatch(input, 0, 1), 1
		}
		if index == len(input)-1 {
			return rowsMatch(input, len(input)-2, len(input)-1), 1
		}
		i, j := 0, 0
		for i, j = index, index+1; i >= 0 && j < len(input); i-- {
			if !rowsMatch(input, i, j) {
				return false, 0
			}

			j++
		}

		return true, j - i
	}
}

func rowsMatch(input []string, i1, i2 int) bool {
	return input[i1] == input[i2]
}

func columnsMatch(input []string, i1, i2 int) bool {
	c1, c2 := "", ""

	for _, line := range input {
		c1 += string(line[i1])
		c2 += string(line[i2])
	}
	return c1 == c2
}
