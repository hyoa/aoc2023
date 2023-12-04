package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"math"
	"slices"
	"strings"
)

func init() {
	DayCollection["4"] = Day4
}

type day4 struct {
	linesInput1 []string
}

func Day4(input1, input2 string) (any, any) {
	d := day4{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day4) step1() any {
	result := 0

	for _, l := range d.linesInput1 {
		lSplitted := strings.Split(l, ": ")
		card := strings.Split(lSplitted[1], " | ")

		nbWinningNumbers := getCountOfWinningNumbers(card)
		if nbWinningNumbers != 0 {
			result += int(math.Pow(2, float64(nbWinningNumbers-1)))
		}
	}

	return result
}

func (d day4) step2() any {
	result := 0
	cardResult := make(map[int]int)
	cardPile := make(map[int]int)

	for _, l := range d.linesInput1 {
		lSplitted := strings.Split(l, ": ")
		var cardNumber int
		fmt.Sscanf(lSplitted[0], "Card %d", &cardNumber)

		card := strings.Split(lSplitted[1], " | ")

		nbWinningNumbers := getCountOfWinningNumbers(card)
		cardResult[cardNumber] = 0
		cardPile[cardNumber] = 1

		if nbWinningNumbers != 0 {
			cardResult[cardNumber] = nbWinningNumbers
		}
	}

	pile := make(map[int]int)
	for n, r := range cardResult {
		pile = getCards(n, r, cardResult, cardPile)
	}

	for _, p := range pile {
		result += p
	}

	return result
}

func getCards(idx, n int, cardResult, cardPile map[int]int) map[int]int {
	for i := idx + 1; i <= idx+n; i++ {
		cardPile[i]++
		cardPile = getCards(i, cardResult[i], cardResult, cardPile)
	}

	return cardPile
}

func getCountOfWinningNumbers(card []string) int {
	winningNumbers := strings.Split(card[0], " ")
	myNumbers := strings.Split(card[1], " ")

	nbWinningNumbers := 0
	for _, wn := range winningNumbers {
		if wn == "" {
			continue
		}

		if slices.Contains(myNumbers, wn) {
			nbWinningNumbers++
		}
	}

	return nbWinningNumbers
}
