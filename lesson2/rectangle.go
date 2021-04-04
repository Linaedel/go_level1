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
	fmt.Println("Введите значение сторон прямоугольника, разделённые пробелом:")
	reader := bufio.NewReader(os.Stdin)
	var a, b float64
	for {
		incString, err := reader.ReadString('\n')
		incString = strings.TrimSpace(incString)
		if err != nil || !containsTwoFloatNumbers(incString) {
			log.Println("Допустимо только одно число с точкой в качестве разделителя. Повторите ввод.")
		} else {
			splittedStringArr := strings.Split(incString, " ")
			var err1, err2 error
			a, err1 = strconv.ParseFloat(splittedStringArr[0], 64)
			b, err2 = strconv.ParseFloat(splittedStringArr[1], 64)
			if err1 != nil || err2 != nil {
				log.Println("Ошибка при конвертации строки в число", err1, err2, "Повторите ввод")
			} else {
				break
			}
		}
	}
	fmt.Println("Площадь прямоугольника с заданными сторонами:", a*b)

}

func containsTwoFloatNumbers(str string) bool {
	points := 0
	noSpace := true
	for _, char := range str {
		if (char < '0' || char > '9') && (char != '.') && (char != ' ') {
			return false
		} else if char == ' ' {
			if noSpace {
				noSpace = false
			} else {
				return false
			}
		} else if char == '.' {
			if points < 2 {
				points++
			} else {
				return false
			}
		}
	}
	return true
}
