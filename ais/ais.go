package ais

import (
	"github.com/greggoryhz/cfw/board"
	"math"
)

var PieceValue map[string]int

const (
	Pawn   = 100
	Knight = 320
	Bishop = 325
	Rook   = 500
	Queen  = 975
)

func pieceScore(piece string) int {
	switch piece {
	case "WP","BP":
		return Pawn
	case "WR", "BR":
		return Rook
	case "WH", "BH":
		return Knight
	case "WB", "BB":
		return Bishop
	case "WQ", "BQ":
		return Queen
	default:
		return 0
	}
}

func ScoreBoard(boardAfterMove board.Board, move board.Move) int {
	score := 0

	for _, piece := range boardAfterMove {
		score += pieceScore(piece)
	}

	if boardAfterMove[move.Dest][0] == 'W' {
		if boardAfterMove.Checkmated("black") {
			score = math.MaxInt64
		}

		return score
	} else {
		if boardAfterMove.Checkmated("white") {
			score = math.MaxInt64
		}

		return -score
	}
}

// piece square tables

var PawnTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	5, 5, 10, 27, 27, 10, 5, 5,
	0, 0, 0, 25, 25, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	5, 10, 10, -25, -25, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var KnightTable = []int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-30, 0, 10, 15, 15, 10, 0, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-50, -40, -20, -30, -30, -20, -40, -50,
}

var BishopTable = []int{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 10, 10, 5, 0, -10,
	-10, 5, 5, 10, 10, 5, 5, -10,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-20, -10, -40, -10, -10, -40, -10, -20,
}

var KingTable = []int{
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	20, 20, 0, 0, 0, 0, 20, 20,
	20, 30, 10, 0, 0, 10, 30, 20,
}

var KingTableEndGame = []int{
	-50, -40, -30, -20, -20, -30, -40, -50,
	-30, -20, -10, 0, 0, -10, -20, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -30, 0, 0, 0, 0, -30, -30,
	-50, -30, -30, -30, -30, -30, -30, -50,
}
