package main

// A queen is a bishop on top of a tower

func getAvaiableQueenMoves(board *Board, col int, row int) [][2]int {
	avaiableMoves := make([][2]int, 0, 28)

	avaiableMoves = append(avaiableMoves, getAvaiableRookMoves(board, col, row)...)
	avaiableMoves = append(avaiableMoves, getAvaiableBishopMoves(board, col, row)...)

	return avaiableMoves
}

func checkQueenMove(board *Board, playerColor PieceColor, startCol int, startRow int, endCol int, endRow int) bool {
	return checkRookMove(board, playerColor, startCol, startRow, endCol, endRow) || checkBishopMove(board, playerColor, startCol, startRow, endCol, endRow)
}
