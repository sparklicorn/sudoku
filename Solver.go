package sudoku

func Solve(board *Board) bool {

	//track all empty cells in minheap ordered by number of candidates

	//while heap is not empty
	//	cell := remove heap top
	//	if cell has zero candidates and is empty
	//		TODO backtrack
	//		continue
	//	if cell has only 1 candidate
	//		add cell to backtrack
	//		continue
	//
	//

	return false
}

func FindAllSolutions(board *Board) []Board {
	solutions := make(map[string]bool)
	queue := make([]*Board, 0, 16)
	q2 := make([]*Board, 0, 16)

	copy(q2, queue)

	for len(queue) > 0 {
		b, queue := *queue[0], queue[1:] //dereference Board from queue
		index := findCellWithLeastCandidates(&b)
		candidates := b.GetCandidates(index)

		for candidate := range candidates {
			next := b.Copy()
			next.SetCell(index, candidate)

			if next.IsFull() {
				solutions[b.SimpleString()] = true
			} else {
				if reduceBoard(&next) {
					queue = append(queue, &next) //pointer to next
				}
			}
		}
	}

	results := make([]Board, 0, len(solutions))
	i := 0
	for str := range solutions {
		results[i] = LoadBoard(str)
		i++
	}

	return results
}

func HasSingleSolution(config *Board) (bool, Board) {
	solutions := FindAllSolutions(config)
	if len(solutions) == 1 {
		return true, solutions[0]
	}
	return false, Board{}
}

//Maps raw board value to associated number of candidates
var numCandidates = make(map[int]int)

func init() {
	for i := 0; i <= ALL; i++ {
		count := 0
		for j := 0; j < 9; j++ {
			if (i >> j & 1) == 1 {
				count++
			}
		}
		numCandidates[i] = count
	}
}

func findCellWithLeastCandidates(b *Board) int {
	minCell := -1
	currentMin := 10

	for i := 0; i < NUM_CELLS; i++ {
		count := numCandidates[b.Cell(i)]
		if count > 1 && count < currentMin {
			currentMin = count
			minCell = i
			if currentMin == 2 {
				break
			}
		}
	}

	return minCell
}

//returns false if board is found to be invalid
func reduceBoard(board *Board) bool {
	dirty := []int
	changed := true
	overallChange := false

	for i := 0; i < NUM_CELLS; i++ {
		if Decode(board.Cell(i)) == 0 {
			dirty = append(dirty, i)
		}
	}

	for changed {
		changed = false

		for i := len(dirty) - 1; i >= 0; i-- {
			cell := dirty[len(dirty) - 1]

			if reduce(board, cell) {
				changed = true
				overallChange = true
				if board.decodeCell(cell) == 0 {
					dirty = append(dirty[:cell], dirty[cell+1:]...)
				}
			}
		}

	}

	return false
}

//if cell is empty and has no candidates, return false
//if cell is decoded into a digit, return false

func reduce(board *Board, cellIndex int) bool {

	candidates := board.Cell(cellIndex)
	if candidates == 0 || Decode(candidates) > 0 {
		return false
	}

	forEachInRow(board, getRowIndex(cellIndex), func(index int) {
		if Decode(board.Cell(index)) > 0 {

		}
	})

	return false
}

func forEachInRow(board *Board, row int, runnable func(index int)) {
	for cellIndex := row * 9; cellIndex < (row+1)*9; cellIndex++ {
		runnable(cellIndex)
	}
}

func forEachInColumn(board *Board, col int, runnable func(index int)) {
	for cellIndex := col; cellIndex < NUM_CELLS; cellIndex += 9 {
		runnable(cellIndex)
	}
}

func forEachInRegion(board *Board, region int, runnable func(index int)) {
	anchor := getFirstIndexInRegion(region)
	for n := 0; n < 9; n++ {
		runnable(anchor + (n/3)*9 + (n % 3))
	}
}

func reduceRow(board *Board, cellIndex, candidates int) int {
	if board.decodeCell(cellIndex) > 0 {
		return board.Cell(cellIndex)
	}

	rowIndex := getRowIndex(cellIndex)

	forEachInRow(board, rowIndex, func(i int) {
		if i != cellIndex {
			value := board.Cell(i)
			if board.decodeCell(i) > 0 {
				// if (candidates ^ value) < candidates {
				candidates ^= value
				// }
			}
		}
	})

	// for i := rowIndex * 9; i < (rowIndex+1)*9; i++ {
	// 	if i == cellIndex {
	// 		continue
	// 	}

	// 	value := board.Cell(i)
	// 	if board.decodeCell(i) > 0 {
	// 		if (candidates ^ value) < candidates {
	// 			candidates ^= value
	// 		}
	// 	}
	// }

	return candidates
}

func reduceColumn(board *Board, cellIndex, candidates int) int {
	if board.decodeCell(cellIndex) > 0 {
		return board.Cell(cellIndex)
	}
	col := getColumnIndex(cellIndex)

	for i := col; i < NUM_CELLS; i += 9 {
		if i == cellIndex {
			continue
		}

		value := board.Cell(i)
		if board.decodeCell(i) > 0 {
			if (candidates ^ value) < candidates {
				candidates ^= value
			}
		}
	}

	return candidates
}

func reduceRegion(board *Board, cellIndex, candidates int) int {
	if board.decodeCell(cellIndex) > 0 {
		return board.Cell(cellIndex)
	}
	region := getRegionIndex(cellIndex)
	anchor := getFirstIndexInRegion(region)

	for n := 0; n < 9; n++ {
		i := anchor + (n/3)*9 + (n % 3)
		if i == cellIndex {
			continue
		}

		value := board.Cell(i)
		if board.decodeCell(i) > 0 {
			if (candidates ^ value) < candidates {
				candidates ^= value
			}
		}
	}

	return candidates
}

func getRowIndex(cellIndex int) int {
	return cellIndex / 9
}

func getColumnIndex(cellIndex int) int {
	return cellIndex % 9
}

func getRegionIndex(cellIndex int) int {
	return (cellIndex/9)/3 + (cellIndex%9)/3
}

func getFirstIndexInRegion(region int) int {
	return (region/3)*27 + (region%3)*3
}
