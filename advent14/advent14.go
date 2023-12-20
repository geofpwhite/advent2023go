package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	value       string
	coordinates [2]int
	above       *node
	below       *node
}

type graph struct {
	nodes []node
	lines []string
}

func (g graph) rollDown() {
	columns := make([]string, len(g.lines[0]))

	for i := range g.lines[0] {
		columns[i] = ""
	}
	for _, line := range g.lines {
		for i := range line {
			columns[i] += string(line[i])
		}
	}

	for i := range columns {
		for strings.Contains(columns[i], "O.") {
			columns[i] = strings.Replace(columns[i], "O.", ".O", 1)
		}
	}

	for i := range columns[0] {
		g.lines[i] = ""
	}
	for _, line := range columns {
		for i := range line {
			g.lines[i] += string(line[i])
		}
	}
}
func (g graph) rollUp() {
	columns := make([]string, len(g.lines[0]))

	for i := range g.lines[0] {
		columns[i] = ""
	}
	for _, line := range g.lines {
		for i := range line {
			columns[i] += string(line[i])
		}
	}
	for i := range columns {
		for strings.Contains(columns[i], ".O") {
			columns[i] = strings.Replace(columns[i], ".O", "O.", 1)
		}
	}
	for i := range columns[0] {
		g.lines[i] = ""
	}
	for _, line := range columns {
		for i := range line {
			g.lines[i] += string(line[i])
		}
	}

}

func (g graph) rollLeft() {
	for i := range g.lines {
		for strings.Contains(g.lines[i], ".O") {
			g.lines[i] = strings.Replace(g.lines[i], ".O", "O.", 1)
		}
	}
}
func (g graph) rollRight() {
	for i := range g.lines {
		for strings.Contains(g.lines[i], "O.") {
			g.lines[i] = strings.Replace(g.lines[i], "O.", ".O", 1)
		}
	}
}

func (g graph) checkSum() int {
	columns := make([]string, len(g.lines[0]))

	for i := range g.lines[0] {
		columns[i] = ""
	}
	for _, line := range g.lines {
		for i := range line {
			columns[i] += string(line[i])
		}
	}
	sum := 0
	for i := range columns {
		for j, c := range columns[i] {
			if c == 'O' {
				sum += len(columns[i]) - j
			}
		}
	}
	println(sum)
	return sum

}

func main() {
	g := parse()
	fmt.Println(g.lines)
	// g.rollUp()
	// g.checkSum()

	sums := []int{}
	for i := 0; i < 2000; i++ {
		g.rollUp()
		g.rollLeft()
		g.rollDown()
		g.rollRight()
		sums = append(sums, g.checkSum())

	}
	fmt.Println((sums))
	println(sums[(1000000000+2000)%7+3])
	x := sums[len(sums)-10 : len(sums)-1]
	fmt.Println(x[(1000000000-2000)%9])
	g.checkSum()
	//   91270 91278 91295 91317 91333 91332 91320 91306 91286
}

func parse() graph {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(content), "\n")
	g := graph{
		nodes: []node{},
	}
	coordHold := [2]int{0, 0}
	for _, line := range lines {
		for _, char := range line {
			var n node = node{
				value:       string(char),
				coordinates: coordHold,
			}
			g.nodes = append(g.nodes, n)
			coordHold[1]++
		}
		coordHold[0]++
		coordHold[1] = 0
	}
	g.lines = lines
	return g

}
