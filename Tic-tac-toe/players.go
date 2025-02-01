package main

type Player struct {
	Name         string
	PlayingPiece *PlayingPiece
}

func NewPlayer(name string, t PieceType) *Player {
	return &Player{
		Name:         name,
		PlayingPiece: NewPlayingPiece(t),
	}
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetPieceType() PieceType {
	return p.PlayingPiece.Type()
}
