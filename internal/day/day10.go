package day

import (
	"hyoa/aoc2023/internal/utils"
	"slices"
)

func init() {
	DayCollection["10"] = Day10
}

type day10 struct {
	linesInput1 []string
	linesInput2 []string
}

type node struct {
	x, y   int
	v      string
	wallTo map[nodeCoordinate]bool
}

type nodeCoordinate struct {
	x, y int
}

func Day10(input1, input2 string) (any, any) {
	d := day10{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
		linesInput2: utils.ReadTextFileLinesAsString(input2),
	}

	return d.step1(), d.step2()
}

func (d day10) step1() any {
	var result int
	nodes, startNode := getNodes(d.linesInput1)

	path := getPath(nodes, startNode)

	result = int(len(path) / 2)

	return result
}

func (d day10) step2() any {
	var result int
	nodes, start := getNodes(d.linesInput2)
	path := getPath(nodes, start)

	simplePathMapped := make(map[nodeCoordinate]int)
	for k := range path {
		simplePathMapped[nodeCoordinate{path[k].x, path[k].y}] = k
	}

	var top node
	for _, n := range path {
		if slices.Contains([]string{"7", "F"}, n.v) {
			f := moveTop(n, nodes, simplePathMapped)

			if !f {
				if top.x == 0 && top.y == 0 {
					top = n
				} else if top.y > n.y {
					top = n
				}
			}
		}
	}

	idx := simplePathMapped[nodeCoordinate{top.x, top.y}]
	pathL := path[idx-1]

	var dir int
	if pathL.x == top.x+1 || pathL.y == top.y+1 {
		dir = -1
	} else {
		dir = 1
	}

	prev := top

	for {

		idx += dir

		if idx < 0 {
			idx = len(path) - 1
		} else if idx >= len(path) {
			idx = 0
		}

		idxNext := idx + dir
		if idxNext < 0 {
			idxNext = len(path) - 1
		} else if idxNext >= len(path) {
			idxNext = 0
		}

		next := path[idx]
		nextNext := path[idxNext]

		if next.x == top.x && next.y == top.y {
			break
		}

		if prev.x == next.x && prev.y == next.y {
			continue
		}

		walled := make([]nodeCoordinate, 0)

		if next.x == prev.x+1 {
			walled = append(walled, nodeCoordinate{next.x, next.y - 1})
		} else if next.x == prev.x-1 {
			walled = append(walled, nodeCoordinate{next.x, next.y + 1})
		} else if next.y == prev.y+1 {
			walled = append(walled, nodeCoordinate{next.x + 1, next.y})
		} else if next.y == prev.y-1 {
			walled = append(walled, nodeCoordinate{next.x - 1, next.y})
		}

		if nextNext.x == next.x+1 {
			walled = append(walled, nodeCoordinate{next.x, next.y - 1})
		} else if nextNext.x == next.x-1 {
			walled = append(walled, nodeCoordinate{next.x, next.y + 1})
		} else if nextNext.y == next.y+1 {
			walled = append(walled, nodeCoordinate{next.x + 1, next.y})
		} else if nextNext.y == next.y-1 {
			walled = append(walled, nodeCoordinate{next.x - 1, next.y})
		}

		walledMap := make(map[nodeCoordinate]bool)
		for _, w := range walled {
			walledMap[w] = true
		}

		path[idx].wallTo = walledMap

		prev = next
	}

	grounds := make(map[nodeCoordinate]bool)
	for _, n := range nodes {
		if _, ok := simplePathMapped[nodeCoordinate{n.x, n.y}]; !ok {
			grounds[nodeCoordinate{n.x, n.y}] = true

		}
	}

	for g := range grounds {
		nodeToCheck := nodeCoordinate{x: g.x, y: g.y}
		h := hitPath(nodes[nodeToCheck], nodes, simplePathMapped, path)

		if h {
			nodes[nodeToCheck] = node{
				x: g.x,
				y: g.y,
				v: "I",
			}

			result++
		} else {
			nodes[nodeToCheck] = node{
				x: g.x,
				y: g.y,
				v: ".",
			}
		}
	}

	return result
}

