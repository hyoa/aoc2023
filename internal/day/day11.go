package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"math"
	"time"
)

func init() {
	DayCollection["11"] = Day11
}

type day11 struct {
	linesInput1 []string
}

func Day11(input1, input2 string) (any, any) {
	d := day11{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

type universeNode struct {
	v    string
	x, y int
}

type simpleUniverseNode struct {
	v string
}

type universe struct {
	nodes [][]universeNode
}

func (d day11) step1() any {
	return getForReplacement(d.linesInput1, 2)
}

func (d day11) step2() any {
	return getForReplacement(d.linesInput1, 1000000)
}

func getForReplacement(input []string, replacement int) int {
	var result int
	replacement = replacement - 1

	galaxies := make(map[int]nodeCoordinate, 0)
	for y, line := range input {
		for x, c := range line {
			if string(c) == "#" {
				galaxies[len(galaxies)+1] = nodeCoordinate{x, y}
			}
		}
	}

	farthestYGalaxy := 0
	farthestXGalaxy := 0

	for _, g := range galaxies {
		if g.y > farthestYGalaxy {
			farthestYGalaxy = g.y
		}
		if g.x > farthestXGalaxy {
			farthestXGalaxy = g.x
		}
	}

	rowIncreaseIdx := make([]int, 0)
	for idxL, line := range input {
		c := make(map[string]int)
		for _, n := range line {
			c[string(n)]++
		}

		if len(c) == 1 {
			rowIncreaseIdx = append(rowIncreaseIdx, idxL)
		}
	}

	columnIncreaseIdx := make([]int, 0)
	for i := farthestXGalaxy; i > 0; i-- {
		c := make(map[string]int)
		for j := 0; j < farthestYGalaxy+1; j++ {
			char := input[j][i]
			c[string(char)]++
		}

		if len(c) == 1 {
			columnIncreaseIdx = append(columnIncreaseIdx, i)
		}
	}

	galaxiesShiftX := make(map[int]int)
	for _, idx := range columnIncreaseIdx {
		for g := range galaxies {
			if galaxies[g].x > idx {
				galaxiesShiftX[g]++
			}
		}
	}

	galaxiesShiftY := make(map[int]int)
	for _, idx := range rowIncreaseIdx {
		for g := range galaxies {
			if galaxies[g].y > idx {
				galaxiesShiftY[g]++
			}
		}
	}

	for g := range galaxiesShiftX {
		galaxies[g] = nodeCoordinate{galaxies[g].x + galaxiesShiftX[g]*replacement, galaxies[g].y}
	}

	for g := range galaxiesShiftY {
		galaxies[g] = nodeCoordinate{galaxies[g].x, galaxies[g].y + galaxiesShiftY[g]*replacement}
	}

	var pairs [][2]nodeCoordinate

	var galaxiesSorted []nodeCoordinate
	for _, g := range galaxies {
		galaxiesSorted = append(galaxiesSorted, g)
	}

	for i := 0; i < len(galaxiesSorted); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, [2]nodeCoordinate{galaxiesSorted[i], galaxiesSorted[j]})
		}
	}

	timeStart := time.Now()

	for _, p := range pairs {
		result += findShortestDist(p[0], p[1])
	}

	fmt.Println("time", time.Since(timeStart))

	return result
}

func findShortestDist(a, b nodeCoordinate) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}
