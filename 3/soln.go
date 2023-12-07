package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type SymPos struct {
	sym rune
	y   int
	x   int
}

type NumPos struct {
	num   int
	y     int
	left  int
	right int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (num NumPos) IsAdjacent(pos SymPos) bool {
	return (pos.y == num.y && ((pos.x == num.left-1) || (pos.x == num.right))) ||
		(Abs(pos.y-num.y) == 1 && num.left-1 <= pos.x && pos.x <= num.right)
}

func 

func main() {
	sc := bufio.NewScanner(os.Stdin)

	lines := []string{}

	for sc.Scan() {
		line := sc.Text()

		lines = append(lines, line)
	}

	numRegex := regexp.MustCompile(`(\d+)`)

	numPositions := []NumPos{}
	symbolPositions := []SymPos{}

	for y, line := range lines {
		matchPos := numRegex.FindAllStringIndex(line, -1)

		// fmt.Println(matchPos)

		for _, pos := range matchPos {
			num, _ := strconv.Atoi(line[pos[0]:pos[1]])
			numPositions = append(numPositions, NumPos{num: num, y: y, left: pos[0], right: pos[1]})
		}
	}

	for y := range lines {
		for x, ch := range lines[y] {
			if !unicode.IsDigit(ch) && ch != '.' {
				symbolPositions = append(symbolPositions, SymPos{sym: ch, y: y, x: x})
			}
		}
	}

	fmt.Println(numPositions)
	fmt.Println(symbolPositions)
}
