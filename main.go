package main

import (
	"fmt"
)

func getNotationFromCoords(board *Board, startRow int, startCol int, endRow int, endCol int) string {
	index := getIndexFromCoords(startRow, startCol)
	piece := board[index]

	return fmt.Sprintf("%s%s%d", piece, string('A'+endCol), endRow)
}

func getExpandedNotationFromCoords(board *Board, startRow int, startCol int, endRow int, endCol int) string {
	index := getIndexFromCoords(startRow, startCol)
	piece := board[index]

	return fmt.Sprintf("%s%s%d%s%d", piece, string('A'+startCol), startRow, string('A'+endCol), endRow)
}

func main() {
	renderBoard(&board, White)
	renderBoard(&board, Black)

	_, error := doMove(&board, White, 1, 3, 3, 3)
	renderBoard(&board, White)
	if error != nil {
		fmt.Printf("%s\n", error.Error())
	}

	_, error2 := doMove(&board, Black, 6, 3, 4, 3)
	renderBoard(&board, Black)
	if error2 != nil {
		fmt.Printf("%s\n", error2.Error())
	}

	_, error3 := doMove(&board, White, 1, 2, 3, 2)
	renderBoard(&board, White)
	if error3 != nil {
		fmt.Printf("%s\n", error3.Error())
	}

	_, error4 := doMove(&board, Black, 6, 2, 4, 2)
	renderBoard(&board, Black)
	if error4 != nil {
		fmt.Printf("%s\n\n", error4.Error())
	}
}
