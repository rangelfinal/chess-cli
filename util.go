package main

/*
	      0    1    2    3    4    5    6    7
	   +----+----+----+----+----+----+----+----+
	7  |  0 |  1 |  2 |  3 |  4 |  5 |  6 |  7 |
		 +----+----+----+----+----+----+----+----+
	6  |  8 |  9 | 10 | 11 | 12 | 13 | 14 | 15 |
		 +----+----+----+----+----+----+----+----+
	5  | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 |
		 +----+----+----+----+----+----+----+----+
	4  | 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 |
		 +----+----+----+----+----+----+----+----+
	3  | 32 | 33 | 34 | 35 | 36 | 37 | 38 | 39 |
		 +----+----+----+----+----+----+----+----+
	2  | 40 | 41 | 42 | 43 | 44 | 45 | 46 | 47 |
		 +----+----+----+----+----+----+----+----+
	1  | 48 | 49 | 50 | 51 | 52 | 53 | 54 | 55 |
		 +----+----+----+----+----+----+----+----+
	0  | 56 | 57 | 58 | 59 | 60 | 61 | 62 | 63 |
		 +----+----+----+----+----+----+----+----+
*/

func getIndexFromCoords(col int, row int) int {
	return (7-row)*8 + col
}

func getCoordsFromIndex(index int) (int, int) {
	return index % 8, 7 - index/8
}

func isPieceOnPosition(board *Board, piece Piece, col int, row int) bool {
	index := getIndexFromCoords(col, row)

	if board[index] == &Empty {
		return false
	}

	return board[index].Type == piece.Type && board[index].Color == piece.Color
}

func hasAnyPieceOnPosition(board *Board, col int, row int) bool {
	index := getIndexFromCoords(col, row)

	return board[index] != &Empty
}

func hasOpposingPieceOnPosition(board *Board, playerColor PieceColor, col int, row int) bool {
	index := getIndexFromCoords(col, row)

	return board[index] != &Empty && board[index].Color != playerColor
}

func hasAlliedPieceOnPosition(board *Board, playerColor PieceColor, col int, row int) bool {
	index := getIndexFromCoords(col, row)

	return board[index] != &Empty && board[index].Color == playerColor
}
