Protocol
========

High level
----------

Here's a high level view of a full turn cycle (white plays, then black
plays).

The framework sends the current board to the white AI. The white AI
send back a move and the board that is the result of making that
move. The same procedure is repeated for black. The framework should
validate each move and flag check, checkmate, and stalemate.

Currently the moves have no time constraints. This will be coming later.

Lower level
-----------

A board should be 64 piece symbols, each separated by a space and
ending with a newline. Each symbol is two characters. The first
character identifies the color of the piece and the second character
identies the type of the piece. 00 represents an empty square.

The board starts at the top left and moves by rows to the bottom
right.

So the starting board for a game is as follows:

    BR BH BB BQ BK BB BH BR BP BP BP BP BP BP BP BP 00 00 00 00 00 00
    00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
    00 00 00 00 WP WP WP WP WP WP WP WP WR WH WB WQ WK WB WH WR
    
At the start of a game, the starting board is sent (on stdin) to the
white AI. White then responds (on stdout) with a move
and the new board configuration. The response should be an integer
representing the index of the piece to move, a space, an integer
representing the target of the move, a space, and finally the new
board in the same format as shown above (but with the move made). This
move will be validated by the framework. If the move is invalid, it is
considered a loss. The AIs do not need to flag check, checkmate, or
stalemate because the framework is aware of these conditions and
handles the end of the game.

Running the framework
=====================

Run the framework with any of the following:

    go run main.go -black /path/to/ai -white /path/to/ai

    go build
    ./cfw -black /path/to/ai -white /path/to/ai

    go install
    # works if $GOPATH/bin is on your path
    cfw -black /path/to/ai -white /path/to/ai 

Other Notes
===========

I took a class using a chess framework that requires the use of .NET
(or mono). This framework makes it extremely difficult to run a lot of
tests against two different AIs and it was very tightly coupled with
the user interface. I decided that I wanted to design a similar
framework that effectively only defines a communication protocol and
doesn't enforce the use of any specific programming environment when
creating the AI. cfw should work on any platform that the Go compiler
targets (at least OS X, Linux, and Windows). However, my ability to
test on these various platforms is limited (particularly Windows). If
something doesn't work on any particular platform, it should be
reported as a bug.

A more long term goal is to allow an agent to compete over a network
connection instead of (as is currently the case) requiring the agent
to run locally.

Status
======

This project is very incomplete and currently does NOT satisfy most of
the protocol specified above.
