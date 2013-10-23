package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	//"strings"

	"github.com/greggoryhz/cfw/board"
	c "github.com/greggoryhz/cfw/communicator"
	"os"
)

type Result string

const (
	WhiteWin = "WhiteWin"
	BlackWin = "BlackWin"
	Stalemate = "Stalemate"

	WhiteIllegalMove = "WhiteIllegalMove"
	BlackIllegalMove = "BlackIllegalMove"

	WhiteCrash = "WhiteCrash"
	BlackCrash = "BlackCrash"
)

type Settings struct {
	final_board *bool
	intermediate_boards *bool
	games *int
}

func RunGame(white, black string, settings *Settings, com c.Communicator) Result {
	color := "white"
	brd := board.NewStartingBoard()

	defer func () {
		if !(*(settings.intermediate_boards)) && *(settings.final_board) {
			fmt.Println(brd.Display())
		}
	}()

	for {
		var cmd *exec.Cmd
		var stdin io.Reader

		legalMoves := brd.GetAllMoves(color)

		if color == "white" {
			cmd = exec.Command(white)
			stdin = com.GenerateRequest(true,brd)
			//stdin = strings.NewReader("white " + brd.String() + "\n")
			color = "black"
		} else {
			cmd = exec.Command(black)
			stdin = com.GenerateRequest(false,brd)
			//stdin = strings.NewReader("black " + brd.String() + "\n")
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

		//fmt.Fscanf(cmdStdout, "%d %d", &move.Src, &move.Dest)
		move = com.ProcessResponse(cmdStdout)

		err = cmd.Wait()
		if err != nil {
			fmt.Printf("%s crashed with the following error:", color)
			log.Fatal(err)
		}

		// verify valid move
		legal := false
		for _, lm := range legalMoves {
			if move.Src == lm.Src && move.Dest == lm.Dest {
				legal = true
			}
		}

		if !legal {
			if color == "white" {
				return BlackIllegalMove
			} else {
				return WhiteIllegalMove
			}
		}

		brd = brd.MakeMove(move)

		if *(settings.intermediate_boards) {
			fmt.Println(brd.Display() + "\n")
		}

		if brd.Checkmated(color) {
			if color == "white" {
				return BlackWin
			} else {
				return WhiteWin
			}
		}

		// check for stalemate
		if brd.Stalemate(color) {
			return Stalemate
		}
	}
}

func PrintResults(results map[Result]int) {
	fmt.Printf("=== Results ===\n")
	fmt.Printf("White win: %d\n", results[WhiteWin])
	fmt.Printf("Black win: %d\n", results[BlackWin])
	fmt.Printf("Stalemate: %d\n\n", results[Stalemate])

	fmt.Printf("White Illegal: %d\n", results[WhiteIllegalMove])
	fmt.Printf("Black Illegal: %d\n", results[BlackIllegalMove])
	fmt.Printf("White Crash: %d\n", results[WhiteCrash])
	fmt.Printf("Black Crash: %d\n\n", results[BlackCrash])
}

func main() {
	white := flag.String("white", "ais/random/random", "the path to white's executable.")
	black := flag.String("black", "ais/random/random", "the path to black's executable.")
	//com := flag.String("com","text","What type of communication")

	settings := &Settings {
		flag.Bool("final", true, "show the final game board."),
		flag.Bool("all-boards", true, "show intermediate game boards (implies -final)."),
		flag.Int("games", 1, "how many games to run."),
	}

	flag.Parse()

	fmt.Printf("white: %s\n", *white)
	fmt.Printf("black: %s\n", *black)

	//var communicator *c.Communicator
	//if *com == "json" {
	//	communicator := c.JsonCommunicator{}
	//} else {
		communicator := &c.TextCommunicator{}
	//}

	results := map[Result]int{
		WhiteWin: 0,
		BlackWin: 0,
		Stalemate: 0,
		WhiteIllegalMove: 0,
		BlackIllegalMove: 0,
		WhiteCrash: 0,
		BlackCrash: 0,
	}

	for i := 0; i < *settings.games; i++ {
		fmt.Printf("game #%d\n", i+1)
		results[RunGame(*white, *black, settings,communicator)]++

		PrintResults(results)
	}
}
