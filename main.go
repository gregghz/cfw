package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"log"
	"io"
)

const (
	WhitePawn = "WP"
	WhiteKnight = "WH"
	WhiteBishop = "WB"
	WhiteRook = "WR"
	WhiteQueen = "WQ"
	WhiteKing = "WK"

	Empty = "00"

	BlackPawn = "BP"
	BlackKnight = "BH"
	BlackBishop = "BB"
	BlackRook = "BR"
	BlackQueen = "BQ"
	BlackKing = "BK"
)

type Piece string

type Board struct {
	Squares []string

	readerIndex int
}

func NewBoard() *Board {
	return &Board{
		[]string{
			BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook,
			BlackPawn, BlackPawn,   BlackPawn,   BlackPawn,  BlackPawn, BlackPawn,   BlackPawn,   BlackPawn,
			Empty,     Empty,       Empty,       Empty,      Empty,     Empty,       Empty,       Empty,
			Empty,     Empty,       Empty,       Empty,      Empty,     Empty,       Empty,       Empty,
			Empty,     Empty,       Empty,       Empty,      Empty,     Empty,       Empty,       Empty,
			Empty,     Empty,       Empty,       Empty,      Empty,     Empty,       Empty,       Empty,
			WhitePawn, WhitePawn,   WhitePawn,   WhitePawn,  WhitePawn, WhitePawn,   WhitePawn,   WhitePawn,
			WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook,
		}, 0,
	}
}

func main() {
	white := flag.String("white", "bin/random", "the path to white's executable.")
	black := flag.String("black", "bin/random", "the path to black's executable.")

	whiteTurn := true
	board := NewBoard()

	flag.Parse()

	fmt.Printf("white: %s\n", *white)
	fmt.Printf("black: %s\n", *black)

	for {
		var cmd *exec.Cmd
		var stdin io.Reader
		
		if (whiteTurn) {
			cmd = exec.Command(*white)
			stdin = strings.NewReader("white " + strings.Join(board.Squares, " "))
			whiteTurn = false
		} else {
			cmd = exec.Command(*black)
			stdin = strings.NewReader("black " + strings.Join(board.Squares, " "))
			whiteTurn = true
		}

		cmd.Stdin = stdin

		// get the stdout pipe from cmd to read their response
		cmdStdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		// read the stdout of cmd
		// should be a single line containing a FromIndex, a ToIndex (this pair represents the move)
		// followed by 64 "pieces" representing the game board AFTER the move
		data := make([]byte, 2*64+63)
		var fromIndex int
		var toIndex int

		fmt.Fscanf(cmdStdout, "%d", &fromIndex)
		fmt.Fscanf(cmdStdout, "%d", &toIndex)
		
		_, err = cmdStdout.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))
	}
}
