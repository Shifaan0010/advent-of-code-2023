package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ColorCounts struct {
	red   int
	blue  int
	green int
}

func (counts ColorCounts) LessOrEq(other ColorCounts) bool {
	return counts.red <= other.red && counts.blue <= other.blue && counts.green <= other.green
}

type Game struct {
	id     int
	counts []ColorCounts
}

var gameRegex = regexp.MustCompile(`Game (\d+)`)
var redRegex = regexp.MustCompile(`(\d+) red`)
var blueRegex = regexp.MustCompile(`(\d+) blue`)
var greenRegex = regexp.MustCompile(`(\d+) green`)

func sumPossibleGameIds(games []Game, maxCounts ColorCounts) int {
	idSum := 0

	for _, game := range games {
		possible := true

		for _, colorCounts := range game.counts {
			if !colorCounts.LessOrEq(maxCounts) {
				possible = false
				break
			}
		}

		if possible {
			idSum += game.id
		}
	}

	return idSum
}

func minColorsNeeded(game Game) ColorCounts {
	minColorCount := ColorCounts{}

	for _, colorCount := range game.counts {
		if colorCount.red > minColorCount.red {
			minColorCount.red = colorCount.red
		}

		if colorCount.green > minColorCount.green {
			minColorCount.green = colorCount.green
		}

		if colorCount.blue > minColorCount.blue {
			minColorCount.blue = colorCount.blue
		}
	}

	return minColorCount
}

func powerSum(games []Game) int {
	sum := 0

	for _, game := range games {
		minCount := minColorsNeeded(game)

		sum += minCount.red * minCount.green * minCount.blue;
	}

	return sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	games := []Game{}

	for sc.Scan() {
		line := sc.Text()

		s := strings.Split(line, ":")

		matches := gameRegex.FindStringSubmatch(s[0])
		gameId, _ := strconv.Atoi(matches[1])

		// fmt.Println(gameId)

		sets := strings.Split(s[1], ";")
		counts := []ColorCounts{}

		for _, set := range sets {
			colorCounts := ColorCounts{}

			if matches := redRegex.FindStringSubmatch(set); len(matches) > 0 {
				colorCounts.red, _ = strconv.Atoi(matches[1])
			}

			if matches := blueRegex.FindStringSubmatch(set); len(matches) > 0 {
				colorCounts.blue, _ = strconv.Atoi(matches[1])
			}

			if matches := greenRegex.FindStringSubmatch(set); len(matches) > 0 {
				colorCounts.green, _ = strconv.Atoi(matches[1])
			}

			counts = append(counts, colorCounts)
		}

		games = append(games, Game{id: gameId, counts: counts})
	}

	// fmt.Printf("%#v\n", games)

	fmt.Printf("Part 1\n")

	maxCounts := ColorCounts{red: 12, green: 13, blue: 14}
	fmt.Printf("Sum: %d\n", sumPossibleGameIds(games, maxCounts))

	fmt.Printf("\nPart 2\n")
	fmt.Printf("PowerSum: %d\n", powerSum(games))
}
