package main

func getAvailableBishopMoves(board *Board, col byte, row byte) [][2]byte {
	index := getIndexFromCoords(col, row)
	piece := board.placements[index]

	availableMoves := make([][2]byte, 0, 14)

	// We use multiple fors to enable short-circuiting with ease

	// Up-Right
	if col < 7 && row < 7 {
		for i := col + 1; i < 7; i++ {
			if checkBishopMove(board, piece.isWhite, col, row, i, i) {
				availableMoves = append(availableMoves, [2]byte{i, i})
			} else {
				break
			}
		}
	}

	// Up-Left
	if col > 0 && row < 7 {
		j := col
		for i := row + 1; i < 7; i++ {
			j--
			if checkBishopMove(board, piece.isWhite, col, row, j, i) {
				availableMoves = append(availableMoves, [2]byte{j, i})
			} else {
				break
			}
		}
	}

	// Down-Right
	if row < 7 {
		j := col
		for i := row - 1; i > 7; i-- {
			j++
			if checkBishopMove(board, piece.isWhite, col, row, j, i) {
				availableMoves = append(availableMoves, [2]byte{j, i})
			} else {
				break
			}
		}
	}

	// Down-Left
	if col > 0 && row > 0 {
		for i := row - 1; i > 0; i-- {
			if checkBishopMove(board, piece.isWhite, col, row, i, i) {
				availableMoves = append(availableMoves, [2]byte{i, i})
			} else {
				break
			}
		}
	}

	return availableMoves
}

func checkBishopMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	rowMovement := int(endRow) - int(startRow)
	colMovement := int(endCol) - int(startCol)

	// Impossible moves
	if rowMovement != colMovement && rowMovement != (colMovement*-1) {
		return false
	}
	if hasAlliedPieceOnPosition(board, playerIsWhite, endCol, endRow) {
		return false
	}

	i := startRow
	j := startCol

	for i != endRow && j != endCol {
		if hasAnyPieceOnPosition(board, j, i) {
			return false
		}

		if startRow < endRow {
			i++
		} else {
			i--
		}
		if startCol < endCol {
			j++
		} else {
			j--
		}
	}

	return true
}
