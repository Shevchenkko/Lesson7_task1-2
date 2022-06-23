package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// На вхід подано стрінг з цілими числами, котрі розділені пробілами.
// Треба повернути найбільше та найменше число.
// 	Наприклад:
// 	input := "1 2 3 4 5" // повертає "5 1"
// 	input := "1 9 3 4 -5" // повертає "9 -5"
// Уточнення:
// 1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
// 2. В стрінгі завжди буде принаймні одне число.
// 3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
// число). Найбільше число має бути першим.
// 4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.

func main() {
	var min, max int32
	// input := "1 2 3 4 5"
	input := "11 4 5 -67 99 76 2 45 9"

	// розділяємо строку по пробілах
	varInput := strings.Split(input, " ")

	// шукаємо найбільше та найменше число
	for i, val := range varInput {
		// переводимо строку у int
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		// переводимо int у int32
		int32Val := int32(intVal)

		if i == 0 {
			max, min = int32Val, int32Val
		}

		if max < int32Val {
			max = int32Val
		}
		if min > int32Val {
			min = int32Val
		}
	}

	// переводимо отримані дані у string
	result := strconv.Itoa(int(max)) + " " + strconv.Itoa(int(min))
	fmt.Println(result)
}
