package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func parse() [][]int {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(content), "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool {
		return s == ""
	})
	nums := make([][]int, len(lines))
	for i, line := range lines {
		nums[i] = make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(err)
				continue
			}

			nums[i][j] = num
		}
	}
	return nums
}

//dynamic programming solution

//bestCaseScenarios holds 12 spots for each node, and each of the 12 items in the list equate to the shortest length from the start to that node,
//FOR EACH possible direction and streak of the same direction that has been traveled

// if j is the shortest length path from the start to some node, then j - value(node) is the shortest length path from the start to another node adjacent to the first
// => the shortest length path from start to node n is the minimum of the (48?) possible adjacent values + value(curnode)
//up 1-3
//down 1-3
//left 1-3
//right 1-3

func part1() {
	nums := parse()
	fmt.Println(nums)
	bestCaseScenarios := make([][][12]int, len(nums))
	for i := range bestCaseScenarios {
		bestCaseScenarios[i] = make([][12]int, len(nums[0]))
	}
	for i := range bestCaseScenarios {
		for j := range bestCaseScenarios[i] {
			for k := range bestCaseScenarios[i][j] {
				bestCaseScenarios[i][j][k] = -1
			}
		}
	}
	for i := 0; i < 12; i++ {
		bestCaseScenarios[0][0][i] = 0
	}
	traverse(nums, bestCaseScenarios, [2]int{0, 0}, 0, "D", 0)
	g := bestCaseScenarios[len(nums)-1][len(nums[0])-1][:]
	g = slices.DeleteFunc(g, func(i int) bool { return i <= 0 })
	fmt.Println(slices.Min(g))
}

// udlr
func scenarioIndex(curdirection string, curstreak int) int {
	ary := []string{"U", "D", "L", "R"}
	directionIndex := slices.Index(ary, curdirection) * 3
	return directionIndex + curstreak - 1
}

func prevNode(curnode [2]int, curdirection string) [2]int {
	switch curdirection {
	case "U":
		return [2]int{curnode[0] + 1, curnode[1]}
	case "D":
		return [2]int{curnode[0] - 1, curnode[1]}
	case "L":
		return [2]int{curnode[0], curnode[1] + 1}
	case "R":
		return [2]int{curnode[0], curnode[1] - 1}
	default:
		return [2]int{-1, -1}
	}
}
func nextNode(curnode [2]int, curdirection string) [2]int {
	switch curdirection {
	case "U":
		return [2]int{curnode[0] - 1, curnode[1]}
	case "D":
		return [2]int{curnode[0] + 1, curnode[1]}
	case "L":
		return [2]int{curnode[0], curnode[1] - 1}
	case "R":
		return [2]int{curnode[0], curnode[1] + 1}
	default:
		return [2]int{-1, -1}
	}
}

/*
recursively add the curnode's value to the curscore, update the correct spot in bestCaseScenarios, then go to each neighbor node if the curstreak isn't 3,
*/
func traverse(nums [][]int, bestCaseScenarios [][][12]int, curnode [2]int, curscore int, curdirection string, curstreak int) {

	defer func() {
		if curnode == [2]int{len(nums) - 1, len(nums[0]) - 1} && curscore%1000 == 0 {
			fmt.Println(bestCaseScenarios[len(nums)-1][len(nums[0])-1])
		}
	}()

	bcsIndex := bestCaseScenarios[curnode[0]][curnode[1]][scenarioIndex(curdirection, curstreak)]
	if curnode[0] < 0 ||
		curnode[1] < 0 ||
		curnode[0] >= len(nums) ||
		curnode[1] >= len(nums[0]) ||
		curstreak > 3 ||
		(bcsIndex > 0 && bcsIndex <= curscore+nums[curnode[0]][curnode[1]]) {
		return
	}
	if curnode == [2]int{0, 0} { // AKA if the starting node
		traverse(nums, bestCaseScenarios, [2]int{0, 1}, 0, "R", 1)
		traverse(nums, bestCaseScenarios, [2]int{1, 0}, 0, "D", 1)
	} else {
		curscore += nums[curnode[0]][curnode[1]]
		bestCaseScenarios[curnode[0]][curnode[1]][scenarioIndex(curdirection, curstreak)] = curscore
		// not the starting node, so call traverse on all adjacent nodes EXCEPT the node we just came from or if streak is 3
		prevnode := prevNode(curnode, curdirection)
		nodesToCheck := [][2]int{
			{curnode[0] - 1, curnode[1]},
			{curnode[0] + 1, curnode[1]},
			{curnode[0], curnode[1] - 1},
			{curnode[0], curnode[1] + 1},
		}
		prevIndex := slices.Index(nodesToCheck, prevnode)
		nodesToCheck = slices.Delete(nodesToCheck, prevIndex, prevIndex+1)
		dir := []string{"U", "D", "L", "R"}
		dir = slices.Delete(dir, prevIndex, prevIndex+1)
		streaks := []int{1, 1, 1}
		for i := 0; i < len(nodesToCheck); i++ {
			if nodesToCheck[i][0] < 0 ||
				nodesToCheck[i][1] < 0 ||
				nodesToCheck[i][0] >= len(nums) ||
				nodesToCheck[i][1] >= len(nums[0]) {
				streaks = slices.Delete(streaks, i, i+1)
				dir = slices.Delete(dir, i, i+1)
				nodesToCheck = slices.Delete(nodesToCheck, i, i+1)
				i--
			}
		}
		for i, n := range nodesToCheck {
			if n == nextNode(curnode, curdirection) {
				if curstreak == 3 {
					streaks = slices.Delete(streaks, i, i+1)
					dir = slices.Delete(dir, i, i+1)
					nodesToCheck = slices.Delete(nodesToCheck, i, i+1)
					break
				}
				streaks[i] = curstreak + 1

				break
			}
		}

		for i, n := range nodesToCheck {
			traverse(nums, bestCaseScenarios, n, curscore, dir[i], streaks[i])
		}
	}
}
