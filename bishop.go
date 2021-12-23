package main

const BlackBishopRune = "♝"
const WhiteBishopRune = "♗"

func getAvaiableBishopMoves(board *Board, row int, col int) [][2]int {
	index := getIndexFromCoords(row, col)
	piece := board[index]

	avaiableMoves := make([][2]int, 0, 14)

	// We use multiple fors to enable short-circuiting with ease

	// Up-Right
	if col < 7 && row < 7 {
		for i := col + 1; i < 7; i++ {
			if checkBishopMove(board, piece.Color, row, col, i, i) {
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
			if checkBishopMove(board, piece.Color, row, col, i, j) {
				avaiableMoves = append(avaiableMoves, [2]int{i, j})
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
			if checkBishopMove(board, piece.Color, row, col, i, j) {
				avaiableMoves = append(avaiableMoves, [2]int{i, j})
			} else {
				break
			}
		}
	}

	// Down-Left
	if col > 0 && row > 0 {
		for i := row - 1; i > 0; i-- {
			if checkBishopMove(board, piece.Color, row, col, i, i) {
				avaiableMoves = append(avaiableMoves, [2]int{i, i})
			} else {
				break
			}
		}
	}

	return avaiableMoves
}

func checkBishopMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement != colMovement && rowMovement != (colMovement*-1) {
		return false
	}
	if hasAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	i := startRow
	j := startCol

	for i != endRow && j != endCol {
		if hasAnyPieceOnPosition(board, i, j) {
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
