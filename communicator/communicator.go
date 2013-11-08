package communicator

import (
	"github.com/gregghz/cfw/board"
	"io"
)

type Communicator interface {
	GenerateRequest(isWhite bool,b board.Board) io.Reader
	ProcessResponse(out io.Reader) board.Move
}
