package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	hands, bids := parse()
	ranks := []int{}
	for _, str := range hands {
		ranks = append(ranks, handRank(str))
	}
	sort(&hands, &bids, &ranks)

	sum := 0
	for i, n := range bids {
		sum += n * (i + 1)
	}

	fmt.Println(sum)
}

func part2() {
	hands, bids := parse()
	ranks := []int{}
	for _, str := range hands {
		ranks = append(ranks, handRankPartTwo(str))

	}
	sortPartTwo(&hands, &bids, &ranks)
	sum := 0
	for i, n := range bids {
		sum += n * (i + 1)
	}
	fmt.Println(sum)
}

func handRankPartTwo(str string) (rank int) {
	char := str[0]
	//five of a kind
	if strings.Count(str, string(char)) == 5 {
		return 6
	}

	char2 := str[1]
	//four of a kind but either the 4 or the 1 is a joker
	if (strings.Count(str, string(char)) == 4 || strings.Count(str, string(char2)) == 4) && strings.Count(str, "J") > 0 {
		return 6
	}
	//4 of a kind
	if strings.Count(str, string(char)) == 4 || strings.Count(str, string(char2)) == 4 {
		return 5
	}

	var twoCard rune
	var threeCard rune
	twoCardsTrue := false
	threeCardsTrue := false

	for _, c := range str {
		if strings.Count(str, string(c)) == 2 {
			twoCardsTrue = true
			twoCard = c
		} else if strings.Count(str, string(c)) == 3 {
			threeCardsTrue = true
			threeCard = c
		}
	}
	if twoCardsTrue && threeCardsTrue {
		//full house
		if twoCard == 'J' || threeCard == 'J' {
			//full house but the 2 or the 3 is a joker so 5 of a kind
			return 6
		}

	}

	if twoCardsTrue && threeCardsTrue {
		//full house
		return 4
	} else if threeCardsTrue {
		//3 of a kind and no 2 of a kind but there's 1 or 3 Jokers so 4 of a kind
		if strings.Count(str, "J") > 0 {
			return 5
		}
		//3 of a kind no funny business
		return 3
	}

	var oneCheckChar rune
	oneCheckBool := false
	for _, c := range str {
		if strings.Count(str, string(c)) == 2 && !oneCheckBool {
			oneCheckBool = true
			oneCheckChar = c
		} else if strings.Count(str, string(c)) == 2 && oneCheckBool && c != oneCheckChar {

			//aka two pair
			if strings.Count(str, "J") == 2 {
				//two pair but one of the pair is a joker so 4 of a kind
				return 5
			}
			if strings.Count(str, "J") == 1 {
				//two pair but the extra card is a joker so full house
				return 4
			}
			//two pair no funny business
			return 2
		}
	}

	if oneCheckBool {

		//one pair
		if strings.Count(str, "J") > 0 {
			//one pair but either the pair or one of the others is a joker, either way 3 of a kind
			return 3
		}
		//one pair no funny business
		return 1
	}

	//high card but a joker so one pair
	if strings.Count(str, "J") > 0 {
		return 1
	}
	//high card
	return 0

}

func sortPartTwo(hands *[]string, bids *[]int, ranks *[]int) {
	cardStrength := map[string]int{}
	cardStrength["A"] = 14
	cardStrength["K"] = 13
	cardStrength["Q"] = 12
	cardStrength["J"] = 1
	cardStrength["T"] = 10
	cardStrength["9"] = 9
	cardStrength["8"] = 8
	cardStrength["7"] = 7
	cardStrength["6"] = 6
	cardStrength["5"] = 5
	cardStrength["4"] = 4
	cardStrength["3"] = 3
	cardStrength["2"] = 2

	for i := 0; i < len(*hands); i++ {
		for j := i; j > 0; j-- {
			if (*ranks)[j] < (*ranks)[j-1] {
				(*ranks)[j], (*ranks)[j-1] = (*ranks)[j-1], (*ranks)[j]
				(*hands)[j], (*hands)[j-1] = (*hands)[j-1], (*hands)[j]
				(*bids)[j], (*bids)[j-1] = (*bids)[j-1], (*bids)[j]
			} else if (*ranks)[j] == (*ranks)[j-1] {
				if getSmaller((*hands)[j], (*hands)[j-1], cardStrength) == 0 {
					(*ranks)[j], (*ranks)[j-1] = (*ranks)[j-1], (*ranks)[j]
					(*hands)[j], (*hands)[j-1] = (*hands)[j-1], (*hands)[j]
					(*bids)[j], (*bids)[j-1] = (*bids)[j-1], (*bids)[j]

				} else {
					break
				}
			} else {
				break
			}
		}
	}
}

