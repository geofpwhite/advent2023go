package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// part1()
	// part1()
	// part1recursive()
	part2()

}

var nums [][]int
var bestCaseScenarios [][][12]int
var bestCaseScenariosPart2 [][][40]int
var nodeQueue [][2]int = [][2]int{}
var scoreQueue []int = []int{}
var directionQueue []string = []string{}
var streakQueue []int = []int{}

func parse() [][]int {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
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

func part1recursive() {
	nums = parse()
	fmt.Println(nums)
	bestCaseScenarios = make([][][12]int, len(nums))
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

	traverse([2]int{0, 0}, 0, "D", 0)
	g := bestCaseScenarios[len(nums)-1][len(nums[0])-1][:]
	g = slices.DeleteFunc(g, func(i int) bool { return i <= 0 })
	fmt.Println(slices.Min(g))
}
func part1() {
	nums = parse()
	fmt.Println(nums)
	bestCaseScenarios = make([][][12]int, len(nums))
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

	push([2]int{0, 0}, 0, "D", 0)
	iterativeTraverse()
	g := bestCaseScenarios[len(nums)-1][len(nums[0])-1][:]
	g = slices.DeleteFunc(g, func(i int) bool { return i <= 0 })
	fmt.Println(slices.Min(g))
}
func part2() {
	nums = parse()
	fmt.Println(nums)
	bestCaseScenariosPart2 = make([][][40]int, len(nums))
	for i := range bestCaseScenariosPart2 {
		bestCaseScenariosPart2[i] = make([][40]int, len(nums[0]))
	}
	for i := range bestCaseScenariosPart2 {
		for j := range bestCaseScenariosPart2[i] {
			for k := range bestCaseScenariosPart2[i][j] {
				bestCaseScenariosPart2[i][j][k] = -1
			}
		}
	}
	for i := 0; i < 40; i++ {
		bestCaseScenariosPart2[0][0][i] = 0
	}

	push([2]int{0, 0}, 0, "D", 0)
	traverse2()
	g := bestCaseScenariosPart2[len(nums)-1][len(nums[0])-1][:]
	g = slices.DeleteFunc(g, func(i int) bool { return i == -1 })
	fmt.Println(slices.Min(g))
}

// udlr
func scenarioIndex(curdirection string, curstreak int) int {
	ary := []string{"U", "D", "L", "R"}
	directionIndex := slices.Index(ary, curdirection) * 3
	return directionIndex + curstreak - 1
}
func scenarioIndex2(curdirection string, curstreak int) int {
	ary := []string{"U", "D", "L", "R"}
	directionIndex := slices.Index(ary, curdirection) * 10
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

func pop() ([2]int, int, string, int) {
	x, y, z, w := nodeQueue[0], scoreQueue[0], directionQueue[0], streakQueue[0]
	// nodeQueue = slices.Delete(nodeQueue, 0, 1)
	nodeQueue = nodeQueue[1:]
	scoreQueue = scoreQueue[1:]
	directionQueue = directionQueue[1:]
	streakQueue = streakQueue[1:]
	// scoreQueue = slices.Delete(scoreQueue, 0, 1)
	//
	// directionQueue = slices.Delete(directionQueue, 0, 1)
	// streakQueue = slices.Delete(streakQueue, 0, 1)
	return x, y, z, w
}

func push(curnode [2]int, curscore int, curdirection string, curstreak int) {
	nodeQueue = append(nodeQueue, curnode)
	scoreQueue = append(scoreQueue, curscore)
	directionQueue = append(directionQueue, curdirection)
	streakQueue = append(streakQueue, curstreak)
}

/*
recursively add the curnode's value to the curscore, update the correct spot in bestCaseScenarios, then go to each neighbor node if the curstreak isn't 3,
*/

func iterativeTraverse() {
	for len(nodeQueue) > 0 {
		curnode, curscore, curdirection, curstreak := pop()
		bcsIndex := bestCaseScenarios[curnode[0]][curnode[1]][scenarioIndex(curdirection, curstreak)]
		if curstreak > 3 ||
			(bcsIndex > 0 && bcsIndex <= curscore+nums[curnode[0]][curnode[1]]) {
			continue
		}
		if curnode == [2]int{0, 0} { // AKA if the starting node
			if curscore == 0 {
				push([2]int{0, 1}, 0, "R", 1)
				push([2]int{1, 0}, 0, "D", 1)
			}
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
				push(n, curscore, dir[i], streaks[i])
			}
		}

	}

}

func traverse(curnode [2]int, curscore int, curdirection string, curstreak int) {

	defer func() {
		if curnode == [2]int{len(nums) - 1, len(nums[0]) - 1} && curscore%1000 == 0 {
			fmt.Println(bestCaseScenarios[len(nums)-1][len(nums[0])-1])
		}
	}()

	bcsIndex := bestCaseScenarios[curnode[0]][curnode[1]][scenarioIndex(curdirection, curstreak)]
	if curstreak > 3 ||
		(bcsIndex > 0 && bcsIndex <= curscore+nums[curnode[0]][curnode[1]]) {
		return
	}
	if curnode == [2]int{0, 0} { // AKA if the starting node
		if curscore == 0 {
			traverse([2]int{0, 1}, 0, "R", 1)
			traverse([2]int{1, 0}, 0, "D", 1)
		}
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
			traverse(n, curscore, dir[i], streaks[i])
		}
	}
}

