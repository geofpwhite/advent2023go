package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0
	red := 12
	green := 13
	blue := 14

	ar := strings.Split(string(content), "\n")
	// reds := []int{}
	// greens := []int{}
	// blues := []int{}

	// valid := []int{}
	for _, c := range ar {

		s := string(c)
		substrings := strings.Split(s, ",")
		if !strings.Contains(substrings[0], "Game") {
			break
		}
		gameNumberString := substrings[0][strings.Index(substrings[0], "Game")+5 : strings.Index(substrings[0], ":")]
		otherString := substrings[0][strings.Index(substrings[0], ":")+1:]
		substrings = append(substrings, otherString)
		gameNumber, err := strconv.Atoi(gameNumberString)

		if err != nil {

		}
		valid := true
		for _, str := range substrings {
			if strings.Contains(str, ";") {
				twoStrings := strings.Split(str, ";")
				for _, iStr := range twoStrings {
					if strings.Contains(iStr, "red") {
						numberReds, err := strconv.Atoi(iStr[1 : strings.Index(iStr, "red")-1])

						if err != nil {
							// continue
						}
						if numberReds > red {
							valid = false
						}

					} else if strings.Contains(iStr, "green") {
						numberGreens, err := strconv.Atoi(iStr[1 : strings.Index(iStr, "green")-1])

						if err != nil {

						}
						if numberGreens > green {
							valid = false
						}

					} else if strings.Contains(iStr, "blue") {
						numberBlues, err := strconv.Atoi(iStr[1 : strings.Index(iStr, "blue")-1])
						if err != nil {

						}
						if numberBlues > blue {
							valid = false
						}

					}

				}
			} else {
				if strings.Contains(str, "red") {

					numberReds, err := strconv.Atoi(str[1 : strings.Index(str, "red")-1])

					if err != nil {
						// continue
					}
					if numberReds > red {
						valid = false
					}
				} else if strings.Contains(str, "green") {
					numberGreens, err := strconv.Atoi(str[1 : strings.Index(str, "green")-1])

					if err != nil {

					}
					if numberGreens > green {
						valid = false
					}

				} else if strings.Contains(str, "blue") {
					numberBlues, err := strconv.Atoi(str[1 : strings.Index(str, "blue")-1])
					if err != nil {

					}
					if numberBlues > blue {
						valid = false
					}

				}
			}
		}
		if valid {
			sum += gameNumber
		}

	}
	fmt.Println(sum)
}
func part2() {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	sum := 0

	ar := strings.Split(string(content), "\n")
	// reds := []int{}
	// greens := []int{}
	// blues := []int{}

	// valid := []int{}
	for _, c := range ar {

		red := 0
		green := 0
		blue := 0
		s := string(c)
		substrings := strings.Split(s, ",")
		if !strings.Contains(substrings[0], "Game") {
			break
		}
		// gameNumberString := substrings[0][strings.Index(substrings[0], "Game")+5 : strings.Index(substrings[0], ":")]
		otherString := substrings[0][strings.Index(substrings[0], ":")+1:]
		substrings = append(substrings, otherString)
		// gameNumber, err := strconv.Atoi(gameNumberString)

		if err != nil {

		}
		for _, str := range substrings {
			if strings.Contains(str, ";") {
				twoStrings := strings.Split(str, ";")
				for _, iStr := range twoStrings {
					if strings.Contains(iStr, "red") {
						numberReds, err := strconv.Atoi(iStr[1 : strings.Index(iStr, "red")-1])

						if err != nil {
							// continue
						}
						if numberReds > red {
							red = numberReds
						}

					} else if strings.Contains(iStr, "green") {
						numberGreens, err := strconv.Atoi(iStr[1 : strings.Index(iStr, "green")-1])

						if err != nil {

						}
						if numberGreens > green {
							green = numberGreens
						}

					} else if strings.Contains(iStr, "blue") {
						numberBlues, err := strconv.Atoi(iStr[1 : strings.Index(iStr, "blue")-1])
						if err != nil {

						}
						if numberBlues > blue {
							blue = numberBlues
						}

					}

				}
			} else {
				if strings.Contains(str, "red") {

					numberReds, err := strconv.Atoi(str[1 : strings.Index(str, "red")-1])

					if err != nil {
						// continue
					}
					if numberReds > red {
						red = numberReds
					}
				} else if strings.Contains(str, "green") {
					numberGreens, err := strconv.Atoi(str[1 : strings.Index(str, "green")-1])

					if err != nil {

					}
					if numberGreens > green {
						green = numberGreens
					}

				} else if strings.Contains(str, "blue") {
					numberBlues, err := strconv.Atoi(str[1 : strings.Index(str, "blue")-1])
					if err != nil {

					}
					if numberBlues > blue {
						blue = numberBlues
					}

				}
			}
		}
		sum += (red * green * blue)

	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