func handRank(str string) (rank int) {

	//check five of a kind
	char := str[0]
	if strings.Count(str, string(char)) == 5 {
		return 6
	}

	char2 := str[1]
	if strings.Count(str, string(char)) == 4 || strings.Count(str, string(char2)) == 4 {
		return 5
	}

	twoCardsTrue := false
	threeCardsTrue := false

	for _, c := range str {
		if strings.Count(str, string(c)) == 2 {
			twoCardsTrue = true
		} else if strings.Count(str, string(c)) == 3 {
			threeCardsTrue = true
		}
	}
	if twoCardsTrue && threeCardsTrue {
		return 4
	} else if threeCardsTrue {
		return 3
	}

	var oneCheckChar rune
	oneCheckBool := false
	for _, c := range str {
		if strings.Count(str, string(c)) == 2 && !oneCheckBool {
			oneCheckBool = true
			oneCheckChar = c
		} else if strings.Count(str, string(c)) == 2 && oneCheckBool && c != oneCheckChar {
			return 2
		}
	}
	if oneCheckBool {
		return 1
	}
	return 0
}

func getSmaller(str1 string, str2 string, cardStrength map[string]int) int {
	char1 := str1[0]
	char2 := str2[0]
	i := 1
	for char1 == char2 {
		char1 = str1[i]
		char2 = str2[i]
		i++
	}
	if cardStrength[string(char1)] < cardStrength[string(char2)] {
		return 0
	} else {
		return 1
	}
}

func sort(hands *[]string, bids *[]int, ranks *[]int) {
	cardStrength := map[string]int{}
	cardStrength["A"] = 14
	cardStrength["K"] = 13
	cardStrength["Q"] = 12
	cardStrength["J"] = 11
	cardStrength["T"] = 10
	cardStrength["9"] = 9
	cardStrength["8"] = 8
	cardStrength["7"] = 7
	cardStrength["6"] = 6
	cardStrength["5"] = 5
	cardStrength["4"] = 4
	cardStrength["3"] = 3
	cardStrength["2"] = 2
	for i := 0; i < len(*hands); i++ {
		for j := i; j > 0; j-- {
			if (*ranks)[j] < (*ranks)[j-1] {
				(*ranks)[j], (*ranks)[j-1] = (*ranks)[j-1], (*ranks)[j]
				(*hands)[j], (*hands)[j-1] = (*hands)[j-1], (*hands)[j]
				(*bids)[j], (*bids)[j-1] = (*bids)[j-1], (*bids)[j]
			} else if (*ranks)[j] == (*ranks)[j-1] {
				if getSmaller((*hands)[j], (*hands)[j-1], cardStrength) == 0 {
					(*ranks)[j], (*ranks)[j-1] = (*ranks)[j-1], (*ranks)[j]
					(*hands)[j], (*hands)[j-1] = (*hands)[j-1], (*hands)[j]
					(*bids)[j], (*bids)[j-1] = (*bids)[j-1], (*bids)[j]

				} else {
					break
				}
			} else {
				break
			}
		}
	}
}

func parse() (hands []string, bids []int) {
	content, _ := os.ReadFile("input.txt")
	strs := strings.Split(string(content), "\n")
	strs = strs[:len(strs)-1]
	for _, line := range strs {
		hand := line[:strings.Index(line, " ")]
		bid := line[strings.Index(line, " ")+1:]
		hands = append(hands, strings.Trim(hand, " "))
		bidNumber, _ := strconv.Atoi(strings.Trim(bid, " "))
		bids = append(bids, bidNumber)
	}

	return hands, bids
}

func main() {
	part1()
	part2()
}
