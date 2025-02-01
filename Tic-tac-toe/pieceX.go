package main

type PieceX struct {
	*PlayingPiece
}

func NewPieceX() *PieceX {
	return &PieceX{PlayingPiece: NewPlayingPiece(pieceX)}
}
