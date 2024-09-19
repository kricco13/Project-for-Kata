package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}

func calculate(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неправильный формат ввода")
	}

	num1, isRoman1 := parseNumber(parts[0])
	num2, isRoman2 := parseNumber(parts[2])

	if isRoman1 != isRoman2 {
		panic("Нельзя смешивать римские и арабские числа")
	}

	result := applyOperation(num1, num2, parts[1])
	if isRoman1 && result < 1 {
		panic("В римской системе нет отрицательных чисел")
	}

	if isRoman1 {
		return toRoman(result)
	}
	return strconv.Itoa(result)
}

func parseNumber(s string) (int, bool) {
	if value, isRoman := romanToArabic[s]; isRoman {
		return value, true
	}
	num, err := strconv.Atoi(s)
	if err != nil || num < 1 || num > 10 {
		panic("Число должно быть от 1 до 10")
	}
	return num, false
}

func applyOperation(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль")
		}
		return num1 / num2
	default:
		panic("Неправильный оператор")
	}
}

func toRoman(num int) string {
	var roman strings.Builder
	for _, rv := range arabicToRoman {
		for num >= rv.Value {
			roman.WriteString(rv.Symbol)
			num -= rv.Value
		}
	}
	return roman.String()
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Результат:", calculate(strings.TrimSpace(input)))
}
