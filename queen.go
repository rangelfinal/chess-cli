package main

const BlackQueenRune = "♛"
const WhiteQueenRune = "♕"

// A queen is a bishop on top of a tower

func getAvaiableQueenMoves(board *Board, row int, col int) [][2]int {
	avaiableMoves := make([][2]int, 0, 28)

	avaiableMoves = append(avaiableMoves, getAvaiableRookMoves(board, row, col)...)
	avaiableMoves = append(avaiableMoves, getAvaiableBishopMoves(board, row, col)...)

	return avaiableMoves
}

func checkQueenMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	return checkRookMove(board, playerColor, startRow, startCol, endRow, endCol) || checkBishopMove(board, playerColor, startRow, startCol, endRow, endCol)
}
