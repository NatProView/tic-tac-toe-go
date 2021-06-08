package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//deklaracja arraya zawierającego znak wyswietlany w grze
	field := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var winner int

	//funkcja witająca gracza przy starcie gry
	hello()

	//wybór trybu gry, kazda z funkcji zwraca int odpowiadający temu kto wygral
	switch choice() {

	case 1: //player vs player
		winner = pvp(&field)

	case 2: //player vs AI
		winner = pve(&field)

	case 3: //AI vs AI
		winner = eve(&field)

	default:
		fmt.Println("something went wrong")
	}

	//wyswietlenie ekranu wyników
	endResult(winner, field)

}

func choice() int {
	fmt.Printf("\n\n")
	fmt.Println("[1] Player vs Player")
	fmt.Println("[2] Player vs AI")
	fmt.Println("[3] AI vs AI")
	var temp int
	fmt.Scan(&temp)
	if temp > 0 && temp < 4 {
		return temp
	} else {
		fmt.Println("Wrong input, try again")
		return choice()
	}
}

//czyszczenie ekranu terminala
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//rysowanie planszy
func draw(field []string) {
	fmt.Printf("\tTic Tac Toe\n\n")
	fmt.Printf("\t %s | %s | %s \n", field[1], field[2], field[3])
	fmt.Printf("\t---|---|---\n")
	fmt.Printf("\t %s | %s | %s \n", field[4], field[5], field[6])
	fmt.Printf("\t---|---|---\n")
	fmt.Printf("\t %s | %s | %s \n", field[7], field[8], field[9])
}

func hello() {
	clear()
	fmt.Printf("\n\t WELCOME TO\n\n")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\t t | i | c ")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("\t---|---|---")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\t t | a | c ")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("\t---|---|---")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\t t | o | e ")
}

func endResult(winner int, field []string) {
	draw(field)
	fmt.Printf("\n\n")
	switch winner {
	case 0:
		fmt.Println("[X] HAS WON THE GAME")
	case 1:
		fmt.Println("[O] HAS WON THE GAME")
	case 3:
		fmt.Println("WE'VE GOT A DRAW!")
	}
}
