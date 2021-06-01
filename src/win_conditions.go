package main

//import "strconv"
/*
func isDraw(field []string) bool {
	for i := 1; i < 10; i++ {
		if strconv.Itoa(i) != field[i] {
			return false //jeszcze sie nie skonczyla
		}
	}
	return true //remis
}
*/
func check(field []string) int {
	//if isDraw(field) {
	//	return 2 //gra zakonczona remisem
	//}

	if field[1] == field[2] && field[2] == field[3] { //poziomo 1 2 3
		return 1
	} else if field[4] == field[5] && field[5] == field[6] { //poziomo 4 5 6
		return 1
	} else if field[7] == field[8] && field[8] == field[9] { //poziomo 7 8 9
		return 1
	} else if field[1] == field[4] && field[4] == field[7] { //pionowo 1 4 7
		return 1
	} else if field[2] == field[5] && field[5] == field[8] { //pionowo 2 5 8
		return 1
	} else if field[3] == field[6] && field[6] == field[9] { //pionowo 3 6 9
		return 1
	} else if field[1] == field[5] && field[5] == field[9] { //przekatnie 1 5 9
		return 1
	} else if field[3] == field[5] && field[5] == field[7] { //przekatnie 3 5 7
		return 1
	} else if field[1] != "1" && field[2] != "2" && field[3] != "3" && //TODO: ZROBIC LADNIEJSZA WERSJE TEGO
		field[4] != "4" && field[5] != "5" && field[6] != "6" && //DZIALA ALE IRYTUJE MNIE SWOIM BYTEM
		field[7] != "7" && field[8] != "8" && field[9] != "9" {
		return 2
	} else {
		return 0 //gra dalej trwa
	}
}
