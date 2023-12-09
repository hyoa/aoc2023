package day

import (
	"hyoa/aoc2023/internal/utils"
	"strconv"
	"strings"
)

func init() {
	DayCollection["9"] = Day9
}

type day9 struct {
	linesInput1 []string
}

func Day9(input1, input2 string) (any, any) {
	d := day9{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day9) step1() any {
	var result int
	histories := createHistories(d.linesInput1)
	allSeqs := createSeqs(histories)

	for k := range allSeqs {
		var s int
		cS := allSeqs[k]
		for i := len(cS) - 2; i >= 0; i-- {
			s += cS[i][len(cS[i])-1]
		}

		result += s
	}

	return result
}

func (d day9) step2() any {
	var result int
	histories := createHistories(d.linesInput1)
	allSeqs := createSeqs(histories)

	for k := range allSeqs {
		var s int
		cS := allSeqs[k]
		for i := len(cS) - 2; i >= 0; i-- {
			s = cS[i][0] - s
		}

		result += s
	}

	return result
}

func createHistories(input []string) [][]int {
	var histories [][]int

	for _, line := range input {
		var history []int
		v := strings.Split(line, " ")
		for _, s := range v {
			n, _ := strconv.Atoi(s)
			history = append(history, n)
		}
		histories = append(histories, history)
	}

	return histories
}

func createSeqs(histories [][]int) [][][]int {
	allSeqs := make([][][]int, 0)

	for k := range histories {
		s := getChildSeq(histories[k], make([][]int, 0))
		allSeqs = append(allSeqs, s)
	}

	return allSeqs
}

func getChildSeq(history []int, seqs [][]int) [][]int {
	if len(seqs) == 0 {
		seqs = append(seqs, history)
	}

	seq := make([]int, 0)

	nOfZero := 0
	for i := 0; i < len(history)-1; i++ {
		n := history[i+1] - history[i]
		seq = append(seq, n)

		if n == 0 {
			nOfZero++
		}
	}

	seqs = append(seqs, seq)

	if nOfZero == len(history)-1 {
		return seqs
	}

	return getChildSeq(seq, seqs)
}
