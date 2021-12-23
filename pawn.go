package main

const BlackPawnRune = "♟︎"
const WhitePawnRune = "♙"
const BlackStartingPosition = 6
const WhiteStartingPosition = 1

func getAvaiablePawnMoves(board *Board, row int, col int) [][2]int {
	index := getIndexFromCoords(row, col)
	piece := board[index]

	possibleMoves := [4][2]int{{row + 1, col}, {row + 2, col}, {row, col + 1}, {row, col - 1}}
	avaiableMoves := make([][2]int, 0, 4)

	for _, move := range possibleMoves {
		if checkPawnMove(board, piece.Color, row, col, move[0], move[1]) {
			avaiableMoves = append(avaiableMoves, move)
		}
	}

	return avaiableMoves
}

func checkPawnMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if playerColor == White && (rowMovement < 0 || rowMovement > 2) {
		return false
	}
	if playerColor == Black && (rowMovement < -2 || rowMovement > 0) {
		return false
	}
	if colMovement > 1 || colMovement < -1 {
		return false
	}

	if colMovement == 0 { // Pawn movement
		if rowMovement == 2 || rowMovement == -2 { // Can move 2 spaces if in starting position
			if playerColor == White {
				if startRow != WhiteStartingPosition {
					return false
				}
				if hasAnyPieceOnPosition(board, startRow+1, startCol) {
					return false
				}
			} else {
				if startRow != BlackStartingPosition {
					return false
				}
				if hasAnyPieceOnPosition(board, startRow-1, startCol) {
					return false
				}
			}
		}

		if hasAnyPieceOnPosition(board, endRow, endCol) {
			return false
		}

		return true
	}

	// Pawn capture

	// Impossible moves
	if rowMovement == 2 {
		return false
	}

	// TODO: En Passant
	if !hasOpposingPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	return true
}
