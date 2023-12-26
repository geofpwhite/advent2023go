package main

import (
	"encoding/hex"
	"os"
	"strconv"
	"strings"
)

const (
	R int = 0
	D int = 1
	L int = 2
	U int = 4
)

type instruction struct {
	direction int
	color     string
	length    int
}

func parse() ([1000][1000]rune, []instruction) {
	// content, _ := os.ReadFile("input.txt")
	content, _ := os.ReadFile("test.txt")
	instructionstrings := strings.Split(string(content), "\n")
	if instructionstrings[len(instructionstrings)-1] == "" {
		instructionstrings = instructionstrings[:len(instructionstrings)-1]

	}
	ins := []instruction{}
	for i := range instructionstrings {
		instructionstrings[i] = strings.Replace(instructionstrings[i], "\r", "", -1)
		inst := instruction{}
		switch instructionstrings[i][0] {
		case 'L':
			inst.direction = L
		case 'R':
			inst.direction = R
		case 'U':
			inst.direction = U
		case 'D':
			inst.direction = D
		}
		length, _ := strconv.Atoi(instructionstrings[i][2 : strings.Index(instructionstrings[i][2:], " ")+2])
		inst.length = length
		inst.color = instructionstrings[i][strings.Index(instructionstrings[i][2:], " ")+2:]
		inst.color = inst.color[3 : len(inst.color)-1]
		ins = append(ins, inst)

	}
	area := [1000][1000]rune{}
	{
		for i := 0; i < 1000; i++ {
			for j := 0; j < 1000; j++ {
				area[i][j] = '.'
			}
		}
	}
	return area, ins

}

func opposite(char int) int {
	switch char {
	case U:
		return D
	case D:
		return U
	case L:
		return R
	case R:
		return L
	}
	return -128
}

func pop(ary *[]instruction) instruction {
	l := len(*ary)
	r := (*ary)[l-1]
	*ary = (*ary)[:l-1]
	return r
}

// # python3 program to evaluate
// # area of a polygon using
// # shoelace formula

// # (X[i], Y[i]) are coordinates of i'th point.
func areaFromVertices(vertices [][2]int) int64 { // def polygonArea(X, Y, n):

	var area int64 = 0
	j := len(vertices) - 1
	for i := 0; i < len(vertices); i++ {
		area += int64(vertices[i][0] * vertices[j][1])
		area -= int64((vertices[j][0]) * (vertices[i][1]))
		j = i
	}
	return area / 2

}

//     # Initialize area
//     area = 0.0

//     # Calculate value of shoelace formula
//     j = n - 1
//     for i in range(0,n):
//         area += (X[j] + X[i]) * (Y[j] - Y[i])
//         j = i   # j is previous vertex to i

//     # Return absolute value
//     return int(abs(area / 2.0))

// # Driver program to test above function
// X = [0, 2, 4]
// Y = [1, 3, 7]
// n = len(X)
// print(polygonArea(X, Y, n))

// # This code is contributed by
// # Smitha Dinesh Semwal1

func part2() {
	_, instructions := parse()
	vertices := make([][2]int, len(instructions))
	cur := [2]int{500, 500}
	distTraveled := 0
	for i, instr := range instructions {
		instr.direction, _ = strconv.Atoi(string(instr.color[len(instr.color)-1]))
		x, err := hex.DecodeString("0" + instr.color[:len(instr.color)-1])
		y := 5
		sum := 0
		for i, char := range instr.color[:y] {
			pow := 1
			for m := 0; m < y-i-1; m++ {
				pow *= 16
			}
			switch char {
			case 'a':
				sum += 10 * pow
			case 'b':
				sum += 11 * pow
			case 'c':
				sum += 12 * pow
			case 'd':
				sum += 13 * pow
			case 'e':
				sum += 14 * pow
			case 'f':
				sum += 15 * pow
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				num, _ := strconv.Atoi(string(char))
				sum += num * pow
			}
		}
		instr.length = sum
		println(x)

		if err != nil {
			panic("error")
		}
		switch instr.direction {
		case U:
			cur[0] -= instr.length
		case D:
			cur[0] += instr.length
		case L:
			cur[1] -= instr.length
		case R:
			cur[1] += instr.length
		}
		distTraveled += instr.length - 1
		vertices[i] = cur
	}
	// vertices = append(vertices, [2]int{500, 500})
	x := areaFromVertices(vertices)
	println(x)
	println(distTraveled)
	println(len(vertices))
	println(int(x) + distTraveled - len(vertices))
	println(int(x) + len(vertices)*3)
}

func part1() {
	area, instructions := parse()
	cur := [2]int{500, 500}
	for _, instrctn := range instructions {
		area[cur[0]][cur[1]] = '#'
		for i := 0; i < instrctn.length; i++ {
			area[cur[0]][cur[1]] = '#'
			if instrctn.direction == L {
				cur[1]--
				area[cur[0]][cur[1]] = '#'

			}
			if instrctn.direction == R {
				cur[1]++
				area[cur[0]][cur[1]] = '#'

			}
			if instrctn.direction == U {
				cur[0]--
				area[cur[0]][cur[1]] = '#'

			}
			if instrctn.direction == D {
				cur[0]++
				area[cur[0]][cur[1]] = '#'

			}
		}
	}
	data := []byte{}
	for i := range area {
		for j := range area[i] {
			data = append(data, byte(area[i][j]))

		}
		data = append(data, '\n')
	}
	os.WriteFile("output.txt", data, 0644)

	sum := 0
	markArea(&area, [2]int{603, 661}, &sum)
	sum = 0
	for _, r := range area {
		for _, m := range r {
			if m == '#' {
				sum++
			}
		}
	}
	println(sum)

}

func markArea(areaPointer *[1000][1000]rune, cur [2]int, sum *int) {
	areaPointer[cur[0]][cur[1]] = '#'
	x := cur[0]
	y := cur[1]

	if x < 999 {
		if areaPointer[x+1][y] != '#' {
			markArea(areaPointer, [2]int{x + 1, y}, sum)
		}
	}
	if x > 0 {

		if areaPointer[x-1][y] != '#' {
			markArea(areaPointer, [2]int{x - 1, y}, sum)
		}
	}
	if y < 999 {

		if areaPointer[x][y+1] != '#' {
			markArea(areaPointer, [2]int{x, y + 1}, sum)
		}
	}
	if y > 0 {

		if areaPointer[x][y-1] != '#' {
			markArea(areaPointer, [2]int{x, y - 1}, sum)
		}
	}

}

func main() {
	part1()
	part2()
}
