package marshaller

import (
	"github.com/greggoryhz/cfw/board"
	"io"
)

type Marshaller interface {
	GenerateRequest(isWhite bool,b board.Board) io.Reader
	ProcessResponse(out io.Reader) board.Move
}
