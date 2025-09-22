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
			p.Table[i][j] = " "
		}
	}
	fmt.Println(p.Table)
	return p
}

// func (p *Power) Display() {
// 	fmt.Println(" 0 1 2 3 4 5 6")
// 	for i := 0; i < 6; i++ {
// 		for j := 0; j < 7; j++ {
// 			fmt.Print("|" + p.Table[i][j])
// 		}
// 		fmt.Println("|")
// 	}
// 	fmt.Println("---------------")
// }

// func (p *Power) PlaceToken(column int, token string) bool {
// 	if column < 0 || column > 6 {
// 		return false
// 	}
// 	for i := 5; i >= 0; i-- {
// 		if p.Table[i][column] == " " {
// 			p.Table[i][column] = token
// 			return true
// 		}
// 	}
// 	return false
// }
