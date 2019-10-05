package sudoku

import (
	"math/rand"
	"time"
)

func shuffle(list *[]int, numSwaps int) {
	rand.Seed(time.Now().UnixNano())
	length := len(*list)
	for n := 0; n < length; n++ {
		i, j := rand.Intn(length), rand.Intn(length)
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	}
}

func fillSections(board *Board, sectionsMask int) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for m := 0; m < 9; m++ {
		if sectionsMask&(1<<(8-m)) > 0 {
			shuffle(&list, 18)
			gr, gc := m/3, m%3
			for i := 0; i < 9; i++ {
				index := gr*27 + gc*3 + (i/3)*9 + (i % 3)
				board.SetCell(index, list[i])
			}
		}
	}
}

func GenerateConfig() *Board {
	board := NewBoard()

	fillSections(&board, 0b101010100)
	success, solution := HasSingleSolution(&board)
	for !success {
		fillSections(&board, 0b101010100)
		success, solution = HasSingleSolution(&board)
	}

	return &solution
}

type Node struct {
	value   *Board
	nexts   []*Node
	prev    *Node
	visited bool
}

func GeneratePuzzle(clues, maxPops int) *Board {

	config := GenerateConfig()

	stack := []string{config.SimpleString()}
	visited := map[string]bool{}

	totalPops, pops := 0, 0

	for len(stack) > 0 && pops < maxPops {
		top := stack[len(stack)-1]
		visited[top] = true

		board := LoadBoard(top)
		singleSolution, _ := HasSingleSolution(&board)
		if !singleSolution {
			stack = stack[:len(stack)-2]
			totalPops++
			pops++
			if pops == 10000 {
				stack = []string{config.SimpleString()}
				pops = 0
			}
			continue
		} else if board.numClues <= clues {
			break
		}

	}

}
