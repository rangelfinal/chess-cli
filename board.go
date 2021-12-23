package main

import (
	"errors"
	"fmt"
)

// Will be used to represent spaces on the board without pieces
var Empty = Piece{}

var (
	WhiteQueenRook   = Piece{Rook, White}
	WhiteQueenKnight = Piece{Knight, White}
	WhiteQueenBishop = Piece{Bishop, White}
	WhiteQueen       = Piece{Queen, White}
	WhiteKing        = Piece{King, White}
	WhiteKingBishop  = Piece{Bishop, White}
	WhiteKingKnight  = Piece{Knight, White}
	WhiteKingRook    = Piece{Rook, White}

	WhitePawn1 = Piece{Pawn, White}
	WhitePawn2 = Piece{Pawn, White}
	WhitePawn3 = Piece{Pawn, White}
	WhitePawn4 = Piece{Pawn, White}
	WhitePawn5 = Piece{Pawn, White}
	WhitePawn6 = Piece{Pawn, White}
	WhitePawn7 = Piece{Pawn, White}
	WhitePawn8 = Piece{Pawn, White}

	BlackQueenRook   = Piece{Rook, Black}
	BlackQueenKnight = Piece{Knight, Black}
	BlackQueenBishop = Piece{Bishop, Black}
	BlackQueen       = Piece{Queen, Black}
	BlackKing        = Piece{King, Black}
	BlackKingBishop  = Piece{Bishop, Black}
	BlackKingKnight  = Piece{Knight, Black}
	BlackKingRook    = Piece{Rook, Black}

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
	&WhiteQueenRook, &WhiteQueenKnight, &WhiteQueenBishop, &WhiteKing, &WhiteQueen, &WhiteKingBishop, &WhiteKingKnight, &WhiteKingRook,
	&WhitePawn1, &WhitePawn2, &WhitePawn3, &WhitePawn4, &WhitePawn5, &WhitePawn6, &WhitePawn7, &WhitePawn8,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&BlackPawn1, &BlackPawn2, &BlackPawn3, &BlackPawn4, &BlackPawn5, &BlackPawn6, &BlackPawn7, &BlackPawn8,
	&BlackQueenRook, &BlackQueenKnight, &BlackQueenBishop, &BlackKing, &BlackQueen, &BlackKingBishop, &BlackKingKnight, &BlackKingRook,
}

// Get all avaiable moves from all pieces of a player
func getAvaiableMoves(board *Board, playerColor PieceColor) [][4]int {
	avaiableMoves := make([][4]int, 0, 64)

	for index, piece := range board {
		if piece != &Empty && piece.Color == playerColor {
			row, col := getCoordsFromIndex(index)

			var moves [][2]int
			if piece.Type == Pawn {
				moves = getAvaiablePawnMoves(board, row, col)
			}
			if piece.Type == Rook {
				moves = getAvaiableRookMoves(board, row, col)
			}
			if piece.Type == Knight {
				moves = getAvaiableKnightMoves(board, row, col)
			}
			if piece.Type == Bishop {
				moves = getAvaiableBishopMoves(board, row, col)
			}
			if piece.Type == Queen {
				moves = getAvaiableQueenMoves(board, row, col)
			}
			if piece.Type == King {
				moves = getAvaiableKingMoves(board, row, col)
			}

			for _, move := range moves {
				avaiableMoves = append(avaiableMoves, [4]int{row, col, move[0], move[1]})
			}
		}
	}

	return avaiableMoves
}

// Ensures a move is valid for current board
func checkMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) bool {
	pieceIndex := getIndexFromCoords(startRow, startCol)
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

// Verifies if the player king is in check
func checkCheck(board *Board, playerColor PieceColor) bool {
	var kingRow, kingCol int

	for index, piece := range board {
		if (playerColor == White && piece == &WhiteKing) || (playerColor == Black && piece == &BlackKing) {
			kingRow, kingCol = getCoordsFromIndex(index)
		}
	}

	for index, piece := range board {
		if piece != &Empty && piece.Color != playerColor {
			startRow, startCol := getCoordsFromIndex(index)

			if checkMove(board, piece.Color, startRow, startCol, kingRow, kingCol) {
				return true
			}
		}
	}

	return false
}

/* Verifies if the player king will be captured next turn.
Gets every avaiable move, create all possible future boards,
and checks if at least one is not in check anymore
*/
func checkMate(board *Board, playerColor PieceColor) bool {
	moves := getAvaiableMoves(board, playerColor)

	for _, move := range moves {
		testBoard := *board
		doMove(&testBoard, playerColor, move[0], move[1], move[2], move[3])
		if !checkCheck(&testBoard, playerColor) {
			return false
		}
	}

	return true
}

// Execute a movement if it's valid
func doMove(board *Board, playerColor PieceColor, startRow int, startCol int, endRow int, endCol int) (Board, error) {
	startIndex := getIndexFromCoords(startRow, startCol)
	startPiece := board[startIndex]

	if startPiece.Color != playerColor {
		return *board, errors.New(("Not your piece"))
	}

	if !checkMove(board, playerColor, startRow, startCol, endRow, endCol) {
		return *board, errors.New("Invalid move")
	}

	endIndex := getIndexFromCoords(endRow, endCol)

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
	Render the board to CLI in the format of:

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
