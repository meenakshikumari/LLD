package main

type PieceO struct {
	*PlayingPiece
}

func NewPieceO() *PieceO {
	return &PieceO{PlayingPiece: NewPlayingPiece(pieceO)}
}
