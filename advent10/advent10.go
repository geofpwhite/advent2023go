package main

import (
	"fmt"
	"os"
	"strings"
)

func part2() {
	lines, start := parse()
	lines = lines[:len(lines)-1]
	north := []int{
		start[0] - 1,
		start[1],
	}
	south := []int{
		start[0] + 1,
		start[1],
	}
	east := []int{
		start[0],
		start[1] + 1,
	}
	west := []int{
		start[0],
		start[1] - 1,
	}
	pipeString := "|-LJ7F"
	neighbors := [][]int{}
	var direction string = ""
	for i, ary := range [][]int{
		north, south, east, west,
	} {
		if strings.Contains(pipeString, string(lines[ary[0]][ary[1]])) {
			neighbors = append(neighbors, ary)
			switch i {
			case 0:
				direction = "up"
			case 1:
				direction = "down"
			case 2:
				direction = "right"
			case 3:
				direction = "left"
			}
		}
	}

	fmt.Println(neighbors[0], start, &lines, 1, [][]int{}, direction)
	_, path, weirdLines := travel(neighbors[0], start, &lines, 1, [][]int{}, direction)
	newLines := []string{}
	for i, line := range lines {
		newLines = append(newLines, "")
		for j := range line {
			if (*weirdLines)[i][j] == 'O' {
				newLines[i] += "O"
			} else {
				newLines[i] += "I"
			}

			for _, ary := range path {
				if ary[0] == i && ary[1] == j {
					newLines[i] = newLines[i][:len(newLines[i])-1]
					newLines[i] += string(lines[i][j])
				}
			}

		}
		// fmt.Println((*weirdLines)[i])
		fmt.Println(newLines[i])
	}
	for i := range newLines {
		newLines[i] = "O" + newLines[i]
	}
	newLines = append(newLines, "")
	x := len(newLines) - 1
	for range newLines[0] {
		newLines[x] += "O"
	}
	l := []string{}
	l = append(l, newLines[x])
	for i := range newLines {
		l = append(l, newLines[i])
	}
	sum := 0
	for _, line := range l {
		for _, char := range line {
			if char == 'I' {
				sum++
			}
		}
		fmt.Println(line)
	}
	println(sum)
	markOutsidePaths(&l)

	sum = 0
	for _, line := range l {
		for _, char := range line {
			if char == 'I' {
				sum++
			}
		}
		fmt.Println(line)
	}
	println(sum)

}

func markOutsidePaths(l *[]string) {
	for i, line := range *l {
		for j, char := range line {
			if char == 'O' {
				markPaths(l, []int{i, j})
			}
		}
	}
}
func markPaths(l *[]string, index []int) {
	neighbors := [][]int{
		{
			index[0], index[1] + 1,
		},
		{
			index[0], index[1] - 1,
		},
		{
			index[0] - 1, index[1],
		},
		{
			index[0] + 1, index[1],
		},
		// {
		// 	index[0] + 1, index[1] + 1,
		// },
		// {
		// 	index[0] + 1, index[1] - 1,
		// },
		// {
		// 	index[0] - 1, index[1] - 1,
		// },
		// {
		// 	index[0] - 1, index[1] + 1,
		// },
	}
	for _, ary := range neighbors {
		if ary[0] > 0 && ary[0] < len((*l)) && ary[1] > 0 && ary[1] < len((*l)[0]) {
			if (*l)[ary[0]][ary[1]] == 'I' {

				(*l)[ary[0]] = (*l)[ary[0]][:ary[1]] + "O" + (*l)[ary[0]][ary[1]+1:]
				markPaths(l, ary)
			}
		}
	}
}

// func part1() {
// 	lines, start := parse()
// 	north := []int{
// 		start[0] - 1,
// 		start[1],
// 	}
// 	south := []int{
// 		start[0] + 1,
// 		start[1],
// 	}
// 	east := []int{
// 		start[0],
// 		start[1] + 1,
// 	}
// 	west := []int{
// 		start[0],
// 		start[1] - 1,
// 	}
// 	pipeString := "|-LJ7"
// 	neighbors := [][]int{}
// 	for _, ary := range [][]int{
// 		north, south, east, west,
// 	} {
// 		if strings.Contains(pipeString, string(lines[ary[0]][ary[1]])) {
// 			neighbors = append(neighbors, ary)
// 		}
// 	}
// 	// distances, _ := travel(neighbors[0], start, lines, 1, [][]int{})
// 	// fmt.Println(distances / 2)
//
// }

