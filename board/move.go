package board

import (
	"github.com/gregghz/cfw/pieces"
)

type Move struct {
	Src, Dest int
}

func colorMatch(piece, color string) bool {
	return (piece[0] == 'W' && color == "white") || (piece[0] == 'B' && color == "black")
}

func filter(moves []Move, f func(Move) bool) []Move {
	filtered := []Move{}

	for _, move := range moves {
		if f(move) {
			filtered = append(filtered, move)
		}
	}

	return filtered
}

func (board Board) findKing(color string) int {
	for i, piece := range board {
		if color == "white" && piece == "WK" {
			return i
		} else if color == "black" && piece == "BK" {
			return i
		}
	}

	return -1
}

func (board Board) exposesNoCheck(move Move) bool {
	checkBoard := board.MakeMove(move)

	color := "white"
	if board[move.Src][0] == 'B' {
		color = "black"
	}

	// find color king
	colorKing := checkBoard.findKing(color)

	opColor := "black"
	if color == "black" {
		opColor = "white"
	}

	opMoves := checkBoard.getUnfilteredMoves(opColor)

	for _, opMove := range opMoves {
		if opMove.Dest == colorKing {
			return false
		}
	}
	
	return true
}

func (board Board) getUnfilteredMoves(color string) []Move {
	moves := []Move{}
	
	for i, piece := range board {
		if !colorMatch(piece, color) {
			continue
		}

		switch piece {
		case "WP", "BP":
			moves = append(moves, board.GetPawnMoves(i)...)
		case "WR", "BR": // rooks
			moves = append(moves, board.GetRookMoves(i)...)
		case "WH", "BH": // horses/knights
			moves = append(moves, board.GetHorseMoves(i)...)
		case "WB", "BB": // bishops
			moves = append(moves, board.GetBishopMoves(i)...)
		case "WQ", "BQ": // queens
			moves = append(moves, board.GetBishopMoves(i)...)
			moves = append(moves, board.GetRookMoves(i)...)
		case "WK", "BK": // kings
			moves = append(moves, board.GetKingMoves(i)...)
		}
	}

	return moves
}

func (board Board) GetAllMoves(color string) []Move {
	moves := board.getUnfilteredMoves(color)
	return filter(moves, board.exposesNoCheck)
}

//  0  1  2  3  4  5  6  7
//  8  9 10 11 12 13 14 15
// 16 17 18 19 20 21 22 23
// 24 25 26 27 28 29 30 31
// 32 33 34 35 36 37 38 39
// 40 41 42 43 44 45 46 47
// 48 49 50 51 52 53 54 55
// 56 57 58 59 60 61 62 63

func (board Board) GetPawnMoves(i int) []Move {
	moves := []Move{}

	if board[i][0] == 'W' {
		if i-8 >= 0 && board[i-8] == "00" {
			moves = append(moves, Move{i, i - 8})

			if i > 47 && i < 56 && board[i-16] == "00" {
				moves = append(moves, Move{i, i - 16})
			}
		}

		switch i {
		case 0, 8, 16, 24, 32, 40, 48, 56:
		default:
			if i-9 >= 0 && board[i-9] != "00" && !colorMatch(board[i-9], "white") {
				moves = append(moves, Move{i, i - 9})
			}
		}

		switch i {
		case 7, 15, 23, 31, 39, 47, 55, 63:
		default:
			if i-7 >= 0 && board[i-7] != "00" && !colorMatch(board[i-7], "white") {
				moves = append(moves, Move{i, i - 7})
			}
		}
	} else { //black
		if i+8 <= 63 && board[i+8] == "00" {
			moves = append(moves, Move{i, i + 8})

			if i > 7 && i < 16 && board[i+16] == "00" {
				moves = append(moves, Move{i, i + 16})
			}
		}

		switch i {
		case 7, 15, 23, 31, 39, 47, 55, 63:
		default:
			if i+9 <= 63 && board[i+9] != "00" && !colorMatch(board[i+9], "black") {
				moves = append(moves, Move{i, i + 9})
			}
		}
		switch i {
		case 0, 8, 16, 24, 32, 40, 48, 56:
		default:
			if i+7 <= 63 && board[i+7] != "00" && !colorMatch(board[i+7], "black") {
				moves = append(moves, Move{i, i + 7})
			}
		}
	}

	return moves
}

