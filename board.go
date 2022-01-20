package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/Delta456/box-cli-maker/v2"
)

type PieceType uint

const (
	Pawn PieceType = iota
	Bishop
	Knight
	Rook
	King
	Queen
)

type CastingAvailability uint

const (
	NoCastling         CastingAvailability = 0
	WhiteKingCastling                      = 1
	WhiteQueenCastling                     = 2
	BlackKingCastling                      = 4
	BlackQueenCastling                     = 8
)

type Piece struct {
	Type    PieceType
	isWhite bool
}

func (p Piece) String() string {
	if !p.isWhite {

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

type Board struct {
	placements      [64]*Piece
	whiteActive     bool
	castling        byte
	enPassant       byte
	halfmoveCounter uint
	moveCounter     uint
}

const startingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Get all available moves from all pieces of a player
func getAvailableMoves(board *Board, playerIsWhite bool) [][4]byte {
	availableMoves := make([][4]byte, 0, 64)

	for index, piece := range board.placements {
		if piece != &Empty && piece.isWhite == playerIsWhite {
			col, row := getCoordsFromIndex(byte(index))

			var moves [][2]byte
			if piece.Type == Pawn {
				moves = getAvailablePawnMoves(board, col, row)
			}
			if piece.Type == Rook {
				moves = getAvailableRookMoves(board, col, row)
			}
			if piece.Type == Knight {
				moves = getAvailableKnightMoves(board, col, row)
			}
			if piece.Type == Bishop {
				moves = getAvailableBishopMoves(board, col, row)
			}
			if piece.Type == Queen {
				moves = getAvailableQueenMoves(board, col, row)
			}
			if piece.Type == King {
				moves = getAvailableKingMoves(board, col, row)
			}

			for _, move := range moves {
				availableMoves = append(availableMoves, [4]byte{col, row, move[0], move[1]})
			}
		}
	}

	return availableMoves
}

// Ensures a move is valid for current board
func checkMove(board *Board, pieceIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) bool {
	pieceIndex := getIndexFromCoords(startCol, startRow)
	piece := board.placements[pieceIndex]

	if piece == &Empty {
		return false
	}

	switch piece.Type {
	case Pawn:
		return checkPawnMove(board, pieceIsWhite, startCol, startRow, endCol, endRow)
	case Rook:
		return checkRookMove(board, pieceIsWhite, startCol, startRow, endCol, endRow)
	case Knight:
		return checkKnightMove(board, pieceIsWhite, startCol, startRow, endCol, endRow)
	case Bishop:
		return checkBishopMove(board, pieceIsWhite, startCol, startRow, endCol, endRow)
	case Queen:
		return checkQueenMove(board, pieceIsWhite, startCol, startRow, endCol, endRow)
	case King:
		return checkKingMove(board, pieceIsWhite, startCol, startRow, endCol, endRow)
	}

	return false
}

// Verifies if the player king is in check
func checkCheck(board *Board, playerIsWhite bool) bool {
	var kingRow, kingCol byte

	for index, piece := range board.placements {
		if piece.Type == King && ((playerIsWhite && piece.isWhite) || (!playerIsWhite && piece.isWhite)) {
			kingRow, kingCol = getCoordsFromIndex(byte(index))
		}
	}

	for index, piece := range board.placements {
		if piece != &Empty && piece.isWhite != playerIsWhite {
			startCol, startRow := getCoordsFromIndex(byte(index))

			if checkMove(board, piece.isWhite, startCol, startRow, kingRow, kingCol) {
				return true
			}
		}
	}

	return false
}

/* Verifies if the player king will be captured next turn.
Gets every available move, create all possible future boards,
and checks if at least one is not in check anymore
*/
func checkMate(board *Board, playerIsWhite bool) bool {
	moves := getAvailableMoves(board, playerIsWhite)

	for _, move := range moves {
		testBoard := *board
		doMove(&testBoard, playerIsWhite, move[0], move[1], move[2], move[3])
		if !checkCheck(&testBoard, playerIsWhite) {
			return false
		}
	}

	return true
}

// Execute a movement if it's valid
func doMove(board *Board, playerIsWhite bool, startCol byte, startRow byte, endCol byte, endRow byte) (Board, error) {
	startIndex := getIndexFromCoords(startCol, startRow)
	startPiece := board.placements[startIndex]

	if startPiece.isWhite != playerIsWhite {
		return *board, errors.New("Not your piece")
	}

	if !checkMove(board, playerIsWhite, startCol, startRow, endCol, endRow) {
		return *board, errors.New("Invalid move")
	}

	endIndex := getIndexFromCoords(endCol, endRow)

	backup := board.placements[endIndex]
	board.placements[endIndex] = board.placements[startIndex]
	board.placements[startIndex] = &Empty

	if checkCheck(board, playerIsWhite) {
		board.placements[startIndex] = board.placements[endIndex]
		board.placements[endIndex] = backup

		return *board, errors.New("Would lose the game")
	}

	// Update metadata
	if checkCheck(board, !playerIsWhite) || backup != &Empty {
		board.halfmoveCounter = 0
	} else {
		board.halfmoveCounter++
	}
	if !playerIsWhite {
		board.moveCounter++
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
func renderBoard(board *Board, playerIsWhite bool) {
	var output string

	output += "───┬───┬───┬───┬───┬───┬───┬───┬───┤\n"

	for index := 0; index < 64; index++ {
		var piece *Piece
		if playerIsWhite {
			piece = board.placements[63-index]

		} else {
			piece = board.placements[index]
		}

		if index%8 == 0 {
			var colNumber int
			if playerIsWhite {
				colNumber = 8 - (index / 8)

			} else {
				colNumber = (index / 8) + 1
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
				if playerIsWhite {
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

func exportBoard(board *Board) string {
	var output string

	// 1 - Piece Placement
	emptyCounter := 0
	for index, piece := range board.placements {
		if piece == &Empty {
			emptyCounter++
		} else {
			if emptyCounter != 0 {
				output += strconv.Itoa(emptyCounter)
				emptyCounter = 0
			}

			if piece.isWhite {
				output += string(unicode.ToUpper(rankCharMap[piece.Type]))
			} else {
				output += string(unicode.ToLower(rankCharMap[piece.Type]))

			}
		}

		if (index+1)%8 == 0 {
			if emptyCounter > 0 {
				output += strconv.Itoa(emptyCounter)
			}
			emptyCounter = 0
			output += "/"
		}
	}

	output += " "

	// 2 - Active color
	if board.whiteActive {
		output += "w"
	} else {
		output += "b"
	}

	output += " "

	// 3 - Castling availability
	castlingFlag := board.castling
	castlingString := ""

	if castlingFlag>>3 > 0 {
		castlingFlag >>= 1
		castlingString += "q"
	}
	if castlingFlag>>2 > 0 {
		castlingFlag >>= 1
		castlingString += "k"
	}
	if castlingFlag>>1 > 0 {
		castlingFlag >>= 1
		castlingString += "Q"
	}
	if castlingFlag>>0 > 0 {
		castlingFlag >>= 1
		castlingString += "K"
	}

	if len(castlingString) == 0 {
		output += "-"
	} else {
		output += Reverse(castlingString)
	}

	output += " "

	// 4 - En passant target square
	if board.enPassant == 0 {
		output += "-"
	} else {
		col, row := getCoordsFromIndex(board.enPassant)
		output += string(rune(col+'a')) + string(row+1)
	}

	output += " "

	// 5 - Halfmove clock
	output += fmt.Sprint(board.halfmoveCounter)

	output += " "

	// 6 - Fullmove number
	output += fmt.Sprint(board.moveCounter)

	return output
}

func importBoard(s string) Board {
	var board Board
	fields := strings.Split(s, " ")

	// 1 - Piece Placement
	placements := strings.Split(fields[0], "/")
	index := 0

	for _, row := range placements {
		for _, placement := range row {
			if unicode.IsDigit(placement) {
				spaces := placement - '0'

				for spaces > 0 {
					board.placements[index] = &Empty
					index++
					spaces--
				}
			} else {
				var isWhite bool
				rank := charRankMap[unicode.ToUpper(placement)]

				if placement == unicode.ToUpper(placement) {
					isWhite = true
				} else {
					isWhite = false
				}

				board.placements[index] = &Piece{rank, isWhite}

				index++
			}
		}
	}

	// 2 - Active color
	if fields[1] == "w" {
		board.whiteActive = true
	} else {
		board.whiteActive = false
	}

	// 3 - Castling availability
	for _, availability := range fields[2] {
		if availability == 'K' {
			board.castling += WhiteKingCastling
		}
		if availability == 'Q' {
			board.castling += WhiteQueenCastling
		}
		if availability == 'k' {
			board.castling += BlackKingCastling
		}
		if availability == 'q' {
			board.castling += BlackQueenCastling
		}
	}

	// 4 - En passant target square
	if fields[3] != "-" {
		targetColumn := fields[3][0] - 'a'
		targetRow := fields[3][1] - '0'
		board.enPassant = getIndexFromCoords(targetColumn, targetRow)
	}

	// 5 - Halfmove clock
	halfmoveCounter, _ := strconv.ParseUint(fields[4], 10, 8)
	board.halfmoveCounter = uint(halfmoveCounter)

	// 6 - Fullmove number
	moveCounter, _ := strconv.ParseUint(fields[5], 10, 8)
	board.moveCounter = uint(moveCounter)

	return board
}
