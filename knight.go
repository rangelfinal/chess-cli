package main

func getAvailableKnightMoves(board *Board, col byte, row byte) [][2]byte {
	index := getIndexFromCoords(col, row)
	piece := board.placements[index]

	possibleMoves := [8][2]byte{
		{col + 1, row + 2}, {col - 1, row + 2},
		{col + 2, row + 1}, {col - 2, row + 1},
		{col + 2, row - 1}, {col - 2, row - 1},
		{col + 1, row - 2}, {col - 1, row - 2},
	}
	availableMoves := make([][2]byte, 0, 8)

	for _, move := range possibleMoves {
		if checkKnightMove(board, piece.isWhite, col, row, move[0], move[1]) {
			availableMoves = append(availableMoves, move)
		}
	}

	return availableMoves
}

func checkKnightMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	rowMovement := int(endRow) - int(startRow)
	colMovement := int(endCol) - int(startCol)

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
	if hasAlliedPieceOnPosition(board, playerIsWhite, endCol, endRow) {
		return false
	}

	return true
}
