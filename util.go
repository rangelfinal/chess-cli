package main

/*
	      0    1    2    3    4    5    6    7
	   +----+----+----+----+----+----+----+----+
	0  |  0 |  1 |  2 |  3 |  4 |  5 |  6 |  7 |
		 +----+----+----+----+----+----+----+----+
	1  |  8 |  9 | 10 | 11 | 12 | 13 | 14 | 15 |
		 +----+----+----+----+----+----+----+----+
	2  | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 |
		 +----+----+----+----+----+----+----+----+
	3  | 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 |
		 +----+----+----+----+----+----+----+----+
	4  | 32 | 33 | 34 | 35 | 36 | 37 | 38 | 39 |
		 +----+----+----+----+----+----+----+----+
	5  | 40 | 41 | 42 | 43 | 44 | 45 | 46 | 47 |
		 +----+----+----+----+----+----+----+----+
	6  | 48 | 49 | 50 | 51 | 52 | 53 | 54 | 55 |
		 +----+----+----+----+----+----+----+----+
	7  | 56 | 57 | 58 | 59 | 60 | 61 | 62 | 63 |
		 +----+----+----+----+----+----+----+----+
*/

func getIndexFromCoords(row int, col int) int {
	return row*8 + col
}

func getCoordsFromIndex(index int) (int, int) {
	return index / 8, index % 8
}

func isPieceOnPosition(board *Board, piece Piece, row int, col int) bool {
	index := getIndexFromCoords(row, col)

	if board[index] == &Empty {
		return false
	}

	return board[index].Type == piece.Type && board[index].Color == piece.Color
}

func hasAnyPieceOnPosition(board *Board, row int, col int) bool {
	index := getIndexFromCoords(row, col)

	return board[index] != &Empty
}

func hasOpposingPieceOnPosition(board *Board, playerColor PieceColor, row int, col int) bool {
	index := getIndexFromCoords(row, col)

	return board[index] != &Empty && board[index].Color != playerColor
}

func hasAlliedPieceOnPosition(board *Board, playerColor PieceColor, row int, col int) bool {
	index := getIndexFromCoords(row, col)

	return board[index] != &Empty && board[index].Color == playerColor
}
