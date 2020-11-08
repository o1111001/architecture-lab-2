package lab2

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

type fnType func (float64, float64) float64

func plus(a float64, b float64) float64 {
	return b + a
}

func minus(a float64, b float64) float64 {
	return b - a
}

func mult(a float64, b float64) float64 {
	return b * a
}

func divide(a float64, b float64) float64 {
	return b / a
}

func pow(a float64, b float64) float64 {
	return math.Pow(b, a)
}

var functionsMap = map[string] fnType {
	"+": plus,
	"-": minus,
	"/": divide,
	"*": mult,
	"^": pow,
}

func stringToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func EvaluatePostfix(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("Expression is empty string")
	}
	var stack []float64
	s := strings.Split(input, " ")

	for _, elemen := range s {
		if fn, ok := functionsMap[element]; ok {
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], fn(a, b))
		} else {
			floatNumber, err := stringToFloat(element)
			if err != nil {
				return "", fmt.Errorf("Expression contains invalid characters")
			}
			stack = append(stack, floatNumber)
		}
	}
	return fmt.Sprintf("%.0f", stack[0]), nil
}