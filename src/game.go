package main

import (
	"fmt"
	"math/rand"
	"time"
)

//funkcja pozwalająca na odczytanie inputu od uzytkownika
//wywołanie jej skutkuje zmiana jednego elementu arraya na znak X albo Y

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

//funkcja generująca input od bota,
//wywołanie jej skutkuje zmiana jednego elementu arraya na znak X albo Y
func inputBot(field **[]string, char string) {
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

//tryb gry player vs player, bierze dwa inputy od graczy, na zmiane

func pvp(field *[]string) int {
	clear()

	//kazda kolejna iteracja funkcji to kolejna tura gry
	// dla i parzystego tura [X]
	// dla i nieparzystego tura [O]
	for i := 0; i < 9; i++ {

		draw(*field)
		turn := i % 2

		if turn == 0 { //i parzyste, tura [X]
			input((&field), "X")
		} else { //i nieparzyste, tura [O]
			input((&field), "O")
		}

		switch check(*field) {
		case 1:
			clear()
			return turn // 0: [X]	1: [O]
		case 2:
			clear()
			return 3 // remis
		}

		clear()

	}
	return -1 //błąd
}

//AI vs AI, zero inputu od uzytkownika
func eve(field *[]string) int {
	clear()
	// kazda kolejna iteracja funkcji to kolejna tura gry
	// dla i parzystego tura [X]
	// dla i nieparzystego tura [O]
	for i := 0; i < 9; i++ {

		draw(*field)
		turn := i % 2

		if turn == 0 {
			inputBot((&field), "X")
		} else {
			inputBot((&field), "O")
		}
		switch check(*field) {
		case 1:
			clear()
			return turn // 0: [X]	1: [O]
		case 2:
			clear()
			return 3 // remis
		}

		clear()

	}
	return -1
}

//gracz vs AI, jeden input od uzytkownika
func pve(field *[]string) int {
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
	clear()
	var turn bool

	// kazda kolejna iteracja funkcji to kolejna tura gry
	// dla i parzystego tura [X]
	// dla i nieparzystego tura [O]
	for i := 0; i < 9; i++ {

		if i%2 == 0 {
			turn = true // [X]
		} else {
			turn = false // [O]
		}

		draw(*field)

		switch {
		case turn && playerStarts: //tura parzysta [X], gracz zaczyna
			input((&field), "X")
		case turn && !playerStarts: //tura przysta [X], AI zaczyna
			inputBot((&field), "X")
		case !turn && playerStarts: //tura nieparzysta [O], gracz zaczyna
			inputBot((&field), "O")
		case !turn && !playerStarts: //tura nieparzysta [O], AI zaczyna
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
