package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (c Calc) Operate(input, operator string) {
	values := strings.Split(input, operator)

	x, err := parseInt(values[0])
	if err != nil {
		fmt.Printf("%s no es un numero", values[0])

		return
	}

	y, err := parseInt(values[1])
	if err != nil {
		fmt.Printf("%s no es un numero", values[0])

		return
	}

	switch operator {
	case "+":
		fmt.Printf("La suma de %d + %d = %d\n", x, y, x+y)

	case "-":
		fmt.Printf("La resta de %d + %d = %d\n", x, y, x-y)

	case "*":
		fmt.Printf("La multiplicacion de %d + %d = %d\n", x, y, x*y)

	case "/":
		if y == 0 {
			fmt.Println("La division entre 0 no es soportada")

			return
		}

		fmt.Printf("La division de %d + %d = %d\n", x, y, x/y)

	default:
		fmt.Printf("El operador %s no esta soportado\n", operator)
	}
}

func parseInt(input string) (int, error) {
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func main() {
	input := ReadInput()
	operator := ReadInput()

	c := Calc{}
	c.Operate(input, operator)
}
