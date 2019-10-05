package sudoku

import (
	"fmt"
	"strings"
)

type Board struct {
	values   [NUM_CELLS]int
	numClues int
}

const (
	ALL       = 0x1FF
	NUM_CELLS = 81
)

var decoder [1 << 9]int

func init() {
	for i := 1; i <= 9; i++ {
		decoder[1<<uint(i-1)] = i
	}
}

func NewBoard() Board {
	var board Board

	for i := 0; i < NUM_CELLS; i++ {
		board.values[i] = ALL
	}

	board.numClues = 0

	return board
}

func LoadBoard(boardString string) Board {
	board := NewBoard()

	if len(boardString) != 81 {
		return board
	}

	charToValueMap := map[rune]int{
		'0': ALL,
		'1': Encode(1),
		'2': Encode(2),
		'3': Encode(3),
		'4': Encode(4),
		'5': Encode(5),
		'6': Encode(6),
		'7': Encode(7),
		'8': Encode(8),
		'9': Encode(9),
		'.': ALL,
	}

	for i := 0; i < NUM_CELLS; i++ {
		val := charToValueMap[rune(boardString[i])]
		if val == 0 {
			val = ALL
		}
		board.SetCell(i, val)
	}

	return board
}

func Decode(value int) int {
	if value < 0 || value > len(decoder) {
		return 0
	}

	return decoder[value]
}

func Encode(value int) int {
	if value < 0 || value > 9 {
		return 0
	}

	return 1 << (value - 1)
}

func DecodeBoard(board *Board) Board {
	decodedBoard := NewBoard()
	decodedBoard.numClues = board.numClues

	for i := 0; i < NUM_CELLS; i++ {
		decodedBoard.values[i] = Decode(board.values[i])
	}

	return decodedBoard
}

func EncodeBoard(board *Board) Board {
	encodedBoard := NewBoard()
	encodedBoard.numClues = board.numClues

	for i := 0; i < NUM_CELLS; i++ {
		encodedBoard.values[i] = Encode(board.values[i])
	}

	return encodedBoard
}

func (board *Board) decodeCell(index int) int {
	return decoder[board.values[index]]
}

func (board *Board) decode() [NUM_CELLS]int {
	var result [NUM_CELLS]int
	for i := 0; i < NUM_CELLS; i++ {
		result[i] = board.decodeCell(i)
	}

	return result
}

func (board *Board) Clear() {
	for i := 0; i < NUM_CELLS; i++ {
		board.values[i] = ALL
	}
	board.numClues = 0
}

func (board Board) EmptySpaces() int {
	return NUM_CELLS - board.numClues
}

func (board *Board) SetCell(index, newValue int) {
	oldVal := board.values[index]
	board.values[index] = newValue
	if decoder[newValue] > 0 && decoder[oldVal] == 0 {
		board.numClues++
	} else if decoder[newValue] == 0 && decoder[oldVal] > 0 {
		board.numClues--
	}
}

func (board *Board) Cell(index int) int {
	return board.values[index]
}

func (board *Board) IsRowValid(row int) bool {
	check := 0
	for i := row * 9; i < (row+1)*9; i++ {
		digit := Decode(board.values[i])
		if digit > 0 && digit <= 9 {
			mask := 1 << (digit - 1)
			if check&mask != 0 {
				return false
			}
			check |= mask
		}
	}
	return true
}

func (board *Board) IsColumnValid(col int) bool {
	check := 0
	for i := 0; i < NUM_CELLS; i += 9 {
		digit := Decode(board.values[i])
		if digit > 0 && digit <= 9 {
			mask := 1 << (digit - 1)
			if check&mask != 0 {
				return false
			}
			check |= mask
		}
	}
	return true
}

func (board *Board) IsRegionValid(region int) bool {
	check := 0
	regionRow := region / 3
	regionCol := region % 3

	for i := 0; i < 9; i++ {
		digit := Decode(board.values[regionRow*27+regionCol*3+(i/3)*9+(i%3)])
		if digit > 0 && digit <= 9 {
			mask := 1 << (digit - 1)
			if check&mask != 0 {
				return false
			}
			check |= mask
		}
	}
	return true
}

func (board *Board) IsValid() bool {
	for i := 0; i < 9; i++ {
		if !board.IsRowValid(i) ||
			!board.IsColumnValid(i) ||
			!board.IsRegionValid(i) {

			return false
		}
	}
	return true
}

func (board *Board) IsFull() bool {
	return board.numClues == NUM_CELLS
}

func (board *Board) IsSolved() bool {
	return board.IsFull() && board.IsValid()
}

func (board *Board) GetCandidates(index int) []int {
	return make([]int, 0, 9)
}

func (board *Board) Copy() Board {
	newBoard := NewBoard()
	newBoard.numClues = board.numClues
	for i := 0; i < NUM_CELLS; i++ {
		newBoard.values[i] = board.values[i]
	}
	return newBoard
}

func (board *Board) SimpleString() string {
	var b strings.Builder
	for i := 0; i < NUM_CELLS; i++ {
		fmt.Fprintf(&b, "%d", Decode(board.values[i]))
	}
	return b.String()
}
