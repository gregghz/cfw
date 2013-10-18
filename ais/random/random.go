package main

import (
	"fmt"
	"github.com/greggoryhz/cfw/board"
	"math/rand"
	"time"
)

type move [2]int

func main() {
	rand.Seed(time.Now().Unix())
	brd := board.NewBoard()

	var color string
	fmt.Scanf("%s", &color)

	// read a board out of stdin
	for i := 0; i < 64; i++ {
		fmt.Scanf("%s", &brd[i])
	}

	moves := brd.GetAllMoves(color)
	mv := moves[rand.Intn(len(moves))]

	fmt.Printf("%d %d %s\n", mv.Src, mv.Dest, brd.MakeMove(mv).String())
}
