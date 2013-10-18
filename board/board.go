package board

import (
	"fmt"
	"github.com/greggoryhz/cfw/pieces"
	"strings"
)

type Board []string

func NewBoard() Board {
	return make(Board, 64)
}

func NewStartingBoard() Board {
	return Board{
		pieces.BlackRook, pieces.BlackKnight, pieces.BlackBishop, pieces.BlackQueen, pieces.BlackKing, pieces.BlackBishop, pieces.BlackKnight, pieces.BlackRook,
		pieces.BlackPawn, pieces.BlackPawn, pieces.BlackPawn, pieces.BlackPawn, pieces.BlackPawn, pieces.BlackPawn, pieces.BlackPawn, pieces.BlackPawn,
		pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty,
		pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty,
		pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty,
		pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty, pieces.Empty,
		pieces.WhitePawn, pieces.WhitePawn, pieces.WhitePawn, pieces.WhitePawn, pieces.WhitePawn, pieces.WhitePawn, pieces.WhitePawn, pieces.WhitePawn,
		pieces.WhiteRook, pieces.WhiteKnight, pieces.WhiteBishop, pieces.WhiteQueen, pieces.WhiteKing, pieces.WhiteBishop, pieces.WhiteKnight, pieces.WhiteRook,
	}
}

func (board Board) String() string {
	return fmt.Sprintf("%s", strings.Join(board, " "))
}

func (board Board) Display() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		strings.Join(board[:8], " "), strings.Join(board[8:16], " "),
		strings.Join(board[16:24], " "), strings.Join(board[24:32], " "),
		strings.Join(board[32:40], " "), strings.Join(board[40:48], " "),
		strings.Join(board[48:56], " "), strings.Join(board[56:64], " "))
}

func (board Board) MakeMove(mv Move) Board {
	// copy the slice
	newBoard := NewBoard()
	copy(newBoard, board)

	piece := newBoard[mv.Src]
	newBoard[mv.Dest] = piece
	newBoard[mv.Src] = pieces.Empty

	// check for a queen'd pawn
	if piece[1] == 'P' {
		if piece[0] == 'W' && mv.Dest < 8 {
			newBoard[mv.Dest] == pieces.WhiteQueen
		} else if piece[0] == 'B' && mv.Dest > 55 {
			newBoard[mv.Dest] == pieces.BlackQueen
		}
	}

	return newBoard
}
