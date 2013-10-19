package board

import (
	"github.com/greggoryhz/cfw/pieces"
)

type Move struct {
	Src, Dest int
}

func colorMatch(piece, color string) bool {
	return (piece[0] == 'W' && color == "white") || (piece[0] == 'B' && color == "black")
}

func (board Board) GetAllMoves(color string) []Move {
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
			moves = append(moves, Move{i, j})
		} else if piece[0] != board[j][0] {
			moves = append(moves, Move{i, j})
			break
		} else {
			break
		}
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

	right := func(t int) {
		switch i {
		case 7, 15, 23, 31, 39, 47, 55, 63:
		default:
			moves = append(moves, Move{i, t})
		}
	}

	left := func(t int) {
		switch i {
		case 7, 15, 23, 31, 39, 47, 55, 63:
		default:
			moves = append(moves, Move{i, t})
		}
	}

	// up
	if i > 7 {
		moves = append(moves, Move{i, i-8})
		right(i-7)
		left(i-9)
	}

	right(i+1)

	// down
	if i < 56 {
		moves = append(moves, Move{i, i+8})
		right(i+9)
		left(i-7)
	}

	left(i-1)

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

	if i > 15 { // can move up 2
		if i%8 < 7 { // can move right 1
			moves = append(moves, Move{i, i - 15})
		}

		if i%8 > 0 { // can move left 1
			moves = append(moves, Move{i, i - 17})
		}
	}

	if i > 7 { // can move up 1
		if i%8 < 6 { // can move right 2
			moves = append(moves, Move{i, i - 6})
		}
		if i%8 > 1 { // can move left 2
			moves = append(moves, Move{i, i - 10})
		}
	}

	if i < 48 { // can move down 2
		if i%8 < 7 { // can move right 1
			moves = append(moves, Move{i, i + 17})
		}
		if i%8 > 0 { // can move left 1
			moves = append(moves, Move{i, i + 15})
		}
	}

	if i < 56 { // can move down 1
		if i%8 < 6 { // can move right 2
			moves = append(moves, Move{i, i + 10})
		}
		if i%8 > 1 { // can move left 2
			moves = append(moves, Move{i, i + 6})
		}
	}

	return moves
}
