package main

import (
	"os"
	"strconv"
	"strings"
)

func parse() []string {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		panic("error")
	}
	lines := strings.Split((string(content)), ",")
	lines = lines[:len(lines)-1]

	return lines
}

func part2() {
	lines := parse()
	boxes := [256][]string{}
	for _, line := range lines {
		var operator rune
		if strings.Contains(line, "-") {
			operator = '-'
		} else {
			operator = '='
		}
		var lensLabel string = line[:strings.Index(line, string(operator))]
		boxNumber := 0
		for _, char := range lensLabel {
			boxNumber += int(char)
			boxNumber *= 17
			boxNumber %= 256
		}

		if operator == '-' {
			if boxes[boxNumber] != nil {
				for i, lens := range boxes[boxNumber] {
					if strings.Contains(lens, lensLabel) {
						boxes[boxNumber] = append(boxes[boxNumber][:i], boxes[boxNumber][i+1:]...)
					}
				}
			}

		} else if operator == '=' {
			focalLength := line[strings.Index(line, "=")+1:]
			label := lensLabel + " " + focalLength
			hold := true
			if boxes[boxNumber] != nil {
				for i, lens := range boxes[boxNumber] {
					if strings.Contains(lens, lensLabel) {
						// boxes[boxNumber] = append(boxes[boxNumber][:i], boxes[boxNumber][i+1:]...)
						boxes[boxNumber][i] = label
						hold = false
						break
					}
				}
				if hold {
					boxes[boxNumber] = append(boxes[boxNumber], label)
				}
			} else {

				boxes[boxNumber] = append(boxes[boxNumber], label)
			}
		}

	}
	sum := 0
	for i, box := range boxes {
		for j, lens := range box {
			length, _ := strconv.Atoi(lens[strings.Index(lens, " ")+1:])
			score := (1 + i) * (1 + j) * length
			sum += score
			println(lens[:3], score, 1+i, j+1, length)
		}
	}
	println(sum)
}

func part1() {
	lines := parse()
	sum := 0
	for _, line := range lines {
		score := 0
		for _, c := range strings.Trim(line, "\n") {
			score += int(c)
			score = (score * 17) % 256
		}
		println(score)
		println(line)
		sum += score
	}
	println(sum)
}
func main() {
	part1()
	part2()
}
