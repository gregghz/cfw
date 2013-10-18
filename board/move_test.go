package board

import (
	"testing"
)

func TestGetRookMoves(t *testing.T) {
	board := FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 WR 00 BP 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 WK 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves := board.GetRookMoves(20)

	if len(moves) != 10 {
		t.Errorf("expected 10 moves, found %d", len(moves))
	}

	valid_moves := []int{4,12,16,17,18,19,21,22,28,36}
	checkMoves(t, valid_moves, moves)
}

func TestBishopMoves(t *testing.T) {
	board := FromDisplay(
`00 00 00 00 00 00 00 00
00 00 WK 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 WB 00 00 00
00 00 00 00 00 BP 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves := board.GetBishopMoves(28)

	if len(moves) != 9 {
		t.Errorf("expect 9 moves, found %d", len(moves))
	}

	valid_moves := []int{7, 14, 19, 21, 35, 37, 42, 49, 56}
	checkMoves(t, valid_moves, moves)
}

func checkMoves(t *testing.T, valid_moves []int, moves []Move) {
	for _, move := range moves {
		found := false
		for _, valid := range valid_moves {
			if valid == move.Dest {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("invalid move: (%d, %d)", move.Src, move.Dest)
		}
	}

	for _, valid := range valid_moves {
		found := false
		for _, move := range moves {
			if valid == move.Dest {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("missing move: (%d, %d)", moves[0].Src, valid)
		}
	}
}
