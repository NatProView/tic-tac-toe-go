package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pvp(field *[3][3]string) int {
	clear()
	var x, y int
	//
	for i := 0; i < 9; i++ {
		clear()
		draw(*field)

		turn := i % 2

		if turn == 0 {
			fmt.Printf("\n\tTurn of [X]\n")
			x, y = playerInput(&field)
			(*field)[x][y] = "X"
		} else {
			fmt.Printf("\n\tTurn of [O]\n")
			x, y = playerInput(&field)
			(*field)[x][y] = "O"
		}

		switch check(*field) {
		case 1:
			return turn // 0: [X]	1: [O]
		case 2:
			return 3 // remis
		}
	}
	return -1
}

func pve(field *[3][3]string) int {
	var playerStarts bool
	var temp string
	if rand.Intn(2) == 1 {
		fmt.Println("Player gets [X] and starts the game")
		playerStarts = true //player gets [x]
	} else {
		fmt.Println("AI gets [X] and starts the game")
		playerStarts = false //bot gets [x]
	}

	for temp != "y" {
		fmt.Println("Ready? [y]")
		fmt.Scan(&temp)
	}
	var turn bool

	for i := 0; i < 9; i++ {
		clear()
		var x int
		var y int
		if i%2 == 0 {
			turn = true
		} else {
			turn = false
		}

		draw(*field)
		switch {
		case turn && playerStarts: //tura parzysta [X], gracz zaczyna
			fmt.Printf("\n\tTurn of [X]\n")
			x, y = playerInput(&field)
			(*field)[x][y] = "X"
		case turn && !playerStarts: //tura przysta [X], AI zaczyna
			fmt.Printf("\n\tTurn of [X]\n")
			x, y = botInput(&field)
			(*field)[x][y] = "X"
		case !turn && playerStarts: //tura nieparzysta [O], gracz zaczyna
			fmt.Printf("\n\tTurn of [O]\n")
			x, y = botInput(&field)
			(*field)[x][y] = "O"
		case !turn && !playerStarts: //tura nieparzysta [O], AI zaczyna
			fmt.Printf("\n\tTurn of [O]\n")
			x, y = playerInput(&field)
			(*field)[x][y] = "O"
		}
		switch check(*field) {
		case 1:
			return i % 2 // 0 X 1 O
		case 2:
			return 3 //remis
		}
	}

	return -1

}

func eve(field *[3][3]string) int {
	var x, y int

	for i := 0; i < 9; i++ {
		clear()
		draw(*field)
		turn := i % 2

		if turn == 0 {
			fmt.Printf("\n\tTurn of [X]\n")
			x, y = botInput(&field)
			(*field)[x][y] = "X"
		} else {
			fmt.Printf("\n\tTurn of [O]\n")
			x, y = botInput(&field)
			(*field)[x][y] = "O"
		}
		switch check(*field) {
		case 1:
			return i % 2
		case 2:
			return 3
		}
	}
	return -1
}
func playerInput(field **[3][3]string) (x, y int) {
	var input int
	isOk := false
	for !isOk {
		fmt.Scan(&input)
		x, y = toMultipleValues(input)
		if (**field)[x][y] != "O" && (**field)[x][y] != "X" {
			isOk = true
			return x, y
		}
		if !isOk {
			fmt.Println("Wrong input, tr again")
		}
	}
	return -1, -1
}
func botInput(field **[3][3]string) (x, y int) {
	isOk := false
	time.Sleep(2 * time.Second)
	for !isOk {
		x, y = toMultipleValues(rand.Intn(9) + 1)
		if (**field)[x][y] != "O" && (**field)[x][y] != "X" {
			isOk = true
			return x, y
		}
	}
	return -1, -1
}
func toMultipleValues(input int) (x, y int) {
	switch input {
	case 1:
		return 0, 0
	case 2:
		return 0, 1
	case 3:
		return 0, 2
	case 4:
		return 1, 0
	case 5:
		return 1, 1
	case 6:
		return 1, 2
	case 7:
		return 2, 0
	case 8:
		return 2, 1
	case 9:
		return 2, 2
	default:
		fmt.Println("something went wrong")
	}
	return
}
