package main

func getAvaiableBishopMoves(board *Board, col int, row int) [][2]int {
	index := getIndexFromCoords(col, row)
	piece := board[index]

	avaiableMoves := make([][2]int, 0, 14)

	// We use multiple fors to enable short-circuiting with ease

	// Up-Right
	if col < 7 && row < 7 {
		for i := col + 1; i < 7; i++ {
			if checkBishopMove(board, piece.Color, col, row, i, i) {
				avaiableMoves = append(avaiableMoves, [2]int{i, i})
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
			if checkBishopMove(board, piece.Color, col, row, j, i) {
				avaiableMoves = append(avaiableMoves, [2]int{j, i})
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
			if checkBishopMove(board, piece.Color, col, row, j, i) {
				avaiableMoves = append(avaiableMoves, [2]int{j, i})
			} else {
				break
			}
		}
	}

	// Down-Left
	if col > 0 && row > 0 {
		for i := row - 1; i > 0; i-- {
			if checkBishopMove(board, piece.Color, col, row, i, i) {
				avaiableMoves = append(avaiableMoves, [2]int{i, i})
			} else {
				break
			}
		}
	}

	return avaiableMoves
}

func checkBishopMove(board *Board, playerColor PieceColor, startCol int, startRow int, endCol int, endRow int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement != colMovement && rowMovement != (colMovement*-1) {
		return false
	}
	if hasAlliedPieceOnPosition(board, playerColor, endCol, endRow) {
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
