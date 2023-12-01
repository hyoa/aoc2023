package day

import (
	"hyoa/aoc2023/internal/utils"
	"strconv"
	"strings"
)

func init() {
	DayCollection["1"] = Day1
}

type day1 struct {
	linesInput1 []string
	linesInput2 []string
}

func Day1(input1, input2 string) (any, any) {
	d := day1{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
		linesInput2: utils.ReadTextFileLinesAsString(input2),
	}

	return d.step1(), d.step2()
}

func (d day1) step1() any {
	search := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	values := 0
	for _, line := range d.linesInput1 {
		v := getNumber(line, search)

		vInt, _ := strconv.Atoi(v)

		values += vInt
	}

	return values
}

func (d day1) step2() any {
	search := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	values := 0
	for _, line := range d.linesInput2 {
		v := getNumber(line, search)

		vInt, _ := strconv.Atoi(v)

		values += vInt
	}

	return values
}

func getNumber(lookUp string, search []string) string {
	type number struct {
		value string
		index int
	}

	lowest := number{
		value: "",
		index: 999999999,
	}

	latest := number{
		value: "",
		index: -1,
	}

	for _, s := range search {
		idxFirst := strings.Index(lookUp, s)

		if idxFirst > -1 && idxFirst < lowest.index {
			lowest.index = idxFirst
			lowest.value = s
		}

		idxLast := strings.LastIndex(lookUp, s)
		if idxLast > -1 && idxLast > latest.index {
			latest.index = idxLast
			latest.value = s
		}
	}

	var mapInt = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if _, ok := mapInt[lowest.value]; ok {
		lowest.value = mapInt[lowest.value]
	}

	if _, ok := mapInt[latest.value]; ok {
		latest.value = mapInt[latest.value]
	}

	return lowest.value + latest.value
}
