package board

import (
	"fmt"
)

func ExampleNewStartingBoard() {
	board := NewStartingBoard()
	fmt.Println(board.Display())
	// Output:
	// BR BH BB BQ BK BB BH BR
	// BP BP BP BP BP BP BP BP
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 00 00 00 00
	// WP WP WP WP WP WP WP WP
	// WR WH WB WQ WK WB WH WR
}

func ExampleMakeMove() {
	board := NewStartingBoard()
	board = board.MakeMove(Move{8, 16})
	fmt.Println(board.Display())
	// Output:
	// BR BH BB BQ BK BB BH BR
	// 00 BP BP BP BP BP BP BP
	// BP 00 00 00 00 00 00 00
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 00 00 00 00
	// WP WP WP WP WP WP WP WP
	// WR WH WB WQ WK WB WH WR
}
