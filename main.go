package main

import (
	"fmt"
	"strings"
)

func getFigurineAlgebraicFromCoords(board *Board, startCol int, startRow int, endCol int, endRow int) string {
	index := getIndexFromCoords(startCol, startRow)
	piece := board[index]

	return fmt.Sprintf("%s%s%d", piece, string('A'+endCol), endRow)
}

func getExpandedNotationFromCoords(board *Board, startCol int, startRow int, endCol int, endRow int) string {
	index := getIndexFromCoords(startCol, startRow)
	piece := board[index]

	return fmt.Sprintf("%s%s%d%s%d", piece, string('A'+startCol), startRow, string('A'+endCol), endRow)
}

func parseExpandedNotationToCoords(not string) (int, int, int, int) {
	var startColS, endColS string
	var startCol, startRow, endCol, endRow int

	fmt.Sscanf(not, "%1s%d%1s%d", &startColS, &startRow, &endColS, &endRow)
	fmt.Printf("%s - %d - %s - %d", startColS, startRow, endColS, endRow)

	startCol = int(startColS[0] - 'A')
	endCol = int(endColS[0] - 'A')
	startRow--
	endRow--
	fmt.Printf("%d - %d - %d - %d\n", startCol, startRow, endCol, endRow)

	return startCol, startRow, endCol, endRow
}

func exportBoard(board *Board) string {
	var output string

	for index, piece := range board {
		if piece != &Empty {
			output += fmt.Sprintf("%2d%1d%s;", index, piece.Color, rankCharMap[piece.Type])
		}
	}

	return output
}

func importBoard(s string) Board {
	var board Board
	pieces := strings.Split(s, ";")

	for _, piece := range pieces {
		var index int
		var color PieceColor
		var rank string

		fmt.Sscanf(piece, "%2d%1d%s", &index, &color, &rank)

		board[index] = &Piece{charRankMap[rank], color}
	}

	return board
}

func main() {
	renderBoard(&board, White)

	/*var input string
	fmt.Scanf("%s", &input)
	startCol, startRow, endCol, endRow := parseExpandedNotationToCoords(input)

	doMove(&board, White, startCol, startRow, endCol, endRow)*/

	renderBoard(&board, Black)

	startCol, startRow, endCol, endRow := parseExpandedNotationToCoords("E2E4")

	_, error := doMove(&board, White, startCol, startRow, endCol, endRow)

	renderBoard(&board, White)

	if error != nil {
		fmt.Printf("%s\n", error.Error())
	}
}
