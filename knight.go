package main

func getAvaiableKnightMoves(board *Board, col int, row int) [][2]int {
	index := getIndexFromCoords(col, row)
	piece := board[index]

	possibleMoves := [8][2]int{
		{col + 1, row + 2}, {col - 1, row + 2},
		{col + 2, row + 1}, {col - 2, row + 1},
		{col + 2, row - 1}, {col - 2, row - 1},
		{col + 1, row - 2}, {col - 1, row - 2},
	}
	avaiableMoves := make([][2]int, 0, 8)

	for _, move := range possibleMoves {
		if checkKnightMove(board, piece.Color, col, row, move[0], move[1]) {
			avaiableMoves = append(avaiableMoves, move)
		}
	}

	return avaiableMoves
}

func checkKnightMove(board *Board, playerColor PieceColor, startCol int, startRow int, endCol int, endRow int) bool {
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
	if hasAlliedPieceOnPosition(board, playerColor, endCol, endRow) {
		return false
	}

	return true
}
