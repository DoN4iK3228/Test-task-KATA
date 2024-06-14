package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arabicNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
		"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
		"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
		"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
		"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
		"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
		"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
		"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	}
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
		if Dig1 > 10 {
			panic("Выдача паники, так как первое число не входит в диапазон от 1 до 10 включительно")
		}
		arDig1 = true
	} else {
		Dig1, err = strconv.Atoi(Parts[0])
		if Dig1 < 1 || Dig1 > 10 {
			panic("Выдача паники, так как первое число не входит в диапазон от 1 до 10 включительно")
		}
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
	}

	if value, ok := arabicNumerals[Parts[2]]; ok {
		Dig2 = value
		if Dig2 > 10 {
			panic("Выдача паники, так как второе число не входит в диапазон от 1 до 10 включительно")
		}
		arDig2 = true
	} else {
		Dig2, err = strconv.Atoi(Parts[2])
		if Dig2 < 1 || Dig2 > 10 {
			panic("Выдача паники, так как второе число не входит в диапазон от 1 до 10 включительно")
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
