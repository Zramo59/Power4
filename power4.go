package main

import "fmt"

type Power struct {
	Table   [6][7]string
	player1 map[string]int
	player2 map[string]int
}

func (p *Power) PlayerAffect() {
	fmt.Println("Que voulez-vous comme symbole pour le joueur 1 ? (\"X/O\")")
	input := ""
	fmt.Scanln(&input)
	if input == "X" {
		p.player1["X"] = 1
		p.player2["O"] = 2
	}
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

func (p *Power) Game() {
	// Game logic will be implemented here
	if len(p.player1) == 0 {
		fmt.Println("C'est au joueur 1 de jouer")
	} else {
		fmt.Println("C'est au joueur 2 de jouer")
	}
	// Continuez avec la logique du jeu
	// Par exemple, vous pouvez demander au joueur de choisir une colonne et de placer son jeton
}
