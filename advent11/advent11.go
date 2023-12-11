package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {
	strings := parse()
	galaxies := [][]int{}
	distances := []int{}
	for i, line := range strings {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}
	for i, coords := range galaxies {
		for _, edgeCoords := range galaxies[i+1:] {
			xdistance := edgeCoords[0] - coords[0]
			ydistance := edgeCoords[1] - coords[1]
			if xdistance < 0 {
				xdistance *= -1
			}
			if ydistance < 0 {
				ydistance *= -1
			}
			distances = append(distances, xdistance+ydistance)
		}
	}

	sum := 0
	for _, num := range distances {
		sum += num
		fmt.Println(num)
	}
	fmt.Println(sum)
}

func parse() []string {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	strings := []string{}
	for i := range lines {
		empty := true
		str := ""
		if lines[i] == "" || lines[i] == " " {
			continue
		}
		for j := range lines[i] {
			char := lines[i][j]
			str += string(char)
			if char == '#' {
				empty = false
			}
		}
		if empty {
			strings = append(strings, str)
		}
		strings = append(strings, str)
	}
	emptyColumns := []int{}
	for _, line := range strings {
		fmt.Println(line)
	}

	for i := 0; i < len(lines[0]); i++ {
		empty := true
		for j := range lines {
			if lines[j] == "" {
				continue
			}
			fmt.Println(lines[j])
			if lines[j][i] == '#' {
				empty = false
			}
		}
		if empty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	add := 0
	for _, column := range emptyColumns {
		for i := range strings {
			fmt.Println(strings[i])
			strings[i] = strings[i][:column+add] + "." + strings[i][column+add:]
			fmt.Println(strings[i])
		}
		add++
	}

	for _, line := range strings {
		fmt.Println(line)
	}
	return strings

}

func main() {
	part1()
}