type cmpf func(int)bool
func gte(n int) cmpf {
	return func(x int) bool {
		return x >= n
	}
}

func lte(n int) cmpf {
	return func(x int) bool {
		return x <= n
	}
}

func bish_cmpf(cmp cmpf, inc, xbound int) cmpf {
	return func(n int) bool {
		x := (n-inc) % 8

		if x == xbound {
			return false
		}

		return cmp(n)
	}
}

func (board Board) moveRange(i, inc int, cmp cmpf) []Move {
	moves := []Move{}
	piece := board[i]

	for j := i + inc; cmp(j); j += inc {
		if board[j] == pieces.Empty {
			moves = board.addValid(moves, Move{i, j})
		} else if piece[0] != board[j][0] {
			moves = board.addValid(moves, Move{i, j})
			break
		} else {
			break
		}
	}

	return moves
}

func (board Board) addValid(moves []Move, move Move) []Move {
	if board[move.Dest] == pieces.Empty || board[move.Src][0] != board[move.Dest][0] {
		moves = append(moves, move)
	}

	return moves
}

func (board Board) GetBishopMoves(i int) []Move {
	moves := []Move{}

	// up right
	moves = append(moves, board.moveRange(i, -7, bish_cmpf(gte(0), -7, 7))...)

	// up left
	moves = append(moves, board.moveRange(i, -9, bish_cmpf(gte(0), -9, 0))...)

	// down right
	moves = append(moves, board.moveRange(i, 9, bish_cmpf(lte(63), 9, 7))...)

	// down left
	moves = append(moves, board.moveRange(i, 7, bish_cmpf(lte(63), 7, 0))...)
	
	// return moves
	return moves
}

func (board Board) GetKingMoves(i int) []Move {
	moves := []Move{}

	if i > 7 {
		moves = board.addValid(moves, Move{i, i-8})

		switch i {
		case 7, 15, 23, 31, 39, 47, 55, 63:
		default:
			moves = board.addValid(moves, Move{i, i-7})
		}

		switch i {
		case 0, 8, 16, 24, 32, 40, 48, 56:
		default:
			moves = board.addValid(moves, Move{i, i-9})
		}
	}

	if i < 56 {
		moves = board.addValid(moves, Move{i, i+8})

		switch i {
		case 7, 15, 23, 31, 39, 47, 55, 63:
		default:
			moves = board.addValid(moves, Move{i, i+9})
		}

		switch i {
		case 0, 8, 16, 24, 32, 40, 48, 56:
		default:
			moves = board.addValid(moves, Move{i, i+7})
		}
	}

	switch i {
	case 7, 15, 23, 31, 39, 47, 55, 63:
	default:
		moves = board.addValid(moves, Move{i, i+1})
	}

	switch i {
	case 0, 8, 16, 24, 32, 40, 48, 56:
	default:
		moves = board.addValid(moves, Move{i, i-1})
	}

	return moves
}

func (board Board) GetRookMoves(i int) []Move {
	moves := []Move{}

	// up
	moves = append(moves, board.moveRange(i, -8, gte(0))...)

	// down
	moves = append(moves, board.moveRange(i, 8, lte(63))...)

	// left
	base_left := (i / 8) * 8
	moves = append(moves, board.moveRange(i, -1, gte(base_left))...)

	// right
	base_right := base_left + 7
	moves = append(moves, board.moveRange(i, 1, lte(base_right))...)

	return moves
}

func (board Board) GetHorseMoves(i int) []Move {
	moves := []Move{}

	right1 := lte(6)
	right2 := lte(5)
	left1 := gte(1)
	left2 := gte(2)

	addIf := func(guard cmpf, move int) {
		if guard(i%8) {
			moves = board.addValid(moves, Move{i, move})
		}
	}
	
	if i > 15 { // can move up 2
		addIf(right1, i-15)
		addIf(left1, i-17)
	}

	if i > 7 { // can move up 1
		addIf(right2, i-6)
		addIf(left2, i-10)
	}

	if i < 48 { // can move down 2
		addIf(right1, i+17)
		addIf(left1, i+15)
	}

	if i < 56 { // can move down 1
		addIf(right2, i+10)
		addIf(left2, i+6)
	}

	return moves
}