func hitPath(curr node, nodes map[nodeCoordinate]node, pathMapped map[nodeCoordinate]int, path []node) bool {
	next := nodeCoordinate{curr.x, curr.y - 1}

	if _, ok := nodes[next]; !ok {
		return false
	}

	if v, ok := pathMapped[next]; ok {
		if _, ok := path[v].wallTo[nodeCoordinate{curr.x, curr.y}]; ok {
			return false
		}

		return true
	}

	return hitPath(nodes[next], nodes, pathMapped, path)
}

func moveTop(curr node, nodes map[nodeCoordinate]node, path map[nodeCoordinate]int) bool {
	next := nodeCoordinate{curr.x, curr.y - 1}

	if _, ok := nodes[next]; !ok {
		return false
	}

	if _, ok := path[next]; ok {
		return true
	}

	return moveTop(nodes[next], nodes, path)
}

func getNextForStart(s node, nodes map[nodeCoordinate]node) []node {
	nextStartAvailable := make([]node, 0)

	startLeft := nodes[nodeCoordinate{s.x - 1, s.y}]
	if startLeft.v == "-" || startLeft.v == "L" || startLeft.v == "F" {
		nextStartAvailable = append(nextStartAvailable, startLeft)
	}

	startRight := nodes[nodeCoordinate{s.x + 1, s.y}]
	if startRight.v == "-" || startRight.v == "J" || startRight.v == "7" {
		nextStartAvailable = append(nextStartAvailable, startRight)
	}

	startUp := nodes[nodeCoordinate{s.x, s.y - 1}]
	if startUp.v == "|" || startUp.v == "F" || startUp.v == "7" {
		nextStartAvailable = append(nextStartAvailable, startUp)
	}

	startDown := nodes[nodeCoordinate{s.x, s.y + 1}]
	if startDown.v == "|" || startDown.v == "L" || startDown.v == "J" {
		nextStartAvailable = append(nextStartAvailable, startDown)
	}

	return nextStartAvailable
}

func findNextPipe(nodes map[nodeCoordinate]node, current, previous node) (node, node) {
	var next node
	switch current.v {
	case "|":
		if previous.y > current.y {
			next = nodes[nodeCoordinate{current.x, current.y - 1}]
		} else {
			next = nodes[nodeCoordinate{current.x, current.y + 1}]
		}
	case "-":
		if previous.x > current.x {
			next = nodes[nodeCoordinate{current.x - 1, current.y}]
		} else {
			next = nodes[nodeCoordinate{current.x + 1, current.y}]
		}
	case "L":
		if previous.x > current.x {
			next = nodes[nodeCoordinate{current.x, current.y - 1}]
		} else {
			next = nodes[nodeCoordinate{current.x + 1, current.y}]
		}
	case "J":
		if previous.x == current.x && previous.y < current.y {
			next = nodes[nodeCoordinate{current.x - 1, current.y}]
		} else {
			next = nodes[nodeCoordinate{current.x, current.y - 1}]
		}
	case "7":
		if previous.x < current.x {
			next = nodes[nodeCoordinate{current.x, current.y + 1}]
		} else {
			next = nodes[nodeCoordinate{current.x - 1, current.y}]
		}
	case "F":
		if previous.x == current.x && previous.y > current.y {
			next = nodes[nodeCoordinate{current.x + 1, current.y}]
		} else {
			next = nodes[nodeCoordinate{current.x, current.y + 1}]
		}
	}

	return next, current
}

func getPath(nodes map[nodeCoordinate]node, startNode node) []node {
	nextStartAvailable := getNextForStart(startNode, nodes)
	var path []node

main:
	for k := range nextStartAvailable {
		current := nextStartAvailable[k]
		previous := startNode

		path = make([]node, 0)

		path = append(path, startNode)
		path = append(path, current)
		for {
			var next node
			next, previous = findNextPipe(nodes, current, previous)
			path = append(path, next)

			if next.v == "S" {
				break main
			}

			current = next
		}
	}

	return path
}

func getNodes(input []string) (map[nodeCoordinate]node, node) {
	y := 0

	var startNode node
	nodes := make(map[nodeCoordinate]node)
	for _, line := range input {
		for x, v := range line {
			n := node{
				x: x,
				y: y,
				v: string(v),
			}
			nodes[nodeCoordinate{x, y}] = n

			if string(v) == "S" {
				startNode = n
			}
		}
		y++
	}

	return nodes, startNode
}
