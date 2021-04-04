package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите площадь круга:")
	reader := bufio.NewReader(os.Stdin)
	var num float64
	for {
		incString, err := reader.ReadString('\n')
		incString = strings.TrimSpace(incString)
		if err != nil || !isFloatNumber(incString) {
			log.Println("Допустимо только одно число с точкой в качестве разделителя. Повторите ввод.")
		} else {
			num, err = strconv.ParseFloat(incString, 32)
			if err != nil {
				log.Println("Ошибка при конвертации строки в число", err, "Повторите ввод")
			} else {
				break
			}
		}
	}
	r := math.Sqrt(num / math.Pi)
	fmt.Printf("Диаметр круга: %0.2f \n", 2*r)
	fmt.Printf("Длина окружности:  %0.2f \n", 2*math.Pi*r)
}

func isFloatNumber(str string) bool {
	noPoints := true
	for _, char := range str {
		if (char < '0' || char > '9') && (char != '.') {
			return false
		} else if char == '.' {
			if noPoints {
				noPoints = false
			} else {
				return false
			}
		}
	}
	return true
}
