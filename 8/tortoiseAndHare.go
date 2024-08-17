package main

type CycleInfo struct {
	cycleLen     int
	stepsToCycle int
}

func tortoiseAndHare(camelmap CamelMap, startNode string) (CycleInfo, error) {
	const hareSteps = 2

	startState := CamelMapState{
		node:    startNode,
		pathIdx: 0,
	}

	steps := 0
	tortoise, hare := startState, startState

	// fmt.Println("\nStarting Tortoise & Hare\n")
	// fmt.Println("Tortoise\tHare")
	// fmt.Printf("%v\t\t%v\n", tortoise, hare)
	for met := false; !met; met = (tortoise == hare) {
		// cannot use "tortoise, err := ..." since that would create a new local variable
		nextTortoise, err := camelmap.nextState(tortoise)
		if err != nil {
			return CycleInfo{}, err
		}

		tortoise = nextTortoise

		for i := 0; i < hareSteps; i += 1 {
			// cannot use "hare, err := ..." since that would create a new local variable
			nextHare, err := camelmap.nextState(hare)
			if err != nil {
				return CycleInfo{}, err
			}

			hare = nextHare
		}

		// fmt.Printf("%v\t\t%v\n", tortoise, hare)

		steps += 1
	}

	// node[i] = node[i + k * cycleLen]
	// node[steps] = node[2 * steps]
	// 2 * steps = steps + k * cycleLen
	// steps = k * cycleLen

	// hare pos = tortoise pos = k * cycleLen
	// stepsToCycle = stepsToCycle + k * cycleLen (mod cycleLen)

	// reset tortoise to start and move both tortoise & hare by 1 to find the cycleLen
	// fmt.Println("\nReset tortoise to start and move both tortoise & hare by 1 to find the cycleLen\n")

	tortoise = startState
	stepsToCycle := 0

	// fmt.Println("Tortoise\tHare")
	// fmt.Printf("%v\t\t%v\n", tortoise, hare)

	for tortoise != hare {
		// cannot use "tortoise, err := ..." since that would create a new local variable
		nextTortoise, err := camelmap.nextState(tortoise)
		if err != nil {
			return CycleInfo{}, err
		}

		// cannot use "hare, err := ..." since that would create a new local variable
		nextHare, err := camelmap.nextState(hare)
		if err != nil {
			return CycleInfo{}, err
		}

		tortoise, hare = nextTortoise, nextHare

		// fmt.Printf("%v\t\t%v\n", tortoise, hare)

		stepsToCycle += 1
	}

	// tortoise stays in place while hare moves
	// fmt.Println("\nTortoise stays in place while hare moves\n")

	cycleLen := 0

	// fmt.Println("Tortoise\tHare")
	// fmt.Printf("%v\t\t%v\n", tortoise, hare)

	for met := false; !met; met = (tortoise == hare) {
		// cannot use "hare, err := ..." since that would create a new local variable
		nextHare, err := camelmap.nextState(hare)
		if err != nil {
			return CycleInfo{}, err
		}

		hare = nextHare

		// fmt.Printf("%v\t\t%v\n", tortoise, hare)

		cycleLen += 1
	}

	return CycleInfo{cycleLen: cycleLen, stepsToCycle: stepsToCycle}, nil
}
