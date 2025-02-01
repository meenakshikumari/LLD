package main

func main() {
	board := NewBoard(3)
	var players []*Player
	players = append(players, NewPlayer("Alok", pieceX))
	players = append(players, NewPlayer("Manisha", pieceO))
	game := NewGame(board, players)
	game.StartGame()
}
