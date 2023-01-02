package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	table = [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C"},
	}
)

func arab2roman(arab int) (string, error) {

	if arab <= 0 || arab > 100 {
		return "", errors.New("Такого числа нет")
	}

	var (
		roman = ""
		m     = 100
	)
	for i := 2; i >= 0; i-- {
		d := arab / m
		roman += table[i][d]
		arab %= m
		m /= 10
	}

	return roman, nil
}

func str_conv(a string) (int, error, bool) {
	i, err := strconv.Atoi(a)
	if err != nil {
		rom := append(table[0], "X")
		for i := 1; i < len(rom); i++ { // len возвращает количество элементов
			if a == rom[i] {
				return i, nil, true
				break
			}
		}
		return 0, err, false
	}
	return i, nil, false

}

func calc(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("Такого знака нет")

	}
}

func main() {

	var a, b, c string
	fmt.Scanln(&a, &c, &b)
	x, err, flag := str_conv(a)
	if err != nil {
		panic(err)
	}
	y, err, flag2 := str_conv(b)
	if err != nil {
		panic(err)
	}
	if flag == flag2 {
		res, err := calc(x, y, c)
		if err != nil {
			panic(err)
		} else if flag2 {
			if (res < 1) {
				panic("Результат меньше единицы")
			}
			n, err := arab2roman(res)
			if err != nil {
				panic(err)
			}
			fmt.Println(n)
		} else {
			fmt.Println(res)
		}
	} else {
		panic("Римские и Арабские не поддерживаются")
	}

}
