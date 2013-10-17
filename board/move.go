package board

type Move struct {
	Src, Dest int
}

func colorMatch(piece, color string) bool {
	return (piece[0] == 'W' && color == "white") || (piece[0] == 'B' && color == "black")
}

func (board Board) GetAllMoves(color string) []Move {
	moves := []Move{}

	for i,piece := range board {
		if !colorMatch(piece, color) {
			continue
		}

		switch piece {
		case "WP", "BP":
			moves = append(moves, board.GetPawnMoves(i,color)...)
		case "WR", "BR": // rooks
		case "WH", "BH": // horses/knights
			moves = append(moves, board.GetHorseMoves(i)...)
		case "WB", "BB": // bishops
		case "WQ", "BQ": // queens
		case "WK", "BK": // kings
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

func (board Board) GetPawnMoves(i int, color string) []Move {
	moves := []Move{}

	if color == "white" {
		if i > 47 && i < 56 {
			if board[i-8] == "00" {
				moves = append(moves, Move{i,i-8})
			}
			if board[i-8] == "00" && board[i-16] == "00" {
				moves = append(moves, Move{i,i-16})
			}
		} else {
			if i-8 >= 0 && board[i-8] == "00" {
				moves = append(moves, Move{i,i-8})
			}
		}
		switch i {
		case 0: case 8: case 16: case 24:
		case 32: case 40: case 48: case 56:
		default:
			if i-9 >= 0 && board[i-9] != "00" && !colorMatch(board[i-9],color) {
				moves = append(moves, Move{i,i-9});
			}
		}
		switch i {
		case 7: case 15: case 23: case 31:
		case 39: case 47: case 55: case 63:
		default:
			if i-7 >= 0 && board[i-7] != "00" && !colorMatch(board[i-7],color) {
				moves = append(moves, Move{i,i-7});
			}
		}
	} else { //black
		if i > 7 && i < 16 {
			if board[i+8] == "00" {
				moves = append(moves, Move{i,i+8})
			}
			if board[i+8] == "00" && board[i+16] == "00" {
				moves = append(moves, Move{i,i+16})
			}
		} else {
			if i+8 >= 0 && board[i+8] == "00" {
				moves = append(moves, Move{i,i+8})
			}
		}
		switch i {
		case 7: case 15: case 23: case 31:
		case 39: case 47: case 55: case 63:
		default:
			if i+9 <= 63 && board[i+9] != "00" && !colorMatch(board[i+9],color) {
				moves = append(moves, Move{i,i+9});
			}
		}
		switch i {
		case 0: case 8: case 16: case 24:
		case 32: case 40: case 48: case 56:
		default:
			if i+7 <= 63 && board[i+7] != "00" && !colorMatch(board[i+7],color) {
				moves = append(moves, Move{i,i+7});
			}
		}
	}

	return moves
}

func (board Board) GetHorseMoves(i int) []Move {
	moves := []Move{}

	if i > 15 { // can move up 2
		if i % 8 < 7 { // can move right 1
			moves = append(moves, Move{i, i-15})
		}

		if i % 8 > 0 { // can move left 1
			moves = append(moves, Move{i, i-17})
		}
	}

	if i > 7 { // can move up 1
		if i % 8 < 6 { // can move right 2
			moves = append(moves, Move{i, i-6})
		}
		if i % 8 > 1 { // can move left 2
			moves = append(moves, Move{i, i-10})
		}
	}

	if i < 48 { // can move down 2
		if i % 8 < 7 { // can move right 1
			moves = append(moves, Move{i, i+17})
		}
		if i % 8 > 0 { // can move left 1
			moves = append(moves, Move{i, i+15})
		}
	}

	if i < 56 { // can move down 1
		if i % 8 < 6 { // can move right 2
			moves = append(moves, Move{i, i+10})
		}
		if i % 8 > 1 { // can move left 2
			moves = append(moves, Move{i, i+6})
		}
	}

	return moves
}
