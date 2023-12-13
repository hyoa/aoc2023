package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func init() {
	DayCollection["12"] = Day12
}

type day12 struct {
	linesInput1 []string
}

func Day12(input1, input2 string) (any, any) {
	d := day12{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

type searchNode struct {
	value string
	left  *searchNode
	right *searchNode
}

func navigateTree(n *searchNode, found *int, rule []int, regMatchPattern *regexp.Regexp, regPartialMatch *regexp.Regexp) {
	if n == nil {
		return
	}

	if n.left == nil && n.right == nil {
		r1, r2, canContinue := replaceMark(n.value)

		if !canContinue {
			if matchPattern(n.value, rule, regMatchPattern) {
				*found++
			}
		}

		if strings.Contains(r1, "?") {
			if ok, v1 := partialMatch(r1, rule, regPartialMatch); ok {
				n.left = &searchNode{
					value: v1,
				}
			}
		} else {
			if matchPattern(r1, rule, regMatchPattern) {
				*found++
			}
		}

		if strings.Contains(r2, "?") {
			if ok, v2 := partialMatch(r2, rule, regPartialMatch); ok {
				n.right = &searchNode{
					value: v2,
				}
			}
		} else {
			if matchPattern(r2, rule, regMatchPattern) {
				*found++
			}
		}

	}

	navigateTree(n.left, found, rule, regMatchPattern, regPartialMatch)
	navigateTree(n.right, found, rule, regMatchPattern, regPartialMatch)

	n.left = nil
	n.right = nil
}

func (d day12) step1() any {
	var result int

	type record struct {
		pattern string
		rule    []int
	}

	regexMatchPattern := `(#+)`
	regMatchPattern := regexp.MustCompile(regexMatchPattern)

	regexPartialMatch := `(#+\?*)`
	regPartialMatch := regexp.MustCompile(regexPartialMatch)

	records := make([]record, 0)
	for _, line := range d.linesInput1 {
		splitted := strings.Split(line, " ")
		pattern := splitted[0]
		ruleString := splitted[1]

		rule := make([]int, 0)
		for _, s := range strings.Split(ruleString, ",") {
			r, _ := strconv.Atoi(s)
			rule = append(rule, r)
		}

		records = append(records, record{
			pattern: pattern,
			rule:    rule,
		})
	}

	for _, r := range records {
		n := &searchNode{
			value: r.pattern,
		}

		found := 0

		navigateTree(n, &found, r.rule, regMatchPattern, regPartialMatch)

		result += found
	}

	return result
}

func (d day12) step2() any {
	var result int

	type record struct {
		pattern string
		rule    []int
	}

	records := make([]record, 0)
	for _, line := range d.linesInput1 {
		splitted := strings.Split(line, " ")
		pattern := splitted[0]
		ruleString := splitted[1]

		for i := 0; i < 4; i++ {
			pattern += "?" + splitted[0]
			ruleString += "," + splitted[1]
		}

		rule := make([]int, 0)
		for _, s := range strings.Split(ruleString, ",") {
			r, _ := strconv.Atoi(s)
			rule = append(rule, r)
		}

		records = append(records, record{
			pattern: pattern,
			rule:    rule,
		})
	}

	regexPartialMatch := `(#+\?*)`
	regPartialMatch := regexp.MustCompile(regexPartialMatch)

	regexMatchPattern := `(#+)`
	regMatchPattern := regexp.MustCompile(regexMatchPattern)

	// wg := errgroup.Group{}

	// type lockResult struct {
	// 	mu     sync.Mutex
	// 	result int
	// 	nbDone int
	// }

	// lr := lockResult{
	// 	result: 0,
	// }

	// var m runtime.MemStats

	// wg.SetLimit(10)

	for _, r := range records {
		n := &searchNode{
			value: r.pattern,
		}

		found := 0

		navigateTree(n, &found, r.rule, regMatchPattern, regPartialMatch)

		result += found
	}

	PrintMemUsage()

	return result
}

func matchPattern(pattern string, rule []int, reg *regexp.Regexp) bool {
	match := reg.FindAllString(pattern, -1)

	if len(match) != len(rule) {
		return false
	}

	for k := range rule {
		if len(match[k]) != rule[k] {
			return false
		}
	}

	return true
}

func partialMatch(pattern string, rule []int, reg *regexp.Regexp) (bool, string) {
	ruleAgg := 0

	for _, r := range rule {
		ruleAgg += r
	}

	countHash := 0
	for _, c := range pattern {
		if c == '#' {
			countHash++
		}
	}

	if countHash > ruleAgg {
		return false, pattern
	}

	match := reg.FindAllString(pattern, -1)
	idx := reg.FindAllStringIndex(pattern, -1)
	mark := strings.Index(pattern, "?")

	counterHash := 0
	isHash := true

	if len(match) == 0 {
		return false, pattern
	}

	for _, c := range match[0] {
		if c == '#' {
			counterHash++
			isHash = true
		} else {
			isHash = false
		}

		if !isHash {
			break
		}
	}

	if counterHash > rule[0] {
		sIdx := strings.Index(pattern, match[0])

		if sIdx != -1 && mark > sIdx {
			return false, pattern

		}
	}

	for m := range match {
		if m > len(rule)-1 {
			return true, pattern
		}

		if strings.Contains(match[m], "?") {
			str := strings.Repeat("#", rule[m]) + "?"
			if match[m] == str {
				pattern = strings.Replace(pattern, "#?", "#.", 1)
				return partialMatch(pattern, rule, reg)
			} else {
				counterHash := 0
				counterMark := 0
				for _, c := range match[m] {
					if c == '#' {
						counterHash++
					}

					if c == '?' {
						counterMark++
					}
				}

				if counterHash > rule[m] && counterMark < rule[m] {
					return false, pattern
				}

				idxS := strings.Index(pattern, match[m])
				idxE := idxS + len(match[m]) + 1

				if idxS != -1 && idxE <= len(pattern)-1 {
					if pattern[idxE] == '.' {
						return false, pattern
					}
				}
			}

			return true, pattern
		}

		if len(match[m]) != rule[m] {
			t := mark < idx[m][1]

			idxStart := idx[m][0] - 1
			idxEnd := idx[m][1]

			if t && mark != -1 && idxStart >= 0 && idxEnd <= len(pattern)-1 {
				char := pattern[idx[m][0]-1 : idxEnd]
				if len(char) == rule[m] {
					pattern = replaceAtIndex(pattern, '#', idx[m][0]-1)

					return partialMatch(pattern, rule, reg)
				}
			}
			return t, pattern
		}
	}

	return true, pattern
}

func replaceMark(s string) (string, string, bool) {
	idx := strings.Index(s, "?")

	if idx == -1 {
		return "", "", false
	}

	r1 := strings.Replace(s, "?", ".", 1)
	r2 := strings.Replace(s, "?", "#", 1)

	return r1, r2, true
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
