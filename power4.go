package main

import "fmt"

type Power struct {
	Table         [][]string
	Player        [2]string
	CurrentPlayer string
	Winner        string
}

var Player = [2]string{"X", "O"}

func Starter(Player [2]string) {
	fmt.Println("Bienvenue dans le jeu Puissance 4!")
	fmt.Println("Joueurs:")
	for i := 0; i < 2; i++ {
		fmt.Printf("Joueur %d: %s\n", i+1, Player[i])
	}
}

func Play(p *Power, column int) bool {
	if column < 0 || column >= len(p.Table[0]) {
		fmt.Println("Colonne invalide. Choisissez une colonne entre 0 et", len(p.Table[0])-1)
		return false
	}
	for i := len(p.Table) - 1; i >= 0; i-- {
		if p.Table[i][column] == "" {
			p.Table[i][column] = p.CurrentPlayer
			p.switchPlayer()
			p.PrintBoard()
			return true
		}
	}
	fmt.Println("Colonne pleine. Choisissez une autre colonne.")
	return false
}

// NewPower retourne une structure Power avec la grille initialisée
func NewPower() *Power {
	p := &Power{}
	p.Player = Player
	p.CurrentPlayer = Player[0]
	p.Table = make([][]string, 6)
	for i := range p.Table {
		p.Table[i] = make([]string, 7)
	}
	return p
}

func (p *Power) switchPlayer() {
	if p.CurrentPlayer == p.Player[0] {
		p.CurrentPlayer = p.Player[1]
	} else {
		p.CurrentPlayer = p.Player[0]
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

func (p *Power) Win() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if p.Table[i][j] != "" {
				// Vérifier horizontal
				if j < 4 && p.Table[i][j] == p.Table[i][j+1] && p.Table[i][j] == p.Table[i][j+2] && p.Table[i][j] == p.Table[i][j+3] {
					fmt.Println("Joueur", p.Table[i][j], "gagne!")
					p.Winner = p.Table[i][j]
					return
				}
				// Vérifier vertical
				if i < 3 && p.Table[i][j] == p.Table[i+1][j] && p.Table[i][j] == p.Table[i+2][j] && p.Table[i][j] == p.Table[i+3][j] {
					fmt.Println("Joueur", p.Table[i][j], "gagne!")
					p.Winner = p.Table[i][j]
					return
				}
				// Vérifier diagonale (bas-gauche vers haut-droite)
				if i < 3 && j < 4 && p.Table[i][j] == p.Table[i+1][j+1] && p.Table[i][j] == p.Table[i+2][j+2] && p.Table[i][j] == p.Table[i+3][j+3] {
					fmt.Println("Joueur", p.Table[i][j], "gagne!")
					p.Winner = p.Table[i][j]
					return
				}
				// Vérifier diagonale (haut-gauche vers bas-droite)
				if i >= 3 && j < 4 && p.Table[i][j] == p.Table[i-1][j+1] && p.Table[i][j] == p.Table[i-2][j+2] && p.Table[i][j] == p.Table[i-3][j+3] {
					fmt.Println("Joueur", p.Table[i][j], "gagne!")
					p.Winner = p.Table[i][j]
					return
				}
			}
		}
	}
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
