package main

func getAvailableKingMoves(board *Board, col byte, row byte) [][2]byte {
	index := getIndexFromCoords(col, row)
	piece := board.placements[index]

	possibleMoves := [10][2]byte{
		{col - 1, row + 1}, {col, row + 1}, {col + 1, row + 1},
		{col, row - 1} /*-----King------*/, {col, row + 1},
		{col - 1, row - 1}, {row - 1, col}, {col + 1, row + 1},
		// Castling
		{col - 2, row}, {col + 2, row},
	}
	availableMoves := make([][2]byte, 0, 8)

	for _, move := range possibleMoves {
		if checkKingMove(board, piece.isWhite, col, row, move[0], move[1]) {
			availableMoves = append(availableMoves, move)
		}
	}

	return availableMoves
}

func checkKingMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	rowMovement := int(endRow) - int(startRow)
	colMovement := int(endCol) - int(startCol)

	// TODO: Castle

	// Impossible moves
	if rowMovement > 1 || rowMovement < -1 || colMovement > 2 || colMovement < -2 {
		return false
	}

	if playerIsWhite && colMovement == 2 && board.castling<<WhiteKingCastling <= 0 {
		return false
	}
	if playerIsWhite && colMovement == -2 && board.castling<<WhiteKingCastling <= 0 {
		return false
	}
	if !playerIsWhite && colMovement == 2 && board.castling<<BlackKingCastling <= 0 {
		return false
	}
	if !playerIsWhite && colMovement == -2 && board.castling<<BlackKingCastling <= 0 {
		return false
	}

	if hasAlliedPieceOnPosition(board, playerIsWhite, endCol, endRow) {
		return false
	}

	return true
}
