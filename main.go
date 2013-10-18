package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/greggoryhz/cfw/board"
	"os"
)

func main() {
	white := flag.String("white", "ais/random/random", "the path to white's executable.")
	black := flag.String("black", "ais/random/random", "the path to black's executable.")

	whiteTurn := true
	brd := board.NewStartingBoard()

	flag.Parse()

	fmt.Printf("white: %s\n", *white)
	fmt.Printf("black: %s\n", *black)

	for {
		var cmd *exec.Cmd
		var stdin io.Reader

		if whiteTurn {
			fmt.Println("\nwhite move . . . ")
			cmd = exec.Command(*white)
			stdin = strings.NewReader("white " + brd.String() + "\n")
			whiteTurn = false
		} else {
			fmt.Println("\nblack move . . . ")
			cmd = exec.Command(*black)
			stdin = strings.NewReader("black " + brd.String() + "\n")
			whiteTurn = true
		}

		cmd.Stdin = stdin
		cmd.Stderr = os.Stderr

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
		newBoard := make(board.Board, 64)
		var fromIndex int
		var toIndex int

		fmt.Fscanf(cmdStdout, "%d %d", &fromIndex, &toIndex)
		fmt.Printf("\n(%d, %d)\n", fromIndex, toIndex)

		for i := 0; i < 64; i++ {
			fmt.Fscanf(cmdStdout, "%s", &newBoard[i])
		}

		err = cmd.Wait()
		if err != nil {
			if whiteTurn {
				fmt.Println("White crashed with the following error:")
				log.Fatal(err)
			} else {
				fmt.Println("Black crashed with the following error:")
				log.Fatal(err)
			}
		}

		brd = newBoard
		fmt.Println(brd.Display())

		// do all verification here first
		// @TODO
		if whiteTurn && len(newBoard.GetAllMoves("white")) == 0 {
			fmt.Println("black wins")
			break
		}

		if !whiteTurn && len(newBoard.GetAllMoves("black")) == 0 {
			fmt.Println("white wins")
			break
		}
	}
}
