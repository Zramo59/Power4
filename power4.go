package main

import "fmt"

type Power struct {
	Table [6][7]string
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
