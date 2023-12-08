package day

import (
	"fmt"
	"hyoa/aoc2023/internal/utils"
	"sort"
	"strings"
)

func init() {
	DayCollection["7"] = Day7
}

type day7 struct {
	linesInput1 []string
}

var scoreCardBalance1 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var scoreCardBalance2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 0,
}

var scoreKind = map[string]int{
	"high card":       1,
	"one pair":        10,
	"two pair":        20,
	"three of a kind": 30,
	"full house":      40,
	"four of a kind":  50,
	"five of a kind":  100,
}

func Day7(input1, input2 string) (any, any) {

	d := day7{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

type hand struct {
	cards string
	bid   int
	kind  string
}

type cardCount map[string]int

func (d day7) step1() any {
	var result int

	hands := make([]hand, 0)
	for _, line := range d.linesInput1 {
		h, _ := getHand(line)

		hands = append(hands, h)
	}

	hands = sortHands(hands, scoreCardBalance1)

	i := len(hands)

	for _, h := range hands {
		result += h.bid * i
		i--
	}

	return result
}

func (d day7) step2() any {
	var result int

	hands := make([]hand, 0)
	for _, line := range d.linesInput1 {
		h, cC := getHand(line)

		if strings.Contains(h.cards, "J") {
			switch h.kind {
			case "high card":
				h.kind = "one pair"
			case "one pair":
				h.kind = "three of a kind"
			case "two pair":
				if cC["J"] == 2 {
					h.kind = "four of a kind"
				} else {
					h.kind = "full house"
				}
			case "three of a kind":
				h.kind = "four of a kind"
			case "four of a kind":
				h.kind = "five of a kind"
			case "full house":
				h.kind = "five of a kind"
			}
		}

		hands = append(hands, h)
	}

	hands = sortHands(hands, scoreCardBalance2)

	i := len(hands)

	for _, h := range hands {
		result += h.bid * i
		i--
	}

	return result
}

func checkFirstCard(c1, c2 string, idx int, score map[string]int) bool {
	if score[c1[idx:idx+1]] == score[c2[idx:idx+1]] {
		return checkFirstCard(c1, c2, idx+1, score)
	}

	if score[c1[idx:idx+1]] > score[c2[idx:idx+1]] {
		return true
	}

	return false
}

func sortHands(hands []hand, score map[string]int) []hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].kind == hands[j].kind {
			return checkFirstCard(hands[i].cards, hands[j].cards, 0, score)
		}

		return scoreKind[hands[i].kind] > scoreKind[hands[j].kind]
	})

	return hands
}

func getHand(input string) (hand, cardCount) {
	var cards string
	var bid int

	fmt.Sscanf(input, "%s %d", &cards, &bid)

	cC := make(cardCount)
	for _, card := range cards {
		if _, ok := cC[string(card)]; !ok {
			cC[string(card)] = 1
		} else {
			cC[string(card)]++
		}
	}

	h := hand{
		cards: cards,
		bid:   bid,
	}

	switch len(cC) {
	case 5:
		h.kind = "high card"
	case 4:
		h.kind = "one pair"
	case 3:
		// check if two pair or three of a kind
		h.kind = "two pair"

		for _, c := range cC {
			if c == 3 {
				h.kind = "three of a kind"
				break
			}
		}
	case 2:
		// check if full house or four of a kind
		h.kind = "full house"
		for _, c := range cC {
			if c == 4 {
				h.kind = "four of a kind"
				break
			}
		}
	case 1:
		h.kind = "five of a kind"
	}

	return h, cC
}
