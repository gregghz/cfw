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
		var move board.Move

		// @TODO
		// verify the move is valid
		// get all legal moves for moving color
		// look up move in this list of legal moves

		fmt.Fscanf(cmdStdout, "%d %d", &move.Src, &move.Dest)
		fmt.Printf("\n(%d, %d)\n", move.Src, move.Dest)

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

		brd = brd.MakeMove(move)
		fmt.Println(brd.Display())

		// check for stalemate
		if (whiteTurn && len(brd.GetAllMoves("white")) == 0) ||
			(!whiteTurn && len(brd.GetAllMoves("black")) == 0) {
			fmt.Println("stalemate")
			break
		}
	}
}
