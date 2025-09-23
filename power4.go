package main

import "fmt"

type Power struct {
	Table   [6][7]string
	player1 string
	player2 string
}

var player1 = "X"
var player2 = "O"

func play(p *Power, column int, player string) bool {
	fmt.Print("Où veux-tu jouer ? (1-7) ")
	var input int
	fmt.Scanln(&input)
	column = input - 1 // Adjust for 0-based index
	// Check if the column is valid
	if column < 0 || column >= 7 {
		fmt.Println("Colonne invalide. Choisissez une colonne entre 1 et 7.")
		return false
	}
	// Find the lowest empty row in the column
	for i := 5; i >= 0; i-- {
		if p.Table[i][column] == "" {
			p.Table[i][column] = player
			return true
		}
	}
	fmt.Println("Colonne pleine. Choisissez une autre colonne.")
	return false
}

// NewPower returns a Power struct with the Table initialized to empty spaces.
func NewPower() *Power {
	p := &Power{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			p.Table[i][j] = ""
		}
		fmt.Println(p.Table)
	}
	return p
}

func tour(p *Power, currentPlayer string) string {
	if currentPlayer == p.player1 {
		return p.player2
	}
	return p.player1
}

func (p *Power) Win() {
	// Logic to check for a win will be implemented here
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			fmt.Println("Checking cell:", i, j)
			// Add win checking logic here
			if p.Table[i][j] != "" {
				// Check horizontal
				if j < 4 && p.Table[i][j] == p.Table[i][j+1] && p.Table[i][j] == p.Table[i][j+2] && p.Table[i][j] == p.Table[i][j+3] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return
				}
				// Check vertical
				if i < 3 && p.Table[i][j] == p.Table[i+1][j] && p.Table[i][j] == p.Table[i+2][j] && p.Table[i][j] == p.Table[i+3][j] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return
				}
				// Check diagonal (bottom-left to top-right)
				if i < 3 && j < 4 && p.Table[i][j] == p.Table[i+1][j+1] && p.Table[i][j] == p.Table[i+2][j+2] && p.Table[i][j] == p.Table[i+3][j+3] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return
				}
				// Check diagonal (top-left to bottom-right)
				if i >= 3 && j < 4 && p.Table[i][j] == p.Table[i-1][j+1] && p.Table[i][j] == p.Table[i-2][j+2] && p.Table[i][j] == p.Table[i-3][j+3] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return
				}
			}
		}
	}
}
