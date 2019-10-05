package sudoku

import "testing"

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	if len(board.values) != NUM_CELLS {
		t.Errorf("len(board.values) = %d, want %d", len(board.values), NUM_CELLS)
		return
	}

	testBoardEmpty(t, board)
}

func testBoardEmpty(t *testing.T, board Board) {
	//All values in empty board should be ALL (0x1FF)
	for i := 0; i < NUM_CELLS; i++ {
		if board.values[i] != ALL {
			t.Errorf("board.values[%d] = %d, want %d", i, board.values[i], ALL)
			return
		}
	}
}

func TestClear(t *testing.T) {
	board := NewBoard()

	for i := 0; i < NUM_CELLS; i++ {
		board.values[i] = 1 << (i % 9)
	}

	//Test that all values are cleared
	board.Clear()
	testBoardEmpty(t, board)
}

func TestEmptySpace(t *testing.T) {
	board := NewBoard()

	if board.EmptySpaces() != NUM_CELLS {
		t.Errorf("board.EmptySpace() = %d, want %d", board.EmptySpaces(), NUM_CELLS)
		return
	}

	for i := 0; i < NUM_CELLS; i++ {
		board.values[i] = 1 << (i % 9)
	}

	if board.EmptySpaces() != 0 {
		t.Errorf("board.EmptySpace() = %d, want %d", board.EmptySpaces(), 0)
		return
	}

	board = NewBoard()
	for i := 0; i < NUM_CELLS/2; i++ {
		board.values[i] = 1 << (i % 9)
	}

	if board.EmptySpaces() != NUM_CELLS/2 {
		t.Errorf("board.EmptySpace() = %d, want %d", board.EmptySpaces(), NUM_CELLS/2)
		return
	}
}

func TestLoadBoard(t *testing.T) {
	testBoardEmpty(t, LoadBoard(""))
	testBoardEmpty(t, LoadBoard("abc"))
	testBoardEmpty(t, LoadBoard("1"))
	testBoardEmpty(t, LoadBoard("12345678912345678912345678912345678912345678912345678912345678912345678912345678"))
	testBoardEmpty(t, LoadBoard("1234567891234567891234567891234567891234567891234567891234567891234567891234567890"))
	testBoardEmpty(t, LoadBoard("12345678912345678912345678912345678912345678912345678912345678912345678912345678999999"))
	testBoardEmpty(t, LoadBoard("................................................................................."))
	testBoardEmpty(t, LoadBoard("000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	testBoardEmpty(t, LoadBoard("askjdflkajhfglkjahldfjahlfkgjhalkjfhglakjfhglakjfhgliaurhlkzjhvliuserlkjlbuhzsdjh"))

	board := LoadBoard("168249753359817462472356819685493127291678345734521698816935274947182536523764981")
	expectedValues := []int{1, 6, 8, 2, 4, 9, 7, 5, 3, 3, 5, 9, 8, 1, 7, 4, 6, 2, 4, 7, 2, 3, 5, 6, 8, 1, 9, 6, 8, 5, 4, 9, 3, 1, 2, 7, 2, 9, 1, 6, 7, 8, 3, 4, 5, 7, 3, 4, 5, 2, 1, 6, 9, 8, 8, 1, 6, 9, 3, 5, 2, 7, 4, 9, 4, 7, 1, 8, 2, 5, 3, 6, 5, 2, 3, 7, 6, 4, 9, 8, 1}

	for i := 0; i < NUM_CELLS; i++ {
		if board.values[i] != Encode(expectedValues[i]) {
			t.Errorf("board.values[%d] = %d, want %d", i, board.values[i], Encode(expectedValues[i]))
			return
		}
	}
}

func TestIsFull(t *testing.T) {
	board := NewBoard()

	if board.IsFull() {
		t.Errorf("board.IsFull() = %t, want %t", board.IsFull(), false)
		return
	}

	board = LoadBoard("238761594491825637657493821389614275712358469546972318823149756165287943974536182")
	if !board.IsFull() {
		t.Errorf("board.IsFull() = %t, want %t", board.IsFull(), true)
		return
	}

}
