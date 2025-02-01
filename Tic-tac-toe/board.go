package main

import "fmt"

type Board struct {
	size int
	grid [][]*PlayingPiece
}

func NewBoard(size int) *Board {
	grid := make([][]*PlayingPiece, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]*PlayingPiece, size)
	}
	return &Board{
		size: size,
		grid: grid,
	}
}

func (b *Board) AddPiece(r, c int, playingPiece *PlayingPiece) bool {
	if !(r >= 0 && r < b.size && c >= 0 && c < b.size) {
		fmt.Println("INVALID PLACE!! PLEASE TRY AGAIN")
		return false
	} else if b.grid[r][c] != nil {
		fmt.Println("ALREADY PRESENT!! PLEASE TRY AGAIN")
		return false
	}
	b.grid[r][c] = playingPiece
	return true
}

func (b *Board) getFreeSpace() int {
	var cnt int
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			if b.grid[r][c] == nil {
				cnt++
			}
		}
	}
	return cnt
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fmt.Print(b.grid[i][j])
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}
