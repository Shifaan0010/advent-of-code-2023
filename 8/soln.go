package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"slices"
	"sort"
)

type Dir int

const Left Dir = 0
const Right Dir = 1

type CamelMap struct {
	graph map[string][2]string
	path  []Dir
}

type CamelMapState struct {
	node    string
	pathIdx int
}

type CamelMapGhostState struct {
	nodes   []string
	pathIdx int
}

func (camelmap CamelMap) nextNode(node string, dir Dir) (string, error) {
	currNodePaths, ok := camelmap.graph[node]
	if !ok {
		return "", errors.New(fmt.Sprintf("Node \"%s\" does not exist on graph", node))
	}

	return currNodePaths[dir], nil
}

func (camelmap CamelMap) nextState(state CamelMapState) (CamelMapState, error) {
	direction := camelmap.path[state.pathIdx]

	nextNode, err := camelmap.nextNode(state.node, direction)
	if err != nil {
		return CamelMapState{}, err
	}

	return CamelMapState{
		node:    nextNode,
		pathIdx: (state.pathIdx + 1) % len(camelmap.path),
	}, nil
}

func (camelmap CamelMap) nextStateGhost(state CamelMapGhostState) (CamelMapGhostState, error) {
	direction := camelmap.path[state.pathIdx]

	for i := range state.nodes {
		nextNode, err := camelmap.nextNode(state.nodes[i], direction)
		if err != nil {
			return CamelMapGhostState{}, err
		}

		state.nodes[i] = nextNode
	}

	state.pathIdx = (state.pathIdx + 1) % len(camelmap.path)

	return state, nil
}

// returns (steps int, reachable bool)
func minStepsToTarget(camelmap CamelMap, start string, end string) (int, bool, error) {
	// fmt.Println("minStepsToTarget", camelmap.graph, camelmap.path, start, end)

	steps := 0

	startState := CamelMapState{
		node:    start,
		pathIdx: 0,
	}

	state := startState

	// fmt.Printf("Node\tPathIdx\tSteps\n")
	// fmt.Printf("%s\t%d\t%d\n", state.node, state.pathIdx, steps)

	for state.node != end {
		// cannot use "state, err := ..." since that would create a new local variable
		nextState, err := camelmap.nextState(state)
		if err != nil {
			return 0, true, err
		}

		state = nextState

		steps += 1

		// fmt.Printf("%s\t%d\t%d\n", state.node, state.pathIdx, steps)

		if state == startState { // cycle
			break
		}
	}

	return steps, state.node == end, nil
}

func equal(arr1 []string, arr2 []string) bool {
	sort.Slice(arr1, func(i int, j int) bool {
		return arr1[i] < arr1[j]
	})

	sort.Slice(arr2, func(i int, j int) bool {
		return arr2[i] < arr2[j]
	})

	return slices.Equal(arr1, arr2)
}

func minStepsToGhostTargetsBruteForce(camelmap CamelMap, startNodes []string, endNodes []string) (int, error) {
	fmt.Println("minStepsToGhostTargetsBruteForce", camelmap.graph, camelmap.path, startNodes, endNodes)

	state := CamelMapGhostState{
		nodes:   startNodes,
		pathIdx: 0,
	}

	steps := 0

	fmt.Println(state)

	for !equal(state.nodes, endNodes) {
		// cannot use "state, err := ..." since that would create a new local variable
		nextState, err := camelmap.nextStateGhost(state)
		if err != nil {
			return 0, err
		}

		state = nextState

		steps += 1

		fmt.Println(state)
	}

	return steps, nil
}

type NodeCycleDistance struct {
	startNode           string
	distancesToEndNodes map[string][]int
	cycleInfo           CycleInfo
}

