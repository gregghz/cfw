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
		t.Errorf("expected 9 moves, found %d", len(moves))
	}

	valid_moves := []int{7, 14, 19, 21, 35, 37, 42, 49, 56}
	checkMoves(t, valid_moves, moves)

	board = FromDisplay(
`00 00 00 00 00 00 00 WK
WH 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 BB 00 00 00`)

	moves = board.GetBishopMoves(60)

	if len(moves) != 7 {
		t.Errorf("expected 7 moves, found %d", len(moves))
	}

	valid_moves = []int{24, 33, 39, 42, 46, 51, 53}
	checkMoves(t, valid_moves, moves)

	t.Logf("down right")
	board = FromDisplay(
`WB 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves = board.GetBishopMoves(0)

	if len(moves) != 7 {
		t.Errorf("expected 7 moves, found %d", len(moves))
	}

	valid_moves = []int{9, 18, 27, 36, 45, 54, 63}
	checkMoves(t, valid_moves, moves)

	t.Logf("down left")
	board = FromDisplay(
`00 00 00 00 00 00 00 WB
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves = board.GetBishopMoves(7)

	if len(moves) != 7 {
		t.Errorf("expected 7 moves, found %d", len(moves))
	}

	valid_moves = []int{14, 21, 28, 35, 42, 49, 56}
	checkMoves(t, valid_moves, moves)

	t.Logf("up right")
	board = FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
WB 00 00 00 00 00 00 00`)

	moves = board.GetBishopMoves(56)

	if len(moves) != 7 {
		t.Errorf("expected 7 moves, found %d", len(moves))
	}

	valid_moves = []int{49, 42, 35, 28, 21, 14, 7}
	checkMoves(t, valid_moves, moves)

	t.Logf("up left")
	board = FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 WB`)

	moves = board.GetBishopMoves(63)

	if len(moves) != 7 {
		t.Errorf("expected 7 moves, found %d", len(moves))
	}

	valid_moves = []int{54, 45, 36, 27, 18, 9, 0}
	checkMoves(t, valid_moves, moves)
}

func TestHorseMoves(t *testing.T) {
	t.Logf("8 moves")
	board := FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 WH 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves := board.GetHorseMoves(27)

	if len(moves) != 8 {
		t.Errorf("expected 8 moves, found %d", len(moves))
	}

	valid_moves := []int{10, 12, 17, 21, 33, 37, 42, 44}
	checkMoves(t, valid_moves, moves)

	t.Logf("top left")
	board = FromDisplay(
`WH 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves = board.GetHorseMoves(0)

	if len(moves) != 2 {
		t.Errorf("expected 2 moves, found %d", len(moves))
	}

	valid_moves = []int{10, 17}
	checkMoves(t, valid_moves, moves)

	t.Logf("top right")
	board = FromDisplay(
`00 00 00 00 00 00 00 WH
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves = board.GetHorseMoves(7)

	if len(moves) != 2 {
		t.Errorf("expected 2 moves, found %d", len(moves))
	}

	valid_moves = []int{13, 22}
	checkMoves(t, valid_moves, moves)

	t.Logf("bottom left")
	board = FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
WH 00 00 00 00 00 00 00`)

	moves = board.GetHorseMoves(56)

	if len(moves) != 2 {
		t.Errorf("expected 2 moves, found %d", len(moves))
	}

	valid_moves = []int{41, 50}
	checkMoves(t, valid_moves, moves)

	t.Logf("bottom right")
	board = FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 WH`)

	moves = board.GetHorseMoves(63)

	if len(moves) != 2 {
		t.Errorf("expected 2 moves, found %d", len(moves))
	}

	valid_moves = []int{46, 53}
	checkMoves(t, valid_moves, moves)

	t.Logf("limited in the middle")
	board = FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 WK 00 00 00
00 00 00 00 00 00 00 00
00 00 00 WH 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 BP 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00`)

	moves = board.GetHorseMoves(27)

	if len(moves) != 7 {
		t.Errorf("expected 7 moves, found %d", len(moves))
	}

	valid_moves = []int{10, 17, 21, 33, 37, 42, 44}
	checkMoves(t, valid_moves, moves)
}

func TestKingMoves(t *testing.T) {
	board := FromDisplay(
`00 00 00 00 00 00 00 00
00 00 00 00 BK 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 WH`)

	moves := board.GetKingMoves(12)

	if len(moves) != 8 {
		t.Errorf("expected 8 moves, found %d", len(moves))
	}

	valid_moves := []int{3, 4, 5, 11, 13, 19, 20, 21}
	checkMoves(t, valid_moves, moves)
}

//  0  1  2  3  4  5  6  7
//  8  9 10 11 12 13 14 15
// 16 17 18 19 20 21 22 23
// 24 25 26 27 28 29 30 31
// 32 33 34 35 36 37 38 39
// 40 41 42 43 44 45 46 47
// 48 49 50 51 52 53 54 55
// 56 57 58 59 60 61 62 63

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
			t.Errorf("missing move: (%d)", valid)
		}
	}
}
