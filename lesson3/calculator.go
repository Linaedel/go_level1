package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	const signs string = "+-*/"
	for {
		fmt.Println("Введите требуемеую математическую операцию в формате <число><оператор><число> " +
			"или exit для выхода")
		reader := bufio.NewReader(os.Stdin)
		var a, b, result float64
		var valid bool
		var operator string

		for {
			incString, err := reader.ReadString('\n')
			incString = strings.TrimSpace(incString)
			incString = strings.ReplaceAll(incString, ",", ".")
			if incString == "exit" {
				os.Exit(0)
			}
			valid, operator = containsAllowedSymbols(incString, signs)
			if err != nil || !valid {
				log.Println("Допустимый формат ввода: <число><оператор><число>. Повторите ввод.")
			} else {
				a, b, err = getProcessedValues(incString, operator)
				switch operator {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b
				case "/":
					result = a / b
				}
				break
			}
		}
		fmt.Println(a, operator, b, "=", result)
	}
}

func containsAllowedSymbols(str string, signs string) (bool, string) {
	hasDot, containsOperator := false, false
	var operator string
	for _, char := range str {
		if (char < '*' || char > '9') && (char != ' ') {
			return false, ""
		} else if char == '.' {
			if !hasDot {
				hasDot = true
			} else {
				return false, ""
			}
		} else if strings.ContainsRune(signs, char) {
			if containsOperator == false {
				containsOperator = true
				operator = string(char)
				hasDot = false
			} else {
				return false, ""
			}
		}
	}
	return true, operator
}

func getProcessedValues(str string, operator string) (float64, float64, error) {
	splittedStringArr := strings.Split(str, operator)
	var resultingError error = nil
	a, err1 := strconv.ParseFloat(strings.TrimSpace(splittedStringArr[0]), 64)
	b, err2 := strconv.ParseFloat(strings.TrimSpace(splittedStringArr[1]), 64)
	// Кривовато, надо бы агрегировать информацию, но пока не досуг читать про создание кастомных ошибок
	if err1 != nil {
		resultingError = err1
	}
	if err2 != nil {
		resultingError = err2
	}
	return a, b, resultingError
}
