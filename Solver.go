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
		}
	}

	return minCell
}

//returns false if board is found to be invalid
func reduceBoard(board *Board) bool {

	return false
}

func reduce(board *Board, cellIndex int) {
	//TODO
}

func reduceRow(board *Board, cellIndex, candidates int) int {
	rowIndex := getRowIndex(cellIndex)

	for i := rowIndex * 9; i < (rowIndex+1)*9; i++ {
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

func reduceColumn(board *Board, cellIndex, candidates int) int {
	col := getColumnIndex(cellIndex)

	for i := col; i < NUM_CELLS; i += 9 {
		if i == cellIndex {
			continue
		}

		value := board.Cell(cellIndex)
		if board.decodeCell(cellIndex) > 0 {
			if (candidates ^ value) < candidates {
				candidates ^= value
			}
		}
	}

	return candidates
}

func reduceRegion(board *Board, cellIndex, candidates int) int {
	col := getRegionIndex(cellIndex)

	for i := col; i < NUM_CELLS; i += 9 {
		if i == cellIndex {
			continue
		}

		value := board.Cell(cellIndex)
		if board.decodeCell(cellIndex) > 0 {
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
