package communicator

import (
	"github.com/greggoryhz/cfw/board"
	"io"
	"strings"
	"fmt"
)

type TextCommunicator struct{}

func (t TextCommunicator) GenerateRequest(isWhite bool,brd board.Board) io.Reader {
	var stdin io.Reader
	if isWhite {
		stdin = strings.NewReader("white " + brd.String() + "\n")
	} else {
		stdin = strings.NewReader("black " + brd.String() + "\n")
	}
	return stdin
}

func (t TextCommunicator) ProcessResponse(out io.Reader) board.Move {
	var move board.Move

	fmt.Fscanf(out, "%d %d", &move.Src, &move.Dest)

	return move
}
