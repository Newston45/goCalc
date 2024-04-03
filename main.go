package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arabicToRoman = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X"}

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

func checkAndConvertToArabic(num string) (int, bool) {
	isRoman := false

	resNum, _ := strconv.Atoi(num) // получаем число int
	if resNum == 0 {
		resNum = romanToArabic[num]
		isRoman = true
		return resNum, isRoman
	}
	return resNum, isRoman
}

func findTwoNumsAndOperator(Text string) (int, int, string) { // находим оператор и два числа
	operators := []string{"+", "-", "/", "*"} // возможные арифметические операции

	Text = strings.Trim(Text, "\n")          // удаляем перенос строки из записанной строки
	Text = strings.ReplaceAll(Text, " ", "") // удаляем все пробелы

	var operator string
	for _, op := range operators { // проходимся циклом по возможным операциям
		if strings.Contains(Text, op) { // и проверяем есть ли в введной строке символ операции
			operator = op // если нашли записываем оператор
			break
		}
	}

	nums := strings.Split(Text, operator) // убираем оператор, получаем два числа из строки

	fmt.Println(nums)

	var num1isRoman, num2isRoman bool
	num1, num1isRoman := checkAndConvertToArabic(nums[0])
	num2, num2isRoman := checkAndConvertToArabic(nums[1])
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num1isRoman)
	fmt.Println(num2isRoman)

	if num1isRoman == true && num2isRoman == true {

	}

	return num1, num2, operator

}

func calculate(num1, num2 int, operator string) int { // выполнение математичских операций
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Это калькулятор!")

	for {
		fmt.Println("Введите математическую операцию c двумя целыми числами:")
		text, _ := reader.ReadString('\n') // записываем введенную пользователем строку

		fmt.Println("Результат вычисления: ", calculate(findTwoNumsAndOperator(text)))
	}

}
