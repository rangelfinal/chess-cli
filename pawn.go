package main

const BlackStartingPosition = 6
const WhiteStartingPosition = 1

func getAvailablePawnMoves(board *Board, col byte, row byte) [][2]byte {
	index := getIndexFromCoords(col, row)
	piece := board.placements[index]

	possibleMoves := [4][2]byte{{col, row + 1}, {col, row + 2}, {col, row + 1}, {col, row - 1}}
	availableMoves := make([][2]byte, 0, 4)

	for _, move := range possibleMoves {
		if checkPawnMove(board, piece.isWhite, col, row, move[0], move[1]) {
			availableMoves = append(availableMoves, move)
		}
	}

	return availableMoves
}

func checkPawnMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	rowMovement := int(endRow) - int(startRow)
	colMovement := int(endCol) - int(startCol)

	// Impossible moves
	if playerIsWhite && (rowMovement < 0 || rowMovement > 2) {
		return false
	}
	if !playerIsWhite && (rowMovement < -2 || rowMovement > 0) {
		return false
	}
	if colMovement > 1 || colMovement < -1 {
		return false
	}

	if colMovement == 0 { // Pawn movement
		if rowMovement == 2 || rowMovement == -2 { // Can move 2 spaces if in starting position
			if playerIsWhite {
				if startRow != WhiteStartingPosition {
					return false
				}
				if hasAnyPieceOnPosition(board, startCol, startRow+1) {
					return false
				}
			} else {
				if startRow != BlackStartingPosition {
					return false
				}
				if hasAnyPieceOnPosition(board, startCol, startRow-1) {
					return false
				}
			}
		}

		if hasAnyPieceOnPosition(board, endCol, endRow) {
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
	if !hasOpposingPieceOnPosition(board, playerIsWhite, endCol, endRow) {
		return false
	}

	return true
}
