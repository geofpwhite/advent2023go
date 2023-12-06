package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	times, distances := parse()
	multipliers := [4]int{0, 0, 0, 0}

	for index, time := range times {
		distanceToBeat := distances[index]
		for i := 0; i < time; i++ {
			curSpeed := 0
			// for j = 0; j < i; j++ {
			// 	curSpeed++
			//
			// }
			curSpeed += i
			distanceCheck := 0
			// for ; j < time; j++ {
			// 	distanceCheck += curSpeed
			// }
			distanceCheck += (curSpeed * (time - i))
			if distanceCheck > distanceToBeat {
				multipliers[index]++
			}

		}
	}
	num := 1
	for i := 0; i < 4; i++ {
		num *= multipliers[i]
	}
	fmt.Println(num)
}

func part2() {
	times, distances := parse()
	timeString := ""
	distanceString := ""
	waysToBeat := 0

	for _, num := range times {
		timeString += strconv.Itoa(num)
	}
	for _, num := range distances {
		distanceString += strconv.Itoa(num)
	}
	time, _ := strconv.Atoi(timeString)
	distance, _ := strconv.Atoi(distanceString)
	for i := 0; i < time; i++ {
		curSpeed := 0
		// for j = 0; j < i; j++ {
		// 	curSpeed++
		//
		// }
		curSpeed += i
		distanceCheck := 0
		// for ; j < time; j++ {
		// 	distanceCheck += curSpeed
		// }

		distanceCheck += (curSpeed * (time - i))
		if distanceCheck > distance {
			waysToBeat++
		}

	}
	fmt.Println(waysToBeat)
}

func parse() ([]int, []int) {
	content, _ := os.ReadFile("input.txt")
	strs := strings.Split(string(content), "\n")
	fmt.Println(strs)
	times := []int{}
	distances := []int{}
	for i, c := range strs {
		str := c[strings.Index(c, ":")+1:]
		str = strings.Trim(str, " ")
		fmt.Println(str)
		if i == 0 {
			x, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			str = str[strings.Index(str, " "):]
			str = strings.Trim(str, " ")
			fmt.Println(x)
			y, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			str = str[strings.Index(str, " "):]
			str = strings.Trim(str, " ")
			fmt.Println(y)
			z, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			str = str[strings.Index(str, " "):]
			str = strings.Trim(str, " ")
			fmt.Println(z)
			str = str + " "
			w, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			fmt.Println(w)
			times = append(times, x, y, z, w)

		} else {
			if str == "" {
				break
			}
			x, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			str = str[strings.Index(str, " "):]
			str = strings.Trim(str, " ")
			fmt.Println(x)
			y, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			str = str[strings.Index(str, " "):]
			str = strings.Trim(str, " ")
			fmt.Println(y)
			z, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			str = str[strings.Index(str, " "):]
			str = strings.Trim(str, " ")
			fmt.Println(z)
			str = str + " "
			w, _ := strconv.Atoi(str[:strings.Index(str, " ")])
			fmt.Println(w)
			distances = append(distances, x, y, z, w)

		}

	}
	return times, distances
}

func main() {
	// part1()
	part2()
}
