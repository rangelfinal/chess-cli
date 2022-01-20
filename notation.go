package main

import "fmt"

func getFigurineAlgebraicFromCoords(board *Board, startCol byte, startRow byte, endCol byte, endRow byte) string {
	index := getIndexFromCoords(startCol, startRow)
	piece := board.placements[index]

	return fmt.Sprintf("%s%s%d", piece, string('A'+endCol), endRow)
}

func getExpandedNotationFromCoords(board *Board, startCol byte, startRow byte, endCol byte, endRow byte) string {
	index := getIndexFromCoords(startCol, startRow)
	piece := board.placements[index]

	return fmt.Sprintf("%s%s%d%s%d", piece, string('A'+startCol), startRow, string('A'+endCol), endRow)
}

func parseExpandedNotationToCoords(not string) (byte, byte, byte, byte) {
	var startColS, endColS string
	var startCol, startRow, endCol, endRow byte

	fmt.Sscanf(not, "%1s%d%1s%d", &startColS, &startRow, &endColS, &endRow)

	startCol = byte(startColS[0] - 'A')
	endCol = byte(endColS[0] - 'A')
	startRow--
	endRow--

	return startCol, startRow, endCol, endRow
}
