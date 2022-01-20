package main

func getAvaiableKingMoves(board *Board, col int, row int) [][2]int {
	index := getIndexFromCoords(col, row)
	piece := board[index]

	possibleMoves := [8][2]int{
		{col - 1, row + 1}, {col, row + 1}, {col + 1, row + 1},
		{col, row - 1} /*-----King------*/, {col, row + 1},
		{col - 1, row - 1}, {row - 1, col}, {col + 1, row + 1},
	}
	avaiableMoves := make([][2]int, 0, 8)

	for _, move := range possibleMoves {
		if checkKingMove(board, piece.Color, col, row, move[0], move[1]) {
			avaiableMoves = append(avaiableMoves, move)
		}
	}

	return avaiableMoves
}

func checkKingMove(board *Board, playerColor PieceColor, startCol int, startRow int, endCol int, endRow int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// TODO: Castle

	// Impossible moves
	if rowMovement > 1 || rowMovement < -1 || colMovement > 1 || colMovement < -1 {
		return false
	}

	if hasAlliedPieceOnPosition(board, playerColor, endCol, endRow) {
		return false
	}

	return true
}
