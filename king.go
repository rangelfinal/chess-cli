package main

const BlackKingRune = "♚"
const WhiteKingRune = "♔"

func getAvaiableKingMoves(board *Board, row int, col int) [][2]int {
	index := getIndexFromCoords(row, col)
	piece := board[index]

	possibleMoves := [8][2]int{
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
		{row, col - 1} /*-----King------*/, {row, col + 1},
		{row - 1, col - 1}, {row - 1, col}, {row + 1, col + 1},
	}
	avaiableMoves := make([][2]int, 0, 8)

	for _, move := range possibleMoves {
		if checkKingMove(board, piece.Color, row, col, move[0], move[1]) {
			avaiableMoves = append(avaiableMoves, move)
		}
	}

	return avaiableMoves
}

func checkKingMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement > 1 || rowMovement < -1 || colMovement > 1 || colMovement < -1 {
		return false
	}

	if hasAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	return true
}
