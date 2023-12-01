package main

import (
	"fmt"
	"hyoa/aoc2023/internal/day"
	"os"
)

func main() {
	days := day.DayCollection

	n := os.Args[2]

	exec, ok := days[n]

	if !ok {
		panic("Day not found")
	}

	kind := os.Args[1]

	input1 := fmt.Sprintf("input/day%s/%s.txt", n, kind)
	input2 := fmt.Sprintf("input/day%s/%s.txt", n, kind)
	if kind == "example" {
		input1 = fmt.Sprintf("input/day%s/%s_1.txt", n, kind)
		input2 = fmt.Sprintf("input/day%s/%s_2.txt", n, kind)
	}

	output1, output2 := exec(input1, input2)

	fmt.Printf("Step 1: %v\n", output1)
	fmt.Printf("Step 2: %v\n", output2)
}
