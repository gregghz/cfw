package main

import (
	"fmt"
	"github.com/greggoryhz/cfw/board"
	"github.com/greggoryhz/cfw/ais"
)

func main() {
	brd := board.NewBoard()

	var color string
	fmt.Scanf("%s", &color)

	for i := 0; i < 64; i++ {
		fmt.Scanf("%s", &brd[i])
	}

	moves := brd.GetAllMoves(color)
	mv := selectBest(brd, moves)

	fmt.Printf("%d %d\n", mv.Src, mv.Dest)
}

func selectBest(brd board.Board, moves []board.Move) board.Move {
	best := moves[0]
	score := ais.ScoreBoard(brd, best)

	for _, move := range moves[1:] {
		s := ais.ScoreBoard(brd.MakeMove(move), move)
		if s > score {
			best = move
			score = s
		}
	}

	return best
}
