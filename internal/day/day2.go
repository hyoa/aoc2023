package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"strings"
)

func init() {
	DayCollection["2"] = Day2
}

type day2 struct {
	linesInput1 []string
}

func Day2(input1, input2 string) (any, any) {
	d := day2{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day2) step1() any {
	colorsMax := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	gamesValid := make([]int, 0)

upper:
	for _, line := range d.linesInput1 {
		nb, seqs := getGameData(line)

		for _, seq := range seqs {
			colors := strings.Split(seq, ",")

			for _, color := range colors {
				var count int
				var name string

				fmt.Sscanf(color, "%d %s", &count, &name)

				if count > colorsMax[name] {
					continue upper
				}
			}
		}

		gamesValid = append(gamesValid, nb)
	}

	score := 0
	for _, nb := range gamesValid {
		score += nb
	}

	return score
}

func (d day2) step2() any {
	type gameColor map[string]int
	gamesColor := make([]gameColor, len(d.linesInput1))

	for _, line := range d.linesInput1 {
		nb, seqs := getGameData(line)

		gamesColor[nb-1] = make(gameColor)
		for _, seq := range seqs {
			colors := strings.Split(seq, ",")

			for _, color := range colors {
				var count int
				var name string

				fmt.Sscanf(color, "%d %s", &count, &name)

				if _, ok := gamesColor[nb-1][name]; !ok {
					gamesColor[nb-1][name] = count
				} else if gamesColor[nb-1][name] < count {
					gamesColor[nb-1][name] = count
				}
			}
		}
	}

	score := 0
	for _, colors := range gamesColor {
		c := 1
		for _, count := range colors {
			c *= count
		}

		score += c
	}

	return score
}

func getGameData(line string) (int, []string) {
	var nb int
	var record string

	lineData := strings.Split(line, ": ")
	fmt.Sscanf(lineData[0], "Game %d", &nb)
	record = lineData[1]

	seqs := strings.Split(record, ";")

	return nb, seqs
}
