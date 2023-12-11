package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {
	strings, emptyLines, emptyColumns := parse()
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
			lx := edgeCoords[0]
			sx := coords[0]
			ly := edgeCoords[1]
			sy := coords[1]
			if xdistance < 0 {
				xdistance *= -1
				lx, sx = sx, lx
			}
			if ydistance < 0 {
				ydistance *= -1
				ly, sy = sy, ly
			}
			distances = append(distances, xdistance+ydistance)

			for _, num := range emptyLines {
				if num > sx && num < lx {
					distances[len(distances)-1] += 999999
				}
			}

			for _, num := range emptyColumns {
				if num > sy && num < ly {
					distances[len(distances)-1] += 999999
				}
			}
		}
	}

	sum := 0
	for _, num := range distances {
		sum += num
		fmt.Println(num)
	}
	fmt.Println(sum)
}

func parse() ([]string, []int, []int) {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	strings := []string{}
	emptyLines := []int{}
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
			emptyLines = append(emptyLines, i)
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
			if lines[j][i] == '#' {
				empty = false
			}
		}
		if empty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	for _, line := range strings {
		fmt.Println(line)
	}
	return strings, emptyLines, emptyColumns

}

func main() {
	part1()
}
