// https://cups.online/ru/tasks/1220

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := Scan()
	str = strings.Trim(str, "\n")
	array := strings.Split(str, " ")

	fmt.Printf("Квадрат со стороной %v\n\n", len(array))
	for i := 0; i < len(array); i++ {
		fmt.Printf("%v\n", str)
	}
}

func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
