package main

import "fmt"

type Power struct {
	Table         [6][7]string
	Player        [2]string
	currentPlayer string
}

var Player = [2]string{"X", "O"}

func Starter(Player [2]string) {
	fmt.Println("Bienvenue dans le jeu Puissance 4!")
	fmt.Println("Joueurs:")
	for i := 0; i < 2; i++ {
		fmt.Printf("Joueur %d: %s\n", i+1, Player[i])
	}
}

func Play(p *Power) bool {
	fmt.Print("Où veux-tu jouer ? (1-7) ")
	var input int
	fmt.Scanln(&input)
	column := input - 1 // Adjust for 0-based index
	// Check if the column is valid
	if column < 0 || column >= 7 {
		fmt.Println("Colonne invalide. Choisissez une colonne entre 1 et 7.")
		return false
	}
	// Find the lowest empty row in the column
	for i := 5; i >= 0; i-- {
		if p.Table[i][column] == "" {
			p.Table[i][column] = p.currentPlayer
			p.switchPlayer()
			p.PrintBoard()
			return true
		}
	}
	fmt.Println("Colonne pleine. Choisissez une autre colonne.")
	return false
}

// NewPower returns a Power struct with the Table initialized to empty spaces.
func NewPower() *Power {
	p := &Power{}
	p.Player = Player
	p.currentPlayer = Player[0]
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			p.Table[i][j] = ""
		}
	}
	return p
}

func (p *Power) switchPlayer() {
	if p.currentPlayer == p.Player[0] {
		p.currentPlayer = p.Player[1]
	} else {
		p.currentPlayer = p.Player[0]
	}
}

func (p *Power) PrintBoard() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if p.Table[i][j] == "" {
				fmt.Print(". ")
			} else {
				fmt.Print(p.Table[i][j], " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func (p *Power) Win() bool {
	// Logic to check for a win will be implemented here
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			// Add win checking logic here
			if p.Table[i][j] != "" {
				// Check horizontal
				if j < 4 && p.Table[i][j] == p.Table[i][j+1] && p.Table[i][j] == p.Table[i][j+2] && p.Table[i][j] == p.Table[i][j+3] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return true
				}
				// Check vertical
				if i < 3 && p.Table[i][j] == p.Table[i+1][j] && p.Table[i][j] == p.Table[i+2][j] && p.Table[i][j] == p.Table[i+3][j] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return true
				}
				// Check diagonal (bottom-left to top-right)
				if i < 3 && j < 4 && p.Table[i][j] == p.Table[i+1][j+1] && p.Table[i][j] == p.Table[i+2][j+2] && p.Table[i][j] == p.Table[i+3][j+3] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return true
				}
				// Check diagonal (top-left to bottom-right)
				if i >= 3 && j < 4 && p.Table[i][j] == p.Table[i-1][j+1] && p.Table[i][j] == p.Table[i-2][j+2] && p.Table[i][j] == p.Table[i-3][j+3] {
					fmt.Println("Player", p.Table[i][j], "wins!")
					return true
				}
			}
		}
	}
	return false
}

func (p *Power) IsDraw() bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if p.Table[i][j] == "" {
				return false
			}
		}
	}
	return true
}