func traverse2() {
	// defer func() {
	// 	if curnode == [2]int{len(nums) - 1, len(nums[0]) - 1} {
	// 		fmt.Println(bestCaseScenariosPart2[curnode[0]][curnode[1]])
	// 	}
	// }()
	for len(nodeQueue) > 0 {
		curnode, curscore, curdirection, curstreak := pop()
		bcsIndex := bestCaseScenariosPart2[curnode[0]][curnode[1]][scenarioIndex2(curdirection, curstreak)]
		xlength, ylength := len(nums), len(nums[0])
		if curstreak > 10 || curnode[0] < 0 || curnode[1] < 0 || curnode[0] > xlength-1 || curnode[1] > ylength-1 ||
			(bcsIndex > 0 && bcsIndex <= curscore+nums[curnode[0]][curnode[1]]) {
			continue
		}
		if curnode == [2]int{0, 0} { // AKA if the starting node
			if curscore == 0 {
				push([2]int{0, 1}, 0, "R", 1)
				push([2]int{1, 0}, 0, "D", 1)
			}
		} else {
			next := nextNode(curnode, curdirection)
			// if curstreak == 1 {
			// 	next2 := nextNode(next, curdirection)
			// 	next3 := nextNode(next2, curdirection)
			// 	if next3[0] < 0 || next3[0] >= xlength || next3[1] < 0 || next3[1] >= xlength {
			// 		return
			// 	}
			// 	traverse2(next3, curscore+nums[curnode[0]][curnode[1]]+nums[next[0]][next[1]]+nums[next2[0]][next2[1]], curdirection, 4)
			// } else
			if curstreak < 4 && next[0] >= 0 && next[1] >= 0 && next[0] < xlength && next[1] < ylength {
				curscore += nums[curnode[0]][curnode[1]]
				bestCaseScenariosPart2[curnode[0]][curnode[1]][scenarioIndex2(curdirection, curstreak)] = curscore
				push(next, curscore, curdirection, curstreak+1)
			} else if curstreak >= 4 {

				curscore += nums[curnode[0]][curnode[1]]
				bestCaseScenariosPart2[curnode[0]][curnode[1]][scenarioIndex2(curdirection, curstreak)] = curscore
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
						nodesToCheck[i][0] >= xlength ||
						nodesToCheck[i][1] >= ylength {
						streaks = slices.Delete(streaks, i, i+1)
						dir = slices.Delete(dir, i, i+1)
						nodesToCheck = slices.Delete(nodesToCheck, i, i+1)
						i--
					}
				}
				for i, n := range nodesToCheck {
					if n == nextNode(curnode, curdirection) {
						if curstreak == 10 {
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
					push(n, curscore, dir[i], streaks[i])
				}
			}
		}
	}
}
