package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"strings"
)

func init() {
	DayCollection["8"] = Day8
}

type day8 struct {
	linesInput1 []string
	linesInput2 []string
}

func Day8(input1, input2 string) (any, any) {
	d := day8{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
		linesInput2: utils.ReadTextFileLinesAsString(input2),
	}

	return d.step1(), d.step2()
}

type move struct {
	right          string
	left           string
	isStart, isEnd bool
}

func (d day8) step1() any {
	moves := strings.Split(d.linesInput1[0], "")

	dest := make(map[string]move)

	for _, l := range d.linesInput1[2:] {
		var idx, right, left string
		fmt.Sscanf(l, "%s = (%s%s", &idx, &left, &right)
		right = strings.Trim(right, ")")
		left = strings.Trim(left, ",")
		dest[idx] = move{right, left, false, false}
	}

	run := 0
	idxMove := 0
	current := "AAA"
	for {
		run++
		switch moves[idxMove] {
		case "R":
			current = dest[current].right
		case "L":
			current = dest[current].left
		}

		if current == "ZZZ" {
			break
		}

		idxMove = (idxMove + 1) % len(moves)
	}

	return run
}

func (d day8) step2() any {
	moves := strings.Split(d.linesInput2[0], "")

	nodes := make(map[string]move)
	starts := make([]string, 0)

	for _, l := range d.linesInput2[2:] {
		var idx, right, left string
		fmt.Sscanf(l, "%s = (%s%s", &idx, &left, &right)
		right = strings.Trim(right, ")")
		left = strings.Trim(left, ",")

		r := []rune(idx)
		lastChar := r[len(r)-1]

		var isStart, isEnd bool
		if string(lastChar) == "A" {
			isStart = true
			starts = append(starts, idx)
		}

		if string(lastChar) == "Z" {
			isEnd = true
		}

		nodes[idx] = move{right, left, isStart, isEnd}
	}

	res := make([]int, 0)
	for k := range starts {
		n := findEnd(starts[k], moves, nodes, 0)
		res = append(res, n)
	}

	return ppmcMultiple(res)
}

func pgdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func ppmc(a, b int) int {
	return a * b / pgdc(a, b)
}

func ppmcMultiple(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = ppmc(result, number)
	}

	return result
}

func findEnd(current string, moves []string, nodes map[string]move, idxMove int) int {
	var run int
	for {
		run++
		switch moves[idxMove] {
		case "R":
			current = nodes[current].right
		case "L":
			current = nodes[current].left
		}

		idxMove = (idxMove + 1) % len(moves)

		if nodes[current].isEnd {
			break
		}
	}

	return run
}
