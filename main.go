package main

import (
	"fmt"
	"math/rand"
)

func main() {
	check()
	var a string
	fmt.Scan(&a)
}

func check() {
	var usernumber int
	generatornumber()
	var programnumber int = generatornumber()
	fmt.Println("Угадай какое я загадал число)))")
	fmt.Scanf("%d\n", &usernumber)
	if usernumber > programnumber {
		fmt.Println("АХАха Лох твоё число больше загаданного. Загаданным число было: ", programnumber)

	} else if usernumber < programnumber {
		fmt.Println("Аахаха Лох твоё число меньше загаданного.Загаданным число было: ", programnumber)

	} else {
		fmt.Println("Хорош!!! Ты угадал!!!! Этим число было: ", programnumber)
	}
	fmt.Println("\nДавай сыграем ещё раз")
	check()
}

func generatornumber() int {
	var generator int = rand.Intn(20)
	return generator

}
