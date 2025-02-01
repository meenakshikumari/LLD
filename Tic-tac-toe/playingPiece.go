package main

type PlayingPiece struct {
	PieceType PieceType
}

func NewPlayingPiece(t PieceType) *PlayingPiece {
	return &PlayingPiece{PieceType: t}
}

func (p *PlayingPiece) Type() PieceType {
	return p.PieceType
}
