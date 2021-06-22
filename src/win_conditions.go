package main

func check(field [3][3]string) int {
	//if isDraw(field) {
	//	return 2 //gra zakonczona remisem
	//}

	if field[0][0] == field[0][1] && field[0][1] == field[0][2] { //poziomo 1 2 3
		return 1
	} else if field[1][0] == field[1][1] && field[1][1] == field[1][2] { //poziomo 4 5 6
		return 1
	} else if field[2][0] == field[2][1] && field[2][1] == field[2][2] { //poziomo 7 8 9
		return 1
	} else if field[0][0] == field[1][0] && field[1][0] == field[2][0] { //pionowo 1 4 7
		return 1
	} else if field[0][1] == field[1][1] && field[1][1] == field[2][1] { //pionowo 2 5 8
		return 1
	} else if field[0][2] == field[1][2] && field[1][2] == field[2][2] { //pionowo 3 6 9
		return 1
	} else if field[0][0] == field[1][1] && field[1][1] == field[2][2] { //przekatnie 1 5 9
		return 1
	} else if field[0][2] == field[1][1] && field[1][1] == field[2][0] { //przekatnie 3 5 7
		return 1
	} else if field[0][0] != "1" && field[0][1] != "2" && field[0][2] != "3" &&
		field[1][0] != "4" && field[1][1] != "5" && field[1][2] != "6" &&
		field[2][0] != "7" && field[2][1] != "8" && field[2][2] != "9" {
		return 2
	} else {
		return 0 //gra dalej trwa
	}
}
