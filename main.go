package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
		var err error
		usernumber, err = PrintUserNumber()
		if err != nil {
			fmt.Println("Ошибка: введите число, а не текст! Введи ещё раз число")
			clearInput()
			i++
			continue
		}

		intArr = append(intArr, usernumber)
		switch {
		case usernumber > programmnumber:
			PrintMessLoseHigh()
		case usernumber < programmnumber:
			PrintMessLoseLow()
		case usernumber == programmnumber:
			PrintMessWin()
			i = -1
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

func clearInput() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func PrintUserNumber() (int, error) {
	var usernumber int
	_, err := fmt.Scan(&usernumber)
	if err != nil {
		return 0, err
	}
	return usernumber, nil
}

func PrintFirstMessGame() {
	fmt.Println("Угадай какое я загадал число за 3 попытки)))")
}

func generatornumber() int {
	var generator int = rand.Intn(20) + 1
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