func travel(index []int, prevIndex []int, lines *[]string, distance int, coordinates [][]int, prevDirection string) (int, [][]int, *[]string) {
	// time.Sleep(100 * time.Millisecond)
	var newIndex []int
	var neighbors [][]int
	newDirection := prevDirection

	println(index)
	switch string((*lines)[index[0]][index[1]]) {
	case "|":
		neighbors = [][]int{
			{
				index[0] + 1, index[1],
			}, {
				index[0] - 1, index[1],
			},
		}
		if prevIndex[0] == neighbors[0][0] && prevIndex[1] == neighbors[0][1] {
			newIndex = neighbors[1]
		} else {
			newIndex = neighbors[0]
		}
		if prevDirection == "up" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]] = (*lines)[index[0]][:index[1]+1] + "O" + (*lines)[index[0]][index[1]+2:]
				fmt.Println((*lines)[index[0]])
			}
		}
		if prevDirection == "down" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]-1])) {
				(*lines)[index[0]] = (*lines)[index[0]][:index[1]-1] + "O" + (*lines)[index[0]][index[1]:]
			}
		}

		break

	case "-":
		if prevDirection == "down" {
			fmt.Println("error")
			os.Exit(1)
		}
		neighbors = [][]int{
			{
				index[0], index[1] + 1,
			}, {
				index[0], index[1] - 1,
			},
		}
		if prevIndex[0] == neighbors[0][0] && prevIndex[1] == neighbors[0][1] {
			newIndex = neighbors[1]
		} else {
			newIndex = neighbors[0]
		}
		if prevDirection == "right" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]+1] = (*lines)[index[0]+1][:index[1]] + "O" + (*lines)[index[0]+1][index[1]+1:]
			}
		}
		if prevDirection == "left" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]-1] = (*lines)[index[0]-1][:index[1]] + "O" + (*lines)[index[0]-1][index[1]+1:]
			}
		}
		break
	case "L":
		neighbors = [][]int{
			{
				index[0] - 1, index[1],
			}, {
				index[0], index[1] + 1,
			},
		}
		if prevIndex[0] == neighbors[0][0] && prevIndex[1] == neighbors[0][1] {
			newIndex = neighbors[1]
		} else {
			newIndex = neighbors[0]
		}
		if prevDirection == "left" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]-1][index[1]+1])) {
				(*lines)[index[0]-1] = (*lines)[index[0]-1][:index[1]+1] + "O" + (*lines)[index[0]-1][index[1]+2:]
			}
			newDirection = "up"
		}
		if prevDirection == "down" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]] = (*lines)[index[0]][:index[1]-1] + "O" + (*lines)[index[0]][index[1]:]
			}
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]+1] = (*lines)[index[0]+1][:index[1]] + "O" + (*lines)[index[0]+1][index[1]+1:]
			}
			newDirection = "right"
		}
		break
	case "J":
		neighbors = [][]int{
			{
				index[0] - 1, index[1],
			}, {
				index[0], index[1] - 1,
			},
		}
		if prevIndex[0] == neighbors[0][0] && prevIndex[1] == neighbors[0][1] {
			newIndex = neighbors[1]
		} else {
			newIndex = neighbors[0]
		}

		if prevDirection == "right" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]] = (*lines)[index[0]][:index[1]+1] + "O" + (*lines)[index[0]][index[1]+2:]
			}
			if !strings.Contains("FLJ|-7", string((*lines)[index[0]+1][index[1]])) {
				(*lines)[index[0]+1] = (*lines)[index[0]+1][:index[1]] + "O" + (*lines)[index[0]+1][index[1]+1:]
			}
			newDirection = "up"
		}
		if prevDirection == "down" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]-1][index[1]-1])) {
				(*lines)[index[0]-1] = (*lines)[index[0]-1][:index[1]-1] + "O" + (*lines)[index[0]-1][index[1]:]
			}
			newDirection = "left"
		}

		break
	case "7":
		neighbors = [][]int{
			{
				index[0], index[1] - 1,
			}, {
				index[0] + 1, index[1],
			},
		}
		if prevIndex[0] == neighbors[0][0] && prevIndex[1] == neighbors[0][1] {
			newIndex = neighbors[1]
		} else {
			newIndex = neighbors[0]
		}
		if prevDirection == "right" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]+1][index[1]-1])) {
				(*lines)[index[0]+1] = (*lines)[index[0]+1][:index[1]-1] + "O" + (*lines)[index[0]+1][index[1]:]
			}
			newDirection = "down"
		}
		if prevDirection == "up" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]-1][index[1]])) {
				(*lines)[index[0]-1] = (*lines)[index[0]-1][:index[1]] + "O" + (*lines)[index[0]-1][index[1]+1:]
			}
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]+1])) {
				(*lines)[index[0]] = (*lines)[index[0]][:index[1]+1] + "O" + (*lines)[index[0]][index[1]+2:]
			}
			newDirection = "left"
		}

		break
	case "F":
		neighbors = [][]int{
			{
				index[0], index[1] + 1,
			}, {
				index[0] + 1, index[1],
			},
		}
		if prevIndex[0] == neighbors[0][0] && prevIndex[1] == neighbors[0][1] {
			newIndex = neighbors[1]
		} else {
			newIndex = neighbors[0]
		}
		if prevDirection == "left" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]][index[1]-1])) {
				(*lines)[index[0]] = (*lines)[index[0]][:index[1]-1] + "O" + (*lines)[index[0]][index[1]:]
			}
			if !strings.Contains("FJL7|-", string((*lines)[index[0]-1][index[1]])) {
				(*lines)[index[0]-1] = (*lines)[index[0]-1][:index[1]] + "O" + (*lines)[index[0]-1][index[1]+1:]
			}
			newDirection = "down"
		}
		if prevDirection == "up" {
			if !strings.Contains("FJL7|-", string((*lines)[index[0]+1][index[1]+1])) {
				(*lines)[index[0]+1] = (*lines)[index[0]+1][:index[1]+1] + "O" + (*lines)[index[0]+1][index[1]+2:]
			}
			newDirection = "right"
		}
		break
	case ".", "O":

		for i := range *lines {
			println((*lines)[i])
		}
		os.Exit(1)

	case "S":
		return distance, coordinates, lines
	}
	newCoordinates := append(coordinates, newIndex)
	fmt.Println(string((*lines)[index[0]][index[1]]))
	fmt.Println(newDirection)
	// println(string((*lines)[newIndex[0]][newIndex[1]]))
	return travel(newIndex, index, lines, distance+1, newCoordinates, newDirection)

}

func parse() ([]string, []int) {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	start := []int{}
	for i := range lines {
		for j, ch := range lines[i] {
			if ch == 'S' {
				start = append(start, i)
				start = append(start, j)
				break
			}
		}
	}
	return lines, start
}
func main() {
	// part1()
	part2()
}
