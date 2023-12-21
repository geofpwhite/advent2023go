package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func parse() []string {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.Replace(lines[i], "\r", "", -1)
		fmt.Println(len(lines[i]))
	}
	fmt.Println(lines)
	return lines
}

const (
	LEFT  int = 0
	RIGHT int = 1
	UP    int = 2
	DOWN  int = 3
)

type laser struct {
	direction   int
	coordinates [2]int
}

type lasers struct {
	lasers    []laser
	field     []string
	usedPipes map[[2]int]bool
}

func (l *lasers) move() {
	toRemove := []int{}
	x := map[[3]int]bool{}
	if len(l.lasers) > 20000 {
		l.lasers = l.lasers[len(l.lasers)-20000:]
		for i := range l.lasers {
			y := [3]int{
				l.lasers[i].coordinates[0],
				l.lasers[i].coordinates[1],
				l.lasers[i].direction,
			}
			x[y] = true
		}
		l.lasers = []laser{}
		for key, value := range x {
			if value {

				lz := laser{
					direction: key[2],
					coordinates: [2]int{
						key[0],
						key[1],
					},
				}
				l.lasers = append(l.lasers, lz)
			}
		}
	}
	// println(len(l.lasers))

	for i := range l.lasers {

		if len(l.field) <= l.lasers[i].coordinates[0] || len(l.field[0]) <= l.lasers[i].coordinates[1] || l.lasers[i].coordinates[0] < 0 || l.lasers[i].coordinates[1] < 0 {
			toRemove = append(toRemove, i)
			continue
		}

		if l.field[l.lasers[i].coordinates[0]][l.lasers[i].coordinates[1]] == '.' {
			x := l.field[l.lasers[i].coordinates[0]]
			x = x[:l.lasers[i].coordinates[1]] + "#" + x[l.lasers[i].coordinates[1]+1:]
			l.field[l.lasers[i].coordinates[0]] = x
		}
		l.usedPipes[l.lasers[i].coordinates] = true

		switch l.lasers[i].direction {

		case LEFT:
			{
				x := l.field[l.lasers[i].coordinates[0]][l.lasers[i].coordinates[1]]
				switch x {
				case '#', '-':
					if l.lasers[i].coordinates[1] > 0 {
						l.lasers[i].coordinates[1]--
					}
				case '\\':
					l.lasers[i].direction = UP
					if l.lasers[i].coordinates[0] > 0 {
						l.lasers[i].coordinates[0]--
					}
				case '/':
					l.lasers[i].direction = DOWN
					if l.lasers[i].coordinates[0] < len(l.field)-1 {
						l.lasers[i].coordinates[0]++
					}
				case '|':
					lasr := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0] + 1,
							l.lasers[i].coordinates[1],
						},
						direction: DOWN}
					lasr2 := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0] - 1,
							l.lasers[i].coordinates[1],
						},
						direction: UP,
					}
					if lasr.coordinates[0] < len(l.field) {
						l.lasers = append(l.lasers, lasr)
					}
					if lasr2.coordinates[0] >= 0 {
						l.lasers = append(l.lasers, lasr2)
					}
					toRemove = append(toRemove, i)

				}

			}

		case RIGHT:
			{
				x := l.field[l.lasers[i].coordinates[0]][l.lasers[i].coordinates[1]]
				switch x {
				case '#', '-':
					if l.lasers[i].coordinates[1] < len(l.field[0])-1 {
						l.lasers[i].coordinates[1]++
					}
				case '/':
					l.lasers[i].direction = UP
					if l.lasers[i].coordinates[0] > 0 {
						l.lasers[i].coordinates[0]--
					}
				case '\\':
					l.lasers[i].direction = DOWN
					if l.lasers[i].coordinates[0] < len(l.field)-1 {
						l.lasers[i].coordinates[0]++
					}
				case '|':
					lasr := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0] + 1,
							l.lasers[i].coordinates[1],
						},
						direction: DOWN}
					lasr2 := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0] - 1,
							l.lasers[i].coordinates[1],
						},
						direction: UP,
					}
					if lasr.coordinates[0] < len(l.field) {
						l.lasers = append(l.lasers, lasr)
					}
					if lasr2.coordinates[0] >= 0 {
						l.lasers = append(l.lasers, lasr2)
					}
					toRemove = append(toRemove, i)

				}

			}
		case UP:
			{
				x := l.field[l.lasers[i].coordinates[0]][l.lasers[i].coordinates[1]]
				switch x {
				case '#', '|':
					if l.lasers[i].coordinates[0] > 0 {
						l.lasers[i].coordinates[0]--
					}
				case '/':
					l.lasers[i].direction = RIGHT
					if l.lasers[i].coordinates[1] < len(l.field[0])-1 {
						l.lasers[i].coordinates[1]++
					}
				case '\\':
					l.lasers[i].direction = LEFT
					if l.lasers[i].coordinates[1] > 0 {

						l.lasers[i].coordinates[1]--
					}
				case '-':
					lasr := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0],
							l.lasers[i].coordinates[1] + 1,
						},
						direction: RIGHT}
					lasr2 := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0],
							l.lasers[i].coordinates[1] - 1,
						},
						direction: LEFT,
					}
					if lasr.coordinates[1] < len(l.field[0]) {
						l.lasers = append(l.lasers, lasr)
					}
					if lasr2.coordinates[1] >= 0 {
						l.lasers = append(l.lasers, lasr2)
					}
					toRemove = append(toRemove, i)

				}

			}
		case DOWN:
			{
				x := l.field[l.lasers[i].coordinates[0]][l.lasers[i].coordinates[1]]
				switch x {
				case '#', '|':
					if l.lasers[i].coordinates[0] < len(l.field)-1 {
						l.lasers[i].coordinates[0]++
					}
				case '\\':
					l.lasers[i].direction = RIGHT
					if l.lasers[i].coordinates[1] < len(l.field[0])-1 {
						l.lasers[i].coordinates[1]++
					}
				case '/':
					l.lasers[i].direction = LEFT
					if l.lasers[i].coordinates[0] > 0 {

						l.lasers[i].coordinates[1]--
					}
				case '-':
					lasr := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0],
							l.lasers[i].coordinates[1] + 1,
						},
						direction: RIGHT}
					lasr2 := laser{
						coordinates: [2]int{
							l.lasers[i].coordinates[0],
							l.lasers[i].coordinates[1] - 1,
						},
						direction: LEFT,
					}
					if lasr.coordinates[1] < len(l.field[0]) {
						l.lasers = append(l.lasers, lasr)
					}
					if lasr2.coordinates[1] >= 0 {
						l.lasers = append(l.lasers, lasr2)
					}
					toRemove = append(toRemove, i)

				}

			}

		}
	}
	sort.Slice(toRemove, func(i, j int) bool {
		return i < j
	})
	for i := range toRemove {
		l.lasers = append(l.lasers[:toRemove[i]], l.lasers[toRemove[i]+1:]...)
		for j := range toRemove[i+1:] {
			toRemove[j]--
		}
	}
}

