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
	fmt.Println("Угадай какое я загадал число за 3 попытки)))")
	fmt.Scanf("%d\n", &usernumber)
	for i := 2; i >= 0; i-- {
		if usernumber > programnumber {
			fmt.Println("АХАха Лох твоё число больше загаданного.")

		} else if usernumber < programnumber {
			fmt.Println("Аахаха Лох твоё число меньше загаданного.")

		} else {
			i = 0
			fmt.Println("Хорош!!! Ты угадал!!!! ")

		}
		if i == 0 {
			fmt.Println("Этим число было: ", programnumber)
			fmt.Println("\nДавай сыграем ещё раз")
			check()
		} else {
			fmt.Println("Осталось попыток: ", i, "Угадай снова какое число загдал")
			fmt.Scanf("%d\n", &usernumber)
		}

	}

}

func generatornumber() int {
	var generator int = rand.Intn(20)
	return generator
}
