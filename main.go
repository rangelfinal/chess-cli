package main

import (
	"fmt"
)

func main() {
	board := importBoard(startingPosition)

	renderBoard(&board, true)

	/*var input string
	fmt.Scanf("%s", &input)
	startCol, startRow, endCol, endRow := parseExpandedNotationToCoords(input)

	doMove(&board, White, startCol, startRow, endCol, endRow)*/

	renderBoard(&board, false)

	println(exportBoard(&board))

	startCol, startRow, endCol, endRow := parseExpandedNotationToCoords("E2E4")

	_, error := doMove(&board, true, startCol, startRow, endCol, endRow)

	if error != nil {
		fmt.Printf("%s\n", error.Error())
	}

	renderBoard(&board, true)

	println(exportBoard(&board))

	board = importBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")

	renderBoard(&board, true)

	println(exportBoard(&board))
}
