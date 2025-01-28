package main

import (
	"fmt"
	"math/rand"
)

var intArr []int

func main() {
	check()
	var a string
	fmt.Scan(&a)
}

func check() {
	var usernumber int
	var programmnumber int = generatornumber()
	PrintFirstMessGame()
	intArr = nil
	for i := 2; i >= 0; i-- {
		usernumber = PrintUserNumber(usernumber)
		intArr = append(intArr, usernumber)
		switch {
		case usernumber > programmnumber:
			PrintMessLoseHigh()
		case usernumber < programmnumber:
			PrintMessLoseLow()
		case usernumber == programmnumber:
			i = 0
			PrintMessWin()
		}
		if i > 0 {
			LeftTry(i)
		}

	}
	Answer(programmnumber)
	PrintMessAgainGame()
	PrintAllAnswer()
	check()

}

func PrintUserNumber(usernumber int) int {
	fmt.Scanf("%d\n", &usernumber)
	return usernumber
}

func PrintFirstMessGame() {
	fmt.Println("Угадай какое я загадал число за 3 попытки)))")
}

func generatornumber() int {
	var generator int = rand.Intn(20)
	return generator
}

func PrintMessLoseLow() {
	fmt.Println("Аахаха Лох твоё число меньше загаданного. ")
}

func PrintMessLoseHigh() {
	fmt.Println("АХАха Лох твоё число больше загаданного. ")
}

func PrintMessAgainGame() {
	fmt.Print("\nДавай сыграем ещё раз. ")
}

func PrintMessWin() {
	fmt.Println("Хорош!!! Ты угадал!!!! ")
}

func LeftTry(i int) {
	fmt.Printf("Осталось попыток: %v Угадай снова какое число загдал\n", i)
}

func Answer(programmnumber int) {
	fmt.Printf("Этим число было: %v", programmnumber)
}

func PrintAllAnswer() {
	fmt.Println("Ваши ответы за игру:", intArr)
}
