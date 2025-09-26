package main

import (
	"fmt"
)

func main() {
	p := NewPower()
	Starter(p.Player)
	p.PrintBoard()
	for {
		if Play(p) {
			if p.Win() {
				break
			} else if p.IsDraw() {
				fmt.Println("Match nul!")
				break
			}
		}
	}
}
