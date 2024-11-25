package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
)

func init() {
	DayCollection["13"] = Day13
}

type day13 struct {
	linesInput1 []string
}

func Day13(input1, input2 string) (any, any) {
	d := day13{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day13) step1() any {
	inputs := make([][]string, 0)

	idx := 0
	inputs = append(inputs, make([]string, 0))
	for _, l := range d.linesInput1 {
		if l == "" {
			idx++
			inputs = append(inputs, make([]string, 0))
			continue
		}

		inputs[idx] = append(inputs[idx], l)
	}

	// for _, input := range inputs {

	// }
	current := inputs[0]
	columns := make([]string, len(current[0]))
	for y := 0; y < len(current); y++ {
		for x := 0; x < len(current[y]); x++ {
			columns[x] += string(current[y][x])
		}
	}

	c := split(columns)

	fmt.Println(c)

	return 0
}

func split(slice []string) int {
	idxSplit := 1

	for {
		if idxSplit == len(slice)-1 {
			break
		}

		sep := idxSplit
		for i := 0; i <= idxSplit; i++ {
		}
	}

	return 0
}

func (d day13) step2() any {
	return 0
}

// func split(slice []string) int {
// 	idxStart := 0
// 	idxEnd := len(slice) - 1

// 	removedItem := 0
// 	for {
// 		if idxEnd == 0 {
// 			slice = slice[idxStart+1:]
// 			removedItem++
// 		}

// 		if idxStart == idxEnd {
// 			break
// 		}

// 		if slice[idxStart] == slice[idxEnd] {
// 			idxStart++
// 			idxEnd--
// 		}
// 		idxEnd--
// 	}

// 	return len(slice)/2 + removedItem
// }

// func debug() {

// 	idxStart := 0
// 	idxEnd := len(columns) - 1
// 	removedColumn := 0
// 	for {
// 		if idxEnd == 0 {
// 			columns = columns[idxStart+1:]
// 			removedColumn++
// 		}

// 		if idxStart == idxEnd {
// 			break
// 		}

// 		if columns[idxStart] == columns[idxEnd] {
// 			idxStart++
// 			idxEnd--
// 		}

// 		idxEnd--
// 	}
// 	fmt.Println(len(columns)/2, removedColumn)

// 	rows := inputs[1]
// 	idxStartRow := 0
// 	idxEndRow := len(rows) - 1

// 	removedRows := 0
// 	for {
// 		if idxEndRow == 0 {
// 			columns = columns[idxStartRow+1:]
// 			removedRows++
// 		}

// 		if idxStartRow == idxEndRow {
// 			break
// 		}

// 		if columns[idxStartRow] == columns[idxEndRow] {
// 			idxStartRow++
// 			idxEndRow--
// 		}
// 		idxEndRow--
// 	}

// 	fmt.Println(len(rows)/2, removedRows)
// }
