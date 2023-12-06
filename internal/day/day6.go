package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	DayCollection["6"] = Day6
}

type day6 struct {
	linesInput1 []string
}

func Day6(input1, input2 string) (any, any) {
	d := day6{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day6) step1() any {
	type race struct {
		time   int
		record int
	}

	reg := regexp.MustCompile(`\d+`)
	times := reg.FindAllString(d.linesInput1[0], -1)
	records := reg.FindAllString(d.linesInput1[1], -1)

	races := make([]race, 0)
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(records[i])
		races = append(races, race{
			time:   t,
			record: d,
		})
	}

	result := 1
	for _, race := range races {
		s := 0
		// from start
		for {
			d := (s * (race.time - s))
			if d > race.record {
				break
			}

			s++
		}

		e := race.time
		// from end
		for {
			d := (e * (race.time - e))
			if d > race.record {
				break
			}

			e--
		}

		diff := e - s + 1
		result *= diff
	}

	return result
}

func (d day6) step2() any {
	timesS := strings.ReplaceAll(d.linesInput1[0], " ", "")
	recordsS := strings.ReplaceAll(d.linesInput1[1], " ", "")

	var time, record int
	fmt.Sscanf(timesS, "Time:%d", &time)
	fmt.Sscanf(recordsS, "Distance:%d", &record)

	s := 0
	// from start
	for {
		d := (s * (time - s))
		if d > record {
			break
		}

		s++
	}

	e := time
	// from end
	for {
		d := (e * (time - e))
		if d > record {
			break
		}

		e--
	}

	return e - s + 1
}
