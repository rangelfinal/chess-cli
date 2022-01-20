package main

import (
	"errors"
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
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
			return BlackPawnRune
		case Bishop:
			return BlackBishopRune
		case Knight:
			return BlackKnightRune
		case Rook:
			return BlackRookRune
		case King:
			return BlackKingRune
		case Queen:
			return BlackQueenRune
		}
	} else {

		switch p.Type {
		case Pawn:
			return WhitePawnRune
		case Bishop:
			return WhiteBishopRune
		case Knight:
			return WhiteKnightRune
		case Rook:
			return WhiteRookRune
		case King:
			return WhiteKingRune
		case Queen:
			return WhiteQueenRune
		}
	}

	return "?"
}

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
	&BlackQueenRook, &BlackQueenKnight, &BlackQueenBishop, &BlackKing, &BlackQueen, &BlackKingBishop, &BlackKingKnight, &BlackKingRook,
	&BlackPawn1, &BlackPawn2, &BlackPawn3, &BlackPawn4, &BlackPawn5, &BlackPawn6, &BlackPawn7, &BlackPawn8,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty, &Empty,
	&WhitePawn1, &WhitePawn2, &WhitePawn3, &WhitePawn4, &WhitePawn5, &WhitePawn6, &WhitePawn7, &WhitePawn8,
	&WhiteQueenRook, &WhiteQueenKnight, &WhiteQueenBishop, &WhiteKing, &WhiteQueen, &WhiteKingBishop, &WhiteKingKnight, &WhiteKingRook,
}

// Get all avaiable moves from all pieces of a player
func getAvaiableMoves(board *Board, playerColor PieceColor) [][4]int {
	avaiableMoves := make([][4]int, 0, 64)

	for index, piece := range board {
		if piece != &Empty && piece.Color == playerColor {
			col, row := getCoordsFromIndex(index)

			var moves [][2]int
			if piece.Type == Pawn {
				moves = getAvaiablePawnMoves(board, col, row)
			}
			if piece.Type == Rook {
				moves = getAvaiableRookMoves(board, col, row)
			}
			if piece.Type == Knight {
				moves = getAvaiableKnightMoves(board, col, row)
			}
			if piece.Type == Bishop {
				moves = getAvaiableBishopMoves(board, col, row)
			}
			if piece.Type == Queen {
				moves = getAvaiableQueenMoves(board, col, row)
			}
			if piece.Type == King {
				moves = getAvaiableKingMoves(board, col, row)
			}

			for _, move := range moves {
				avaiableMoves = append(avaiableMoves, [4]int{col, row, move[0], move[1]})
			}
		}
	}

	return avaiableMoves
}

// Ensures a move is valid for current board
func checkMove(board *Board, playerColor PieceColor, startCol int, startRow int, endCol int, endRow int) bool {
	pieceIndex := getIndexFromCoords(startCol, startRow)
	piece := board[pieceIndex]

	if piece == &Empty {
		return false
	}

	switch piece.Type {
	case Pawn:
		return checkPawnMove(board, playerColor, startCol, startRow, endCol, endRow)
	case Rook:
		return checkRookMove(board, playerColor, startCol, startRow, endCol, endRow)
	case Knight:
		return checkKnightMove(board, playerColor, startCol, startRow, endCol, endRow)
	case Bishop:
		return checkBishopMove(board, playerColor, startCol, startRow, endCol, endRow)
	case Queen:
		return checkQueenMove(board, playerColor, startCol, startRow, endCol, endRow)
	case King:
		return checkKingMove(board, playerColor, startCol, startRow, endCol, endRow)
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
			startCol, startRow := getCoordsFromIndex(index)

			if checkMove(board, piece.Color, startCol, startRow, kingRow, kingCol) {
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
func doMove(board *Board, playerColor PieceColor, startCol int, startRow int, endCol int, endRow int) (Board, error) {
	startIndex := getIndexFromCoords(startCol, startRow)
	startPiece := board[startIndex]

	if startPiece.Color != playerColor {
		return *board, errors.New("Not your piece")
	}

	if !checkMove(board, playerColor, startCol, startRow, endCol, endRow) {
		return *board, errors.New("Invalid move")
	}

	endIndex := getIndexFromCoords(endCol, endRow)

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

───┬───┬───┬───┬───┬───┬───┬───┬───┤
 8 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 7 │ ♟︎ │ ♟︎ │ ♟︎ │ ♟︎ │ ♟︎ │ ♟︎ │ ♟︎ │ ♟︎ │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 6 │   │   │   │   │   │   │   │   │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 5 │   │   │   │   │   │   │   │   │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 4 │   │   │   │   │   │   │   │   │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 3 │   │   │   │   │   │   │   │   │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 2 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
 1 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │
───┼───┼───┼───┼───┼───┼───┼───┼───┤
   │ A │ B │ C │ D │ E │ F │ G │ H │
*/
func renderBoard(board *Board, playerColor PieceColor) {
	var output string

	output += "───┬───┬───┬───┬───┬───┬───┬───┬───┤\n"

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
				colNumber = (index / 8) + 1
			} else {
				colNumber = 8 - (index / 8)
			}
			output += fmt.Sprintf(" %d │", colNumber)
		}

		if piece == &Empty {
			output += "   "
		} else {
			output += fmt.Sprintf(" %s ", piece)
		}

		if index%8 == 7 {
			output += "│\n"
			output += "───┼───┼───┼───┼───┼───┼───┼───┼───┤\n"
			if index == 63 {
				if playerColor == White {
					output += "   │ A │ B │ C │ D │ E │ F │ G │ H │\n"
				} else {
					output += "   │ H │ G │ F │ E │ D │ C │ B │ A │\n"
				}
			}
		} else {
			output += "│"
		}
	}

	Box := box.New(box.Config{Type: "Round", Color: "Cyan"})
	Box.Print("", output)
}
