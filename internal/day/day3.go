package day

import (
	"hyoa/aoc2023/internal/utils"
	"slices"
	"strconv"
	"strings"
)

func init() {
	DayCollection["3"] = Day3
}

type day3 struct {
	linesInput1 []string
}

func Day3(input1, input2 string) (any, any) {
	d := day3{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day3) step1() any {
	sum := 0

	type pos struct {
		xS, xE, y int
	}

	numbers := make(map[pos]int)

	toIgnore := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}
	for y, line := range d.linesInput1 {
		splitted := strings.Split(line, "")

		for x, s := range splitted {
			if !slices.Contains(toIgnore, s) {
				// above
				nT := 0
				aboveLine := []rune(d.linesInput1[y-1])
				if _, err := strconv.Atoi(string(aboveLine[x])); err == nil {
					xL := lookLeft(aboveLine, x)
					xR := lookRight(aboveLine, x)

					nT, _ = strconv.Atoi(string(aboveLine[xL+1 : xR]))

					numbers[pos{
						xS: xL + 1,
						xE: xR,
						y:  y - 1,
					}] = nT
				}

				// below
				nB := 0
				belowLine := []rune(d.linesInput1[y+1])
				if _, err := strconv.Atoi(string(belowLine[x])); err == nil {
					xL := lookLeft(belowLine, x)
					xR := lookRight(belowLine, x)

					nB, _ = strconv.Atoi(string(belowLine[xL+1 : xR]))

					numbers[pos{
						xS: xL + 1,
						xE: xR,
						y:  y + 1,
					}] = nB
				}

				// left
				nL := 0
				leftLine := []rune(d.linesInput1[y])
				if _, err := strconv.Atoi(string(leftLine[x-1])); err == nil {
					xL := lookLeft(leftLine, x-1)
					xR := x

					nL, _ = strconv.Atoi(string(leftLine[xL+1 : xR]))

					numbers[pos{
						xS: xL + 1,
						xE: xR,
						y:  y,
					}] = nL
				}

				// right
				nR := 0
				rightLine := []rune(d.linesInput1[y])
				if _, err := strconv.Atoi(string(rightLine[x+1])); err == nil {
					xL := x
					xR := lookRight(rightLine, x+1)

					nR, _ = strconv.Atoi(string(rightLine[xL+1 : xR]))

					numbers[pos{
						xS: xL + 1,
						xE: xR,
						y:  y,
					}] = nR
				}

				// above left & above right
				nTL := 0
				nTR := 0

				if nT == 0 {
					if _, err := strconv.Atoi(string(aboveLine[x-1])); err == nil {
						xL := lookLeft(aboveLine, x-1)
						xR := x

						nTL, _ = strconv.Atoi(string(aboveLine[xL+1 : xR]))

						numbers[pos{
							xS: xL + 1,
							xE: xR,
							y:  y - 1,
						}] = nTL
					}

					if _, err := strconv.Atoi(string(aboveLine[x+1])); err == nil {
						xL := x
						xR := lookRight(aboveLine, x+1)

						nTR, _ = strconv.Atoi(string(aboveLine[xL+1 : xR]))

						numbers[pos{
							xS: xL + 1,
							xE: xR,
							y:  y - 1,
						}] = nTR
					}
				}

				// below left & below right
				nBL := 0
				nBR := 0

				if nB == 0 {
					if _, err := strconv.Atoi(string(belowLine[x-1])); err == nil {
						xL := lookLeft(belowLine, x-1)
						xR := x

						nBL, _ = strconv.Atoi(string(belowLine[xL+1 : xR]))

						numbers[pos{
							xS: xL + 1,
							xE: xR,
							y:  y + 1,
						}] = nBL
					}

					if _, err := strconv.Atoi(string(belowLine[x+1])); err == nil {
						xL := x
						xR := lookRight(belowLine, x+1)

						nBR, _ = strconv.Atoi(string(belowLine[xL+1 : xR]))

						numbers[pos{
							xS: xL + 1,
							xE: xR,
							y:  y + 1,
						}] = nBR
					}
				}
			}
		}
	}

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func lookRight(s []rune, x int) int {
	xP := x + 1
	for {
		if xP >= len(s) {
			break
		}

		if _, err := strconv.Atoi(string(s[xP])); err != nil {
			break
		}

		xP++
	}

	return xP
}

func lookLeft(s []rune, x int) int {
	xM := x - 1
	for {
		if xM < 0 {
			break
		}

		if _, err := strconv.Atoi(string(s[xM])); err != nil {
			break
		}

		xM--
	}

	return xM
}

func (d day3) step2() any {
	sum := 0

	toIgnore := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}
	for y, line := range d.linesInput1 {
		splitted := strings.Split(line, "")

		for x, s := range splitted {
			if !slices.Contains(toIgnore, s) {
				partNumber := make([]int, 0)

				// above
				nT := 0
				aboveLine := []rune(d.linesInput1[y-1])
				if _, err := strconv.Atoi(string(aboveLine[x])); err == nil {
					xL := lookLeft(aboveLine, x)
					xR := lookRight(aboveLine, x)

					nT, _ = strconv.Atoi(string(aboveLine[xL+1 : xR]))

					partNumber = append(partNumber, nT)
				}

				// below
				nB := 0
				belowLine := []rune(d.linesInput1[y+1])
				if _, err := strconv.Atoi(string(belowLine[x])); err == nil {
					xL := lookLeft(belowLine, x)
					xR := lookRight(belowLine, x)

					nB, _ = strconv.Atoi(string(belowLine[xL+1 : xR]))

					partNumber = append(partNumber, nB)
				}

				// left
				nL := 0
				leftLine := []rune(d.linesInput1[y])
				if _, err := strconv.Atoi(string(leftLine[x-1])); err == nil {
					xL := lookLeft(leftLine, x-1)
					xR := x

					nL, _ = strconv.Atoi(string(leftLine[xL+1 : xR]))

					partNumber = append(partNumber, nL)
				}

				// right
				nR := 0
				rightLine := []rune(d.linesInput1[y])
				if _, err := strconv.Atoi(string(rightLine[x+1])); err == nil {
					xL := x
					xR := lookRight(rightLine, x+1)

					nR, _ = strconv.Atoi(string(rightLine[xL+1 : xR]))

					partNumber = append(partNumber, nR)
				}

				// above left & above right
				nTL := 0
				nTR := 0

				if nT == 0 {
					if _, err := strconv.Atoi(string(aboveLine[x-1])); err == nil {
						xL := lookLeft(aboveLine, x-1)
						xR := x

						nTL, _ = strconv.Atoi(string(aboveLine[xL+1 : xR]))

						partNumber = append(partNumber, nTL)
					}

					if _, err := strconv.Atoi(string(aboveLine[x+1])); err == nil {
						xL := x
						xR := lookRight(aboveLine, x+1)

						nTR, _ = strconv.Atoi(string(aboveLine[xL+1 : xR]))

						partNumber = append(partNumber, nTR)
					}
				}

				// below left & below right
				nBL := 0
				nBR := 0

				if nB == 0 {
					if _, err := strconv.Atoi(string(belowLine[x-1])); err == nil {
						xL := lookLeft(belowLine, x-1)
						xR := x

						nBL, _ = strconv.Atoi(string(belowLine[xL+1 : xR]))

						partNumber = append(partNumber, nBL)
					}

					if _, err := strconv.Atoi(string(belowLine[x+1])); err == nil {
						xL := x
						xR := lookRight(belowLine, x+1)

						nBR, _ = strconv.Atoi(string(belowLine[xL+1 : xR]))

						partNumber = append(partNumber, nBR)
					}
				}

				if len(partNumber) == 2 {
					sum += partNumber[0] * partNumber[1]
				}
			}
		}
	}

	return sum
}
