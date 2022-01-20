package main

const BlackPawnRune = "♟︎"
const BlackBishopRune = "♝"
const BlackKnightRune = "♞"
const BlackRookRune = "♜"
const BlackQueenRune = "♛"
const BlackKingRune = "♚"

const WhitePawnRune = "♙"
const WhiteBishopRune = "♗"
const WhiteKnightRune = "♘"
const WhiteRookRune = "♖"
const WhiteQueenRune = "♕"
const WhiteKingRune = "♔"

var charRankMap = map[rune]PieceType{
	'K': King,
	'Q': Queen,
	'R': Rook,
	'B': Bishop,
	'N': Knight,
	'P': Pawn,
}

var rankCharMap = map[PieceType]rune{
	King:   'K',
	Queen:  'Q',
	Rook:   'R',
	Bishop: 'B',
	Knight: 'N',
	Pawn:   'P',
}
