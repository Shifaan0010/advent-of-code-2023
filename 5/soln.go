package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start  int
	length int
}

type RangeMap struct {
	source int
	dest   int
	length int
}

func (rng RangeMap) contains(n int) bool {
	return rng.source <= n && n < rng.source + rng.length
}

func (rng RangeMap) mapVal(val int) int {
	return val - rng.source + rng.dest
}

type Mapping struct {
	Ranges []RangeMap
}

func NewMapping(ranges []RangeMap) Mapping {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].source < ranges[j].source
	})

	return Mapping{Ranges: ranges}
}

func (mapping Mapping) mapVal(val int) int {
	for _, rng := range mapping.Ranges {
		if rng.contains(val) {
			return rng.mapVal(val)
		}
	}

	return val
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var seeds []int
	mappings := []Mapping{}

	var currRanges []RangeMap

	for sc.Scan() {
		line := sc.Text()

		if strings.HasPrefix(line, "seeds:") {
			split := strings.Split(line, " ")

			seeds = make([]int, 0, len(split)-1)
			for _, s := range split[1:] {
				seed, _ := strconv.Atoi(s)
				seeds = append(seeds, seed)
			}
		} else if strings.HasPrefix(line, "seed-to-soil map:") ||
			strings.HasPrefix(line, "seed-to-soil map:") ||
			strings.HasPrefix(line, "soil-to-fertilizer map:") ||
			strings.HasPrefix(line, "fertilizer-to-water map:") ||
			strings.HasPrefix(line, "water-to-light map:") ||
			strings.HasPrefix(line, "light-to-temperature map:") ||
			strings.HasPrefix(line, "temperature-to-humidity map:") ||
			strings.HasPrefix(line, "humidity-to-location map:") {

			if currRanges != nil {
				mappings = append(mappings, NewMapping(currRanges))
			}
			currRanges = []RangeMap{}
		} else {
			split := strings.Split(line, " ")

			if len(split) != 3 {
				continue
			}

			dest, _ := strconv.Atoi(split[0])
			source, _ := strconv.Atoi(split[1])
			length, _ := strconv.Atoi(split[2])

			currRanges = append(currRanges, RangeMap{dest: dest, source: source, length: length})
		}
	}

	if len(currRanges) > 0 {
		mappings = append(mappings, NewMapping(currRanges))
	}

	// fmt.Println(mappings)

	fmt.Println("Part 1")

	minLocation := math.MaxInt

	mappedSeeds := make([]int, 0, len(seeds))
	for _, val := range seeds {
		for _, mapping := range mappings {
			val = mapping.mapVal(val)
		}

		minLocation = min(minLocation, val)

		mappedSeeds = append(mappedSeeds, val)
	}

	fmt.Println(mappedSeeds)

	fmt.Printf("Min Location: %d\n", minLocation)
}
