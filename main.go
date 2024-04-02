package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkNumVal(val int) bool {
	if val < 1 || val > 10 {
		return false
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	operators := []string{"+", "-", "/", "*"} // возможные арифметические операции
	var result int

	fmt.Println("Это калькулятор!")

	for {
		fmt.Println("Введите что вы хотите посчитать.")

		text, _ := reader.ReadString('\n')       // записываем введенную пользователем строку
		text = strings.Trim(text, "\n")          // удаляем перенос строки из записанной строки
		text = strings.ReplaceAll(text, " ", "") // удаляем все пробелы

		var operator string
		for _, op := range operators { // проходимся циклом по возможным операциям
			if strings.Contains(text, op) { // и проверяем есть ли в введной строке символ операции
				operator = op // если нашли записываем оператор
				break
			}
		}

		nums := strings.Split(text, operator) // убираем оператор, находим два числа из строки
		num1, _ := strconv.Atoi(nums[0])      // получаем первое число int
		num2, _ := strconv.Atoi(nums[1])      // получаем второе число int

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 != 0 {
				result = num1 / num2
			}
			panic("На ноль делить нельзя!")
		}

		fmt.Println("Результат вычисления: ", result)

	}

}
