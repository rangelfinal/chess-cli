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

var charRankMap = map[string]PieceType{
	"K": King,
	"Q": Queen,
	"R": Rook,
	"B": Bishop,
	"N": Knight,
}

var rankCharMap = map[PieceType]string{
	King:   "K",
	Queen:  "Q",
	Rook:   "R",
	Bishop: "B",
	Knight: "N",
}
