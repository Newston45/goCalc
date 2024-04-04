package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkAndConvertToArabic(num string) (int, bool) {
	var romanToArabic = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10}

	isRoman := false

	resNum, _ := strconv.Atoi(num) // получаем число int

	if resNum == 0 {
		resNum = romanToArabic[num] // конвертируем число в арабское
		isRoman = true
		return resNum, isRoman
	}
	if resNum < 1 || resNum > 10 {
		panic("число меньше 1 или больше 10")
	}
	return resNum, isRoman
}

// функция перевода арабского числа в римское
func arabicToRoman(arabic int) string {
	conversions := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conv := range conversions { // проходимся по структуре
		for arabic >= conv.Value { // и сравниваем наше число со значением в ней
			roman.WriteString(conv.Symbol) // записываем римский символ
			arabic -= conv.Value           // вычитаем значение равное римскому символу из исходного числа
		}
	}
	return roman.String() // возвращаем собранную строку
}

// функция выполнения математичских операций
func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("На ноль делить нельзя!")
		}
		return num1 / num2
	default:
		panic("Неверная операция!")
	}
}

func findResult(Text string) string {
	operators := []string{"+", "-", "/", "*"} // возможные арифметические операции

	Text = strings.Trim(Text, "\n") // удаляем перенос строки из записанной строки
	if Text == "exit" {
		os.Exit(0)
	}
	Text = strings.ReplaceAll(Text, " ", "") // удаляем все пробелы

	var operator string
	for _, op := range operators { // проходимся циклом по возможным операциям
		if strings.Contains(Text, op) { // и проверяем есть ли в введной строке символ операции
			operator = op // если нашли записываем оператор
			break
		}
		panic("нет опертора вычисления")
	}

	nums := strings.Split(Text, operator) // убираем оператор, получаем два числа из строки

	if len(nums) > 2 {
		panic("неверное количество опернадов")
	}

	var num1isRoman, num2isRoman bool
	num1, num1isRoman := checkAndConvertToArabic(nums[0]) //проверяем числа и конвертируем их
	num2, num2isRoman := checkAndConvertToArabic(nums[1])

	var result int

	if num1isRoman == false && num2isRoman == true { // проверяем что числа в одной системе счисления
		panic("разные системы счисления")
	} else if num1isRoman == true && num2isRoman == false {
		panic("разные системы счисления")
	} else if num1isRoman == true && num2isRoman == true {
		result = calculate(num1, num2, operator)
		if result <= 0 {
			panic("отрицательное римское число")
		}
		return arabicToRoman(result) // если числа римские то результат выдаем в римских
	}

	result = calculate(num1, num2, operator)
	resultString := strconv.Itoa(result)

	return resultString

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Это калькулятор!")

	for {
		fmt.Println("Введите математическую операцию c двумя целыми числами:")
		fmt.Println("или введите exit для закрытия программы")
		inputText, _ := reader.ReadString('\n') // записываем введенную пользователем строку

		fmt.Println("Результат вычисления: ", findResult(inputText))
	}

}
