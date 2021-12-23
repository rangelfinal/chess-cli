package main

import (
	"errors"
	"fmt"
)

type PieceType int

const (
	Pawn PieceType = iota
	Bishop
	Knight
	Rook
	King
	Queen
)

type PieceColor int

const (
	White PieceColor = iota
	Black
)

type Piece struct {
	Type  PieceType
	Color PieceColor
}

func (p Piece) String() string {
	if p.Color == Black {

		switch p.Type {
		case Pawn:
			return "♟︎"
		case Bishop:
			return "♝"
		case Knight:
			return "♞"
		case Rook:
			return "♜"
		case King:
			return "♚"
		case Queen:
			return "♛"
		}
	} else {

		switch p.Type {
		case Pawn:
			return "♙"
		case Bishop:
			return "♗"
		case Knight:
			return "♘"
		case Rook:
			return "♖"
		case King:
			return "♔"
		case Queen:
			return "♕"
		}
	}

	return "?"
}

var Empty = Piece{}

var (
	WhiteLeftRook    = Piece{Rook, White}
	WhiteLeftKnight  = Piece{Knight, White}
	WhiteLeftBishop  = Piece{Bishop, White}
	WhiteQueen       = Piece{Queen, White}
	WhiteKing        = Piece{King, White}
	WhiteRightBishop = Piece{Bishop, White}
	WhiteRightKnight = Piece{Knight, White}
	WhiteRightRook   = Piece{Rook, White}

	WhitePawn1 = Piece{Pawn, White}
	WhitePawn2 = Piece{Pawn, White}
	WhitePawn3 = Piece{Pawn, White}
	WhitePawn4 = Piece{Pawn, White}
	WhitePawn5 = Piece{Pawn, White}
	WhitePawn6 = Piece{Pawn, White}
	WhitePawn7 = Piece{Pawn, White}
	WhitePawn8 = Piece{Pawn, White}

	BlackLeftRook    = Piece{Rook, Black}
	BlackLeftKnight  = Piece{Knight, Black}
	BlackLeftBishop  = Piece{Bishop, Black}
	BlackQueen       = Piece{Queen, Black}
	BlackKing        = Piece{King, Black}
	BlackRightBishop = Piece{Bishop, Black}
	BlackRightKnight = Piece{Knight, Black}
	BlackRightRook   = Piece{Rook, Black}

	BlackPawn1 = Piece{Pawn, Black}
	BlackPawn2 = Piece{Pawn, Black}
	BlackPawn3 = Piece{Pawn, Black}
	BlackPawn4 = Piece{Pawn, Black}
	BlackPawn5 = Piece{Pawn, Black}
	BlackPawn6 = Piece{Pawn, Black}
	BlackPawn7 = Piece{Pawn, Black}
	BlackPawn8 = Piece{Pawn, Black}
)

type Board [64]*Piece

var board = Board{
	&WhiteLeftRook, &WhiteLeftKnight, &WhiteLeftBishop, &WhiteKing, &WhiteQueen, &WhiteRightBishop, &WhiteRightKnight, &WhiteRightRook,
	&WhitePawn1, &WhitePawn2, &WhitePawn3, &WhitePawn4, &WhitePawn5, &WhitePawn6, &WhitePawn7, &WhitePawn8,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&BlackPawn1, &BlackPawn2, &BlackPawn3, &BlackPawn4, &BlackPawn5, &BlackPawn6, &BlackPawn7, &BlackPawn8,
	&BlackLeftRook, &BlackLeftKnight, &BlackLeftBishop, &BlackKing, &BlackQueen, &BlackRightBishop, &BlackRightKnight, &BlackRightRook,
}

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

func rowColToIndex(row int, col int) int {
	return row*8 + col
}

func indexToRowCol(index int) (int, int) {
	return index / 8, index % 8
}

func isPieceOnPosition(board *Board, piece Piece, row int, col int) bool {
	index := rowColToIndex(row, col)

	if board[index] == &Empty {
		return false
	}

	return board[index].Type == piece.Type && board[index].Color == piece.Color
}

func hasAnyPieceOnPosition(board *Board, row int, col int) bool {
	index := rowColToIndex(row, col)

	return board[index] != &Empty
}

func hasOpposingPieceOnPosition(board *Board, playerColor PieceColor, row int, col int) bool {
	index := rowColToIndex(row, col)

	return board[index] != &Empty && board[index].Color != playerColor
}

func haAlliedPieceOnPosition(board *Board, playerColor PieceColor, row int, col int) bool {
	index := rowColToIndex(row, col)

	return board[index] != &Empty && board[index].Color == playerColor
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
				if startRow != 1 {
					return false
				}
				if hasAnyPieceOnPosition(board, startRow+1, startCol) {
					return false
				}
			} else {
				if startRow != 6 {
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
	if rowMovement == 2 {
		return false
	}

	// TODO: En Passant
	if !hasOpposingPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	return true
}

func checkRookMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement != 0 && colMovement != 0 {
		return false
	}
	if haAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
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

func checkBishopMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement != colMovement && rowMovement != (colMovement*-1) {
		return false
	}
	if haAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	i := startRow
	j := startCol

	for i != endRow && j != endCol {
		if hasAnyPieceOnPosition(board, i, j) {
			return false
		}

		if startRow < endRow {
			i++
		} else {
			i--
		}
		if startCol < endCol {
			j++
		} else {
			j--
		}
	}

	return true
}

func checkKnightMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	if rowMovement < 0 {
		rowMovement *= -1
	}
	if colMovement < 0 {
		colMovement *= -1
	}

	// Impossible moves
	if !((rowMovement == 2 && colMovement == 1) || (colMovement == 2 && rowMovement == 1)) {
		return false
	}
	if haAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	return true
}

func checkQueenMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	return checkRookMove(board, playerColor, startRow, startCol, endRow, endCol) || checkBishopMove(board, playerColor, startRow, startCol, endRow, endCol)
}

func checkKingMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	rowMovement := endRow - startRow
	colMovement := endCol - startCol

	// Impossible moves
	if rowMovement > 1 || rowMovement < -1 || colMovement > 1 || colMovement < -1 {
		return false
	}

	if haAlliedPieceOnPosition(board, playerColor, endRow, endCol) {
		return false
	}

	return true
}

func checkMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	pieceIndex := rowColToIndex(startRow, startCol)
	piece := board[pieceIndex]

	if piece == &Empty {
		return false
	}

	switch piece.Type {
	case Pawn:
		return checkPawnMove(board, playerColor, startRow, startCol, endRow, endCol)
	case Rook:
		return checkRookMove(board, playerColor, startRow, startCol, endRow, endCol)
	case Knight:
		return checkKnightMove(board, playerColor, startRow, startCol, endRow, endCol)
	case Bishop:
		return checkBishopMove(board, playerColor, startRow, startCol, endRow, endCol)
	case Queen:
		return checkQueenMove(board, playerColor, startRow, startCol, endRow, endCol)
	case King:
		return checkKingMove(board, playerColor, startRow, startCol, endRow, endCol)
	}

	return false
}

func checkCheck(board *Board, playerColor PieceColor) bool {
	var kingRow, kingCol int

	for index, piece := range board {
		if (playerColor == White && piece == &WhiteKing) || (playerColor == Black && piece == &BlackKing) {
			kingRow, kingCol = indexToRowCol(index)
		}
	}

	for index, piece := range board {
		if piece != &Empty && piece.Color != playerColor {
			startRow, startCol := indexToRowCol(index)

			if checkMove(board, piece.Color, startRow, startCol, kingRow, kingCol) {
				fmt.Printf("Win: %v - %v - %v - %v - %v - %v - %v\n", playerColor, startRow, startCol, kingRow, kingCol, piece.Type, piece.Color)
				return true
			}
		}
	}

	return false
}

func doMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) (Board, error) {
	startIndex := rowColToIndex(startRow, startCol)
	startPiece := board[startIndex]

	if startPiece.Color != playerColor {
		return *board, errors.New(("Wrong turn"))
	}

	if !checkMove(board, playerColor, startRow, startCol, endRow, endCol) {
		return *board, errors.New("Invalid move")
	}

	endIndex := rowColToIndex(endRow, endCol)

	backup := board[endIndex]
	board[endIndex] = board[startIndex]
	board[startIndex] = &Empty

	if checkCheck(board, playerColor) {
		board[startIndex] = board[endIndex]
		board[endIndex] = backup

		return *board, errors.New("Would lose the game")
	}

	return *board, nil
}

/*
      | A | B | C | D | E | F | G | H |
   1  | ♜ | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ | ♜ |
   2  | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ |
   3  |   |   |   |   |   |   |   |   |
   4  |   |   |   |   |   |   |   |   |
   5  |   |   |   |   |   |   |   |   |
   6  |   |   |   |   |   |   |   |   |
   7  | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ |
   8  | ♖ | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ | ♖ |
*/
func renderBoard(board *Board, playerColor PieceColor) {
	output := "           |---+---+---+---+---+---+---+---|\n"
	if playerColor == White {
		output += "           | A | B | C | D | E | F | G | H |\n"
	} else {
		output += "           | H | G | F | E | D | C | B | A |\n"
	}
	output += "           |---+---+---+---+---+---+---+---|"

	for index := 0; index < 64; index++ {
		var piece *Piece
		if playerColor == Black {
			piece = board[index]
		} else {
			piece = board[63-index]
		}

		if index%8 == 0 {
			var colNumber int
			if playerColor == Black {
				colNumber = 8 - (index / 8)
			} else {
				colNumber = (index / 8) + 1
			}
			output += fmt.Sprintf("\n	%d  |", colNumber)
		}

		if piece == &Empty {
			output += "   |"
		} else {
			output += fmt.Sprintf(" %s |", piece)
		}
	}

	output += "\n           |---+---+---+---+---+---+---+---|\n\n"

	fmt.Print(output)
}

func main() {
	renderBoard(&board, White)
	renderBoard(&board, Black)

	_, error := doMove(&board, White, 1, 3, 3, 3)
	renderBoard(&board, White)
	if error != nil {
		fmt.Printf("%s\n", error.Error())
	}

	_, error2 := doMove(&board, Black, 6, 3, 4, 3)
	renderBoard(&board, Black)
	if error2 != nil {
		fmt.Printf("%s\n", error2.Error())
	}

	_, error3 := doMove(&board, White, 1, 2, 3, 2)
	renderBoard(&board, White)
	if error3 != nil {
		fmt.Printf("%s\n", error3.Error())
	}

	_, error4 := doMove(&board, Black, 6, 2, 4, 2)
	renderBoard(&board, Black)
	if error4 != nil {
		fmt.Printf("%s\n\n", error4.Error())
	}
}
