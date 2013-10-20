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

	color := "white"
	brd := board.NewStartingBoard()

	flag.Parse()

	fmt.Printf("white: %s\n", *white)
	fmt.Printf("black: %s\n", *black)

	for {
		var cmd *exec.Cmd
		var stdin io.Reader

		if color == "white" {
			fmt.Println("\nwhite move . . . ")
			cmd = exec.Command(*white)
			stdin = strings.NewReader("white " + brd.String() + "\n")
			color = "black"
		} else {
			fmt.Println("\nblack move . . . ")
			cmd = exec.Command(*black)
			stdin = strings.NewReader("black " + brd.String() + "\n")
			color = "white"
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

		var move board.Move

		fmt.Fscanf(cmdStdout, "%d %d", &move.Src, &move.Dest)
		fmt.Printf("\n(%d, %d)\n", move.Src, move.Dest)

		err = cmd.Wait()
		if err != nil {
			fmt.Printf("%s crashed with the following error:", color)
			log.Fatal(err)
		}

		// @TODO
		// verify the move is valid
		// get all legal moves for moving color
		// look up move in this list of legal moves

		brd = brd.MakeMove(move)
		fmt.Println(brd.Display())

		if brd.Checkmated(color) {
			fmt.Println("checkmate")
			break
		}

		// check for stalemate
		if brd.Stalemate(color) {
			fmt.Println("stalemate")
			break
		}
	}
}
