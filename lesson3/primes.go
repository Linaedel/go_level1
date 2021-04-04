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
	const needValidNumber = "Необходимо ввести число, умещающееся в int32. Повторите ввод."
	fmt.Println("Введите предел, до которого следует вывести простые числа")
	reader := bufio.NewReader(os.Stdin)
	var a int64

	for {
		incString, err := reader.ReadString('\n')
		incString = strings.TrimSpace(incString)
		if err != nil || !isNumeric(incString) {
			log.Println(needValidNumber)
		} else {
			a, err = strconv.ParseInt(incString, 10, 64)
			if err != nil {
				log.Println(needValidNumber)
			} else {
				break
			}
		}
	}
	primes := checkForPrimes(int(a))
	fmt.Println(primes)

}

func checkForPrimes(n int) []int {
	arr := make([]bool, n)
	for i := 2; i <= int(math.Sqrt(float64(n)+1)); i++ {
		if arr[i] == false {
			for j := i * i; j < n; j += i {
				arr[j] = true
			}
		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			primes = append(primes, i)
		}
	}

	return primes
}

func isNumeric(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
