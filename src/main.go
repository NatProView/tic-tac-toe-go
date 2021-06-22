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
	field := [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	var winner int
	isOn := true
	hello()

	//wybór trybu gry, kazda z funkcji zwraca int odpowiadający temu kto wygral
	for isOn {
		switch choice() {

		case 1: //player vs player

			winner = pvp(&field)
			isOn = false

		case 2: //player vs AI
			winner = pve(&field)
			isOn = false
		case 3: //AI vs AI
			winner = eve(&field)
			isOn = false

		default:
			fmt.Println("something went wrong")
		}
	}
	endResult(winner, field)

}

func endResult(winner int, field [3][3]string) {
	clear()
	draw(field)
	switch winner {
	case 0:
		fmt.Println("[X] HAS WON THE GAME")
	case 1:
		fmt.Println("[O] HAS WoN ThE GAME")
	case 3:
		fmt.Println("IT'S A DRAW")
	case -1:
		fmt.Println("Something went wrong :(")
	}
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
func draw(field [3][3]string) {
	fmt.Printf("\tTic Tac Toe\n\n")
	fmt.Printf("\t %s | %s | %s \n", field[0][0], field[0][1], field[0][2])
	fmt.Printf("\t---|---|---\n")
	fmt.Printf("\t %s | %s | %s \n", field[1][0], field[1][1], field[1][2])
	fmt.Printf("\t---|---|---\n")
	fmt.Printf("\t %s | %s | %s \n", field[2][0], field[2][1], field[2][2])
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