func minStepsToGhostTargets_Modulo(nodeDistances []NodeCycleDistance, startNodes, endNodes []string) (*big.Int, error) {
	// TODO: check if all start and nodes are reachable and inside cycles

	maxStepsToStart := 0
	for _, nodeCycle := range nodeDistances {
		if nodeCycle.cycleInfo.stepsToCycle > maxStepsToStart {
			maxStepsToStart = nodeCycle.cycleInfo.stepsToCycle
		}
	}

	modularEqns := []ModularEqnBig{}
	for _, nodeCycle := range nodeDistances {
		if len(nodeCycle.distancesToEndNodes) != 1 {
			return big.NewInt(-1), errors.New(fmt.Sprintf("Invalid Cycle [%v]", nodeCycle))
		}

		distanceToEndNode := 0
		for _, dist := range nodeCycle.distancesToEndNodes {
			distanceToEndNode = dist[0]
		}

		if distanceToEndNode < nodeCycle.cycleInfo.stepsToCycle {
			return big.NewInt(-1), errors.New(fmt.Sprintf("End Node not in Cycle [%v]", nodeCycle))
		}

		modularEqns = append(modularEqns, ModularEqnBig{
			val: big.NewInt(int64((nodeCycle.cycleInfo.cycleLen + (distanceToEndNode-maxStepsToStart)%nodeCycle.cycleInfo.cycleLen) % nodeCycle.cycleInfo.cycleLen)),
			mod: big.NewInt(int64(nodeCycle.cycleInfo.cycleLen)),
		})
	}

	soln, err := solveModularEqnsBig(modularEqns)

	if err != nil {
		return nil, err
	}

	return soln.val.Add(soln.val, big.NewInt(int64(maxStepsToStart))), nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	camelmap := CamelMap{}

	// Read path
	sc.Scan()
	pathStr := sc.Text()

	camelmap.path = make([]Dir, len(pathStr))
	for i, ch := range pathStr {
		if ch == 'L' {
			camelmap.path[i] = Left
		} else {
			camelmap.path[i] = Right
		}
	}

	// fmt.Printf("Path: %#v\n", path)

	// Read graph
	camelmap.graph = map[string][2]string{}

	r := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	sc.Scan()
	for sc.Scan() {
		line := sc.Text()

		matches := r.FindStringSubmatch(line)

		root := matches[1]

		camelmap.graph[root] = [2]string{Left: matches[2], Right: matches[3]}
	}

	fmt.Println("Part 1")

	steps, reachable, err := minStepsToTarget(camelmap, "AAA", "ZZZ")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if !reachable {
			fmt.Println("Impossible")
		} else {
			fmt.Printf("Steps: %d\n", steps)
		}
	}

	fmt.Println()

	fmt.Println("Part 2")

	startNodes := []string{}
	endNodes := []string{}

	for node := range camelmap.graph {
		if node[len(node)-1] == 'A' {
			startNodes = append(startNodes, node)
		} else if node[len(node)-1] == 'Z' {
			endNodes = append(endNodes, node)
		}
	}

	fmt.Println(startNodes, endNodes)

	// brute force
	// steps2, err := minStepsToGhostTargetsBruteForce(camelmap, append(startNodes[:0:0], startNodes...), append(endNodes[:0:0], endNodes...))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Steps: %d\n", steps2)
	// }

	nodeDistances := []NodeCycleDistance{}

	// walk the cycle to find all end nodes
	fmt.Println("\nNode\tCycleLen\tStepsToCycle")
	for _, node := range startNodes {
		cycleInfo, err := tortoiseAndHare(camelmap, node)

		if err != nil {
			fmt.Printf("%s\tcould not find cycle\n", node)
			continue
		}

		distancesToEndNodes := map[string][]int{}

		state := CamelMapState{
			node:    node,
			pathIdx: 0,
		}
		for i := 0; i < cycleInfo.stepsToCycle+cycleInfo.cycleLen; i += 1 {
			if state.node[len(state.node)-1] == 'Z' {
				if _, exists := distancesToEndNodes[state.node]; !exists {
					distancesToEndNodes[state.node] = []int{}
				}

				distancesToEndNodes[state.node] = append(distancesToEndNodes[state.node], i)
			}

			state, err = camelmap.nextState(state)
		}

		fmt.Printf("%s\t%d\t\t%d\t%v\n", node, cycleInfo.cycleLen, cycleInfo.stepsToCycle, distancesToEndNodes)

		nodeDistances = append(nodeDistances, NodeCycleDistance{
			startNode:           node,
			distancesToEndNodes: distancesToEndNodes,
			cycleInfo:           cycleInfo,
		})
	}

	minStepsToEndNodes, err := minStepsToGhostTargets_Modulo(nodeDistances, startNodes, endNodes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Steps: %d\n", minStepsToEndNodes)
	}
}
