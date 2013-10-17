package board

import (
	"cfw/piece"
	"fmt"
	"strings"
)

type Board []string

func NewBoard() Board {
	return Board{
		piece.BlackRook, piece.BlackKnight, piece.BlackBishop, piece.BlackQueen, piece.BlackKing, piece.BlackBishop, piece.BlackKnight, piece.BlackRook,
		piece.BlackPawn, piece.BlackPawn,   piece.BlackPawn,   piece.BlackPawn,  piece.BlackPawn, piece.BlackPawn,   piece.BlackPawn,   piece.BlackPawn,
		piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,      piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,
		piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,      piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,
		piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,      piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,
		piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,      piece.Empty,     piece.Empty,       piece.Empty,       piece.Empty,
		piece.WhitePawn, piece.WhitePawn,   piece.WhitePawn,   piece.WhitePawn,  piece.WhitePawn, piece.WhitePawn,   piece.WhitePawn,   piece.WhitePawn,
		piece.WhiteRook, piece.WhiteKnight, piece.WhiteBishop, piece.WhiteQueen, piece.WhiteKing, piece.WhiteBishop, piece.WhiteKnight, piece.WhiteRook,
	}
}

func (board Board) String() string {
	return fmt.Sprintf("%s", strings.Join(board, " "))
}
