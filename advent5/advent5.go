package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		print("err")
	}

	strs := strings.Split(string(content), "\n")
	// seeds := []string{}
	seedString := strings.Split(strs[0][strings.Index(strs[0], ":")+1:], " ")
	seedString = seedString[1:]
	var startString string
	boo := 0
	maps := [][]string{}
	// for _, c := range seeds {
	// 	fmt.Println(c)
	// }
	mapAry := []string{}
	for _, c := range strs[1:] {
		if strings.Contains(c, "to") && len(mapAry) > 0 {
			maps = append(maps, mapAry)
			mapAry = []string{}
			continue
		} else {
			mapAry = append(mapAry, c)
		}
	}
	maps = append(maps, mapAry)
	maps = maps[1:]
	// fmt.Println(maps)
	// fmt.Println(seeds)

	final := []int{}
	for _, c := range seedString {
		if c == "" {
			continue
		}
		if boo == 0 {
			startString = c
			boo = 1
		} else {

			boo = 0
			x, _ := strconv.Atoi(strings.Trim(startString, " "))
			y, _ := strconv.Atoi(strings.Trim(c, " "))
			// seeds = append(seeds, strconv.Itoa(x))
			// seeds = append(seeds, strconv.Itoa(x+y))
			i := x
			min := -1
			for i < x+y {
				// seeds = append(seeds, strconv.Itoa(i))
				n := doMap(&maps, i)
				if min < 0 || n < min {
					min = n
				}
				i++
			}
			println(x)
			final = append(final, min)
		}
	}

	// for _, c := range seeds {
	// 	// fmt.Println(c)
	// 	if c == "" {
	// 		continue
	// 	}
	// 	x := strings.Trim(c, " ")
	// 	num, err := strconv.Atoi(x)
	// 	// source := num
	// 	// fmt.Println(c)
	// 	if err != nil {
	// 		fmt.Println("err num ")
	// 	}
	//
	// 	num = doMap(maps, num)
	// 	final = append(final, num)
	// 	// fmt.Println(source)
	// }
	mini := final[0]
	for _, c := range final {
		if c < mini {
			mini = c
		}
	}
	fmt.Println(mini)
	// fmt.Println(seeds)
	// fmt.Println(final)
}
func doMap(maps *[][]string, source int) int {

	num := source
	for _, ary := range *maps {
		// fmt.Println(i)

		for _, str := range ary {
			if str == "" || str == " " {
				continue
			}
			destinationString := str[:strings.Index(str, " ")]
			destination, err := strconv.Atoi(strings.Trim(destinationString, " "))
			if err != nil {
				fmt.Println("err destination")
			}
			sourceString := str[strings.Index(str, " ")+1:]
			rangeString := sourceString[strings.Index(sourceString, " "):]
			_range, err := strconv.Atoi(strings.Trim(rangeString, " "))
			source, err = strconv.Atoi(sourceString[:strings.Index(sourceString, " ")])
			// fmt.Println(destination, source, _range)
			// fmt.Println(i)
			if num >= source && num < source+_range {
				num = destination + (num - source)
				// fmt.Println(num)
				break
			}

		}
	}
	return num
}

func part1() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		print("err")
	}

	strs := strings.Split(string(content), "\n")

	seeds := strings.Split(strs[0][strings.Index(strs[0], ":")+1:], " ")
	maps := [][]string{}
	// for _, c := range seeds {
	// 	fmt.Println(c)
	// }
	mapAry := []string{}
	for _, c := range strs[1:] {
		if strings.Contains(c, "to") && len(mapAry) > 0 {
			maps = append(maps, mapAry)
			mapAry = []string{}
			continue
		} else {
			mapAry = append(mapAry, c)
		}
	}
	maps = append(maps, mapAry)
	fmt.Println(maps)
	maps = maps[1:]

	final := []int{}

	for _, c := range seeds {
		// fmt.Println(c)
		if c == "" {
			continue
		}
		x := strings.Trim(c, " ")
		num, err := strconv.Atoi(x)
		// fmt.Println(c)
		if err != nil {
			fmt.Println("err num ")
		}

		// fmt.Println(maps)
		for i, ary := range maps {
			fmt.Println(i)

			for _, str := range ary {
				if str == "" {
					continue
				}
				destinationString := str[:strings.Index(str, " ")]
				destination, err := strconv.Atoi(destinationString)
				if err != nil {
					fmt.Println("err destination")
				}
				sourceString := str[strings.Index(str, " ")+1:]
				source, err := strconv.Atoi(strings.Trim(sourceString, " "))
				rangeString := sourceString[strings.Index(sourceString, " "):]
				_range, err := strconv.Atoi(strings.Trim(rangeString, " "))
				source, err = strconv.Atoi(sourceString[:strings.Index(sourceString, " ")])
				fmt.Println(destination, source, _range)
				fmt.Println(i)
				if num >= source && num < source+_range {
					num = destination + (num - source)
					fmt.Println(num)
					break
				}

			}
		}
		final = append(final, num)
	}
	fmt.Println(len(final))
	mini := final[0]
	for _, c := range final {
		fmt.Println(c)
		if c < mini {
			mini = c
		}
	}
	fmt.Println(mini)
	fmt.Println(seeds)
}

func main() {
	part2()
}
