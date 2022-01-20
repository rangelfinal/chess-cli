package main

func getAvailableRookMoves(board *Board, col byte, row byte) [][2]byte {
	index := getIndexFromCoords(col, row)
	piece := board.placements[index]

	availableMoves := make([][2]byte, 0, 14)

	// We use multiple fors to enable short-circuiting with ease

	// Right
	if col < 7 {
		for i := col + 1; i < 7; i++ {
			if checkRookMove(board, piece.isWhite, col, row, i, row) {
				availableMoves = append(availableMoves, [2]byte{i, row})
			} else {
				break
			}
		}
	}

	// Left
	if col > 0 {
		for i := col - 1; i > 0; i-- {
			if checkRookMove(board, piece.isWhite, col, row, i, row) {
				availableMoves = append(availableMoves, [2]byte{i, row})
			} else {
				break
			}
		}
	}

	// Up
	if row < 7 {
		for i := row + 1; i < 7; i++ {
			if checkRookMove(board, piece.isWhite, col, row, col, i) {
				availableMoves = append(availableMoves, [2]byte{col, i})
			} else {
				break
			}
		}
	}

	// Down
	if row > 0 {
		for i := row - 1; i > 0; i-- {
			if checkRookMove(board, piece.isWhite, col, row, col, i) {
				availableMoves = append(availableMoves, [2]byte{col, i})
			} else {
				break
			}
		}
	}

	return availableMoves
}

func checkRookMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	rowMovement := int(endRow) - int(startRow)
	colMovement := int(endCol) - int(startCol)

	// Impossible moves
	if rowMovement != 0 && colMovement != 0 {
		return false
	}
	if hasAlliedPieceOnPosition(board, playerIsWhite, endCol, endRow) {
		return false
	}

	if rowMovement > 0 { // Horizontal
		i := startRow

		for i != endRow {
			if hasAnyPieceOnPosition(board, startCol, i) {
				return false
			}

			if startRow < endRow {
				i++
			} else {
				i--
			}
		}
	} else if colMovement > 0 { // Vertical
		i := startCol

		for i != endCol {
			if hasAnyPieceOnPosition(board, startCol, i) {
				return false
			}

			if startCol < endCol {
				i++
			} else {
				i--
			}
		}
	}

	return true
}
