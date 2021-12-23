package main

const BlackKnightRune = "♞"
const WhiteKnightRune = "♘"

func getAvaiableKnightMoves(board *Board, row int, col int) [][2]int {
	index := getIndexFromCoords(row, col)
	piece := board[index]

	possibleMoves := [8][2]int{
		{row + 2, col + 1}, {row + 2, col - 1},
		{row + 1, col + 2}, {row + 1, col - 2},
		{row - 1, col + 2}, {row - 1, col - 2},
		{row - 2, col + 1}, {row - 2, col - 1},
	}
	avaiableMoves := make([][2]int, 0, 8)

	for _, move := range possibleMoves {
		if checkKnightMove(board, piece.Color, row, col, move[0], move[1]) {
			avaiableMoves = append(avaiableMoves, move)
		}
	}

	return avaiableMoves
}

func checkKnightMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	if rowMovement < 0 {
		rowMovement *= -1
	}
	if colMovement < 0 {
		colMovement *= -1
	}

	// Impossible moves
	if !((rowMovement == 2 && colMovement == 1) || (colMovement == 2 && rowMovement == 1)) {
		return false
	}
	if hasAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	return true
}
