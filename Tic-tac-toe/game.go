package main

import "fmt"

type Game struct {
	players []*Player
	board   *Board
}

func NewGame(board *Board, players []*Player) *Game {
	return &Game{
		board:   board,
		players: players,
	}
}

func (g *Game) StartGame() {
	isWinner := false
	for !isWinner {
		var r, c int
		player := g.players[0]
		g.players = g.players[1:]
		//copy(g.players, g.players[1:]) // better to use as no new memory allocated when we do => g.players = g.players[1:]
		g.players = append(g.players, player)

		correctMove := false
		for !correctMove {
			fmt.Printf("Player %s Enter row and col\n", player.Name)
			_, err := fmt.Scan(&r, &c)
			if err != nil {
				continue
			}
			correctMove = g.board.AddPiece(r, c, player.PlayingPiece)
		}

		g.board.PrintBoard()
		//check if the player is winner or game is tie
		isWinner = g.checkWinner(player, r, c)
		if isWinner {
			fmt.Printf("Player %s is winner!\n", player.Name)
		}

		if g.board.getFreeSpace() == 0 {
			break
		}
	}
	if !isWinner {
		fmt.Println("tie")
	}
	return
}

func (g *Game) checkWinner(player *Player, r, c int) bool {
	//row := []int{-1, 0, 1, 0}
	//col := []int{0, -1, 0, 1}
	rowMatch := true
	for i := 0; i < g.board.size; i++ { // check in rows
		if g.board.grid[r][i] == nil || g.board.grid[r][i] != player.PlayingPiece {
			rowMatch = false
			break
		}
	}

	colMatch := true
	for i := 0; i < g.board.size; i++ { // check in colum
		if g.board.grid[i][c] == nil || g.board.grid[i][c] != player.PlayingPiece {
			colMatch = false
			break
		}
	}

	diagonalMatch := true
	antiDiagonalMatch := true
	if r == c { // only possible if the piece has been kept at diagonal position
		for i := 0; i < g.board.size; i++ { // check in diagonal
			if g.board.grid[i][i] == nil || g.board.grid[i][i] != player.PlayingPiece {
				diagonalMatch = false
				break
			}
		}

		for i := g.board.size - 1; i >= 0; i-- { // check in diagonal
			if g.board.grid[i][i] == nil || g.board.grid[i][i] != player.PlayingPiece {
				antiDiagonalMatch = false
				break
			}
		}
	} else {
		diagonalMatch = false
		antiDiagonalMatch = false
	}

	return rowMatch || colMatch || diagonalMatch || antiDiagonalMatch
}

func (g *Game) AddPlayer(player *Player) {
	g.players = append(g.players, player)
}
