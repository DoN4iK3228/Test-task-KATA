package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arabicNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	var result, Dig1, Dig2 int
	var err error
	arDig1 := false
	arDig2 := false

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if len(text) <= 1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	Parts := strings.Split(text, " ")
	oper := Parts[1]

	if len(Parts) < 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
	} else if len(Parts) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	if value, ok := arabicNumerals[Parts[0]]; ok {
		Dig1 = value
		arDig1 = true

	} else {
		Dig1, err = strconv.Atoi(Parts[0])
		if Dig1 < 0 || Dig1 > 10 {
			panic("Выдача паники, так как одно число или более не входит в диапазон от 1 до 10 включительно")
		}
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
	}

	if value, ok := arabicNumerals[Parts[2]]; ok {
		Dig2 = value
		arDig2 = true
	} else {
		Dig2, err = strconv.Atoi(Parts[2])
		if Dig2 < 0 || Dig2 > 10 {
			panic("Выдача паники, так как одно число или более не входит в диапазон от 1 до 10 включительно")
		}
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
	}

	if arDig1 == true && arDig2 == true {
		result = calculate(Dig1, Dig2, oper)
		if result < 0 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		for key, v := range arabicNumerals {
			if v == result {
				fmt.Print(key)
			}
		}
	} else if arDig1 == false && arDig2 == false {
		result = calculate(Dig1, Dig2, oper)
		fmt.Print(result)
	} else {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
}

func calculate(Dig1, Dig2 int, operator string) int {
	switch operator {
	case "+":
		return Dig1 + Dig2
	case "-":
		return Dig1 - Dig2
	case "*":
		return Dig1 * Dig2
	case "/":
		if Dig2 == 0 {
			panic("Выдача паники, так как нельзя делить на 0.")
		}
		return Dig1 / Dig2
	default:
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — оператор (+, -, /, *).")
	}
}
