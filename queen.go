package main

// A queen is a bishop on top of a tower

func getAvailableQueenMoves(board *Board, col byte, row byte) [][2]byte {
	availableMoves := make([][2]byte, 0, 28)

	availableMoves = append(availableMoves, getAvailableRookMoves(board, col, row)...)
	availableMoves = append(availableMoves, getAvailableBishopMoves(board, col, row)...)

	return availableMoves
}

func checkQueenMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	return checkRookMove(board, playerIsWhite, startCol, startRow, endCol, endRow) || checkBishopMove(board, playerIsWhite, startCol, startRow, endCol, endRow)
}