func part1(start laser) int {
	lines := parse()

	lasr := start
	laserss := []laser{lasr}
	m := map[[2]int]bool{}
	lasrs := lasers{lasers: laserss, field: lines, usedPipes: m}
	// fmt.Println(lasrs.field)
	var field, fieldPrev []string
	field = lasrs.field
	fieldPrev = []string{"f"}
	for i := 0; i < 1400; i++ {

		fieldPrev = field
		lasrs.move()
		field = lasrs.field
		// for _, line := range lasrs.field {
		// 	println(line)
		// }

	}
	for strings.Join(field, " ") != strings.Join(fieldPrev, " ") {

		fieldPrev = field
		lasrs.move()
		field = lasrs.field
		// for _, line := range lasrs.field {
		// 	println(line)
		// }

	}
	// fmt.Println(
	// 	lasrs.lasers,
	// )
	sum := 0
	// for key, value := range lasrs.usedPipes {
	// 	if value {

	// 		sum++
	// 		fmt.Println(key)
	// 	}
	// }

	sln := []string{}
	for i, line := range lasrs.field[:len(lasrs.field)-1] {
		sln = append(sln, "")
		for j := range line[:len(line)-1] {
			if lasrs.usedPipes[[2]int{i, j}] {
				sln[i] += "#"
				sum++
			} else {
				sln[i] += string(line[j])
			}
		}
	}
	// for _, line := range sln {
	// 	fmt.Println(line)
	// }
	return sum

}

func part2() {
	sums := []int{}
	for i := 0; i < 110; i++ {
		lasr := laser{direction: DOWN, coordinates: [2]int{0, i}}
		sums = append(sums, part1(lasr))
	}
	for i := 0; i < 110; i++ {
		lasr := laser{direction: UP, coordinates: [2]int{110, i}}
		sums = append(sums, part1(lasr))
	}
	for i := 0; i < 110; i++ {
		lasr := laser{direction: LEFT, coordinates: [2]int{i, 110}}
		sums = append(sums, part1(lasr))
	}
	for i := 0; i < 110; i++ {
		lasr := laser{direction: RIGHT, coordinates: [2]int{i, 0}}
		sums = append(sums, part1(lasr))
	}
	max := 0
	for _, i := range sums {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)

}
func main() {
	lasr := laser{direction: RIGHT, coordinates: [2]int{0, 0}}
	part1(lasr)
	part2()

}
