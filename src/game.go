package main

import (
	"fmt"
	"math/rand"
	"time"
)

//gracz vs gracz, dwa inputy od uzytkownika
func input(field **[]string, char string) { //0 ok | 1 nie ok XD
	var input int
	isOk := false
	fmt.Printf("\n\tTurn of [%s]\n", char)
	for !isOk {
		fmt.Scan(&input)
		for k, v := range **field {
			if k == input && v != "X" && v != "O" {
				(**field)[k] = char
				isOk = true
			}
		}
		if !isOk {
			fmt.Println("[Input] Wrong input, try again")
		}

	}

}

func inputBot(field **[]string, char string) { //0 ok | 1 nie ok XD
	var input int
	isOk := false
	fmt.Printf("\n\tTurn of [%s]\n", char)
	time.Sleep(2 * time.Second)
	for !isOk {
		input = (rand.Intn(9) + 1)
		for k, v := range **field {
			if k == input && v != "X" && v != "O" {
				(**field)[k] = char
				isOk = true
			}
		}

	}

}

func pvp(field *[]string) int {
	clear()
	for i := 0; i < 10; i++ {

		draw(*field)

		if i%2 == 0 {
			input((&field), "X")
		} else {
			input((&field), "O")
		}
		switch check(*field) {
		case 1:
			clear()
			return i % 2 // 0: [X]	1: [O]
		case 2:
			clear()
			return 3 // remis
		}

		clear()

	}
	return -1
}

//AI vs AI, zero inputu od uzytkownika
func eve(field *[]string) int {
	clear()
	for i := 0; i < 9; i++ {

		draw(*field)

		if i%2 == 0 {
			inputBot((&field), "X")
		} else {
			inputBot((&field), "O")
		}
		switch check(*field) {
		case 1:
			clear()
			return i % 2 // 0: [X]	1: [O]
		case 2:
			clear()
			return 3 // remis
		}

		clear()

	}
	return -1
}

//gracs vs AI, jeden input od uzytkownika
func pve(field *[]string) int {
	var random bool
	var temp string
	if rand.Intn(2) == 1 {
		fmt.Println("Player gets [X] and starts the game")
		random = true //player gets [x]
	} else {
		fmt.Println("AI gets [X] and starts the game")
		random = false //bot gets [x]
	}

	for temp != "y" {
		fmt.Println("Ready? [y]")
		fmt.Scan(&temp)
	}
	clear()
	var turn bool
	for i := 0; i < 9; i++ {

		if i%2 == 0 {
			turn = true // [X]
		} else {
			turn = false // [O]
		}

		draw(*field)

		switch {
		case turn && random:
			input((&field), "X")
		case turn && !random:
			inputBot((&field), "X")
		case !turn && random:
			inputBot((&field), "O")
		case !turn && !random:
			input((&field), "O")
		}

		switch check(*field) {
		case 1:
			clear()
			return i % 2 // 0: [X]	1: [O]
		case 2:
			clear()
			return 3 // remis
		}

		clear()

	}
	return -1

}
