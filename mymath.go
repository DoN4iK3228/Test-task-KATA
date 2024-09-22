package mymath

import "math"

// Sqrt возвращает квадратный корень числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Ceil возвращает наименьшее целое число, не меньшее x
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor возвращает наибольшее целое число, не большее x
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Pow возводит число x в степень y
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Max возвращает большее из двух чисел
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min возвращает меньшее из двух чисел
func Min(x, y float64) float64 {
	return math.Min(x, y)
}
