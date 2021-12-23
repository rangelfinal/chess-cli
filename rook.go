package main

const BlackRookRune = "♜"
const WhiteRookRune = "♖"

func getAvaiableRookMoves(board *Board, row int, col int) [][2]int {
	index := getIndexFromCoords(row, col)
	piece := board[index]

	avaiableMoves := make([][2]int, 0, 14)

	// We use multiple fors to enable short-circuiting with ease

	// Right
	if col < 7 {
		for i := col + 1; i < 7; i++ {
			if checkRookMove(board, piece.Color, row, col, row, i) {
				avaiableMoves = append(avaiableMoves, [2]int{row, i})
			} else {
				break
			}
		}
	}

	// Left
	if col > 0 {
		for i := col - 1; i > 0; i-- {
			if checkRookMove(board, piece.Color, row, col, row, i) {
				avaiableMoves = append(avaiableMoves, [2]int{row, i})
			} else {
				break
			}
		}
	}

	// Up
	if row < 7 {
		for i := row + 1; i < 7; i++ {
			if checkRookMove(board, piece.Color, row, col, i, col) {
				avaiableMoves = append(avaiableMoves, [2]int{i, col})
			} else {
				break
			}
		}
	}

	// Down
	if row > 0 {
		for i := row - 1; i > 0; i-- {
			if checkRookMove(board, piece.Color, row, col, i, col) {
				avaiableMoves = append(avaiableMoves, [2]int{i, col})
			} else {
				break
			}
		}
	}

	return avaiableMoves
}

func checkRookMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement != 0 && colMovement != 0 {
		return false
	}
	if hasAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	if rowMovement > 0 { // Horizontal
		i := startRow

		for i != endRow {
			if hasAnyPieceOnPosition(board, i, startCol) {
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
			if hasAnyPieceOnPosition(board, i, startCol) {
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
