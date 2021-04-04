package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Введите трёхзначное число:")
	reader := bufio.NewReader(os.Stdin)
	var stringNum string
	for {
		var err error
		stringNum, err = reader.ReadString('\n')
		stringNum = strings.TrimSpace(stringNum)
		//_, err := fmt.Scanf("%s", &stringNum)
		if err != nil || utf8.RuneCountInString(stringNum) != 3 || !isNumeric(stringNum) {
			log.Println("Неверный формат ввода, ожидается трёхзначное число. Повторите ввод.")
		} else {
			break
		}
	}
	var splitNum [3]int64
	for i, char := range stringNum {
		splitNum[i], _ = strconv.ParseInt(string(char), 10, 32)
	}
	fmt.Printf("Количество сотен: %d \nКоличество десятков: %d \nКоличество единиц: %d \n",
		splitNum[0], splitNum[1], splitNum[2])
}

func isNumeric(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
