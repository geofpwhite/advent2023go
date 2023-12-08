package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

type node struct {
	label string
	left  string
	right string
}

// lcm = | a * b | / gcd(a,b)
func _lcm(numbers []int) int {
	x := 1
	for _, num := range numbers {
		a := x
		b := num
		//euclidean algorithm
		for b != 0 {
			a, b = b, a%b
		}
		x *= num / a
	}
	return x
}

// lcm calculates the least common multiple using the formula: LCM(a, b) = |a * b| / GCD(a, b)
func lcm(a, b *big.Int) *big.Int {
	if a.Sign() == 0 || b.Sign() == 0 {
		return big.NewInt(0)
	}
	gcdAB := gcd(a, b)
	return new(big.Int).Abs(new(big.Int).Div(new(big.Int).Mul(a, b), gcdAB))
}

// findLCM calculates the LCM of a list of integers.
func findLCM(numbers []int) *big.Int {
	if len(numbers) == 0 {
		return big.NewInt(0)
	}

	result := big.NewInt(int64(numbers[0]))
	for _, num := range numbers[1:] {
		bigNum := big.NewInt(int64(num))
		result = lcm(result, bigNum)
	}

	return result
}

func part2() {
	instructions, nodes := parse()
	instructionIndex := 0
	count := 0

	curNodes := []node{}
	cycles := []int{}

	for _, node := range nodes {
		if node.label[2] == 'A' {
			curNodes = append(curNodes, node)
			cycles = append(cycles, 0)
		}
	}

	for i := range curNodes {
		instructionIndex = 0
		count = 0
		for curNodes[i].label[2] != 'Z' {
			if instructions[instructionIndex] == 'R' {
				curNodes[i] = nodes[curNodes[i].right]
			} else {
				curNodes[i] = nodes[curNodes[i].left]
			}
			count++
			cycles[i]++
			instructionIndex = (instructionIndex + 1) % len(instructions)
		}
	}
	fmt.Println(findLCM(cycles))

}

func part1() {
	instructions, nodes := parse()
	curNode := nodes["AAA"]
	count := 0
	instructionIndex := 0
	for curNode.label != "ZZZ" {
		step := instructions[instructionIndex]
		if step == 'R' {
			curNode = nodes[curNode.right]
		} else {
			curNode = nodes[curNode.left]
		}
		count++
		instructionIndex = (instructionIndex + 1) % len(instructions)
	}
	fmt.Println(count)

}

func parse() (string, map[string]node) {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	instructions := lines[0]
	nodes := map[string]node{}
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		newNode := node{}
		newNode.label = line[:strings.Index(line, "=")-1]
		newNode.left = line[strings.Index(line, "(")+1 : strings.Index(line, ",")]
		newNode.right = line[strings.Index(line, ",")+2 : strings.Index(line, ")")]
		nodes[newNode.label] = newNode
	}

	return instructions, nodes

}

func main() {
	part1()
	part2()
}
