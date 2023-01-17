package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//RECURSIVIDADE: experimenta dividir o total pelos numeros e ver se o array contem
// o resultado dessa conta
// senão depois experimenta ver se há algum numero proximo desse com que consigas obter com os restantes numeros

const (
	ASC  = true
	DESC = false

	DifferenceInitialValue = 100
)

type ContainsSolution struct {
	Number     int
	Difference int
}

func ChallengeSolver(total int, availableNumbers ...int) string {
	sliceSum := sumArray(availableNumbers)

	if sliceSum == total {
		return strings.Join(convertToStringArray(availableNumbers), " + ")
	} else if sliceSum > total {
		remainder := sliceSum - total

		if arrayContains(remainder, availableNumbers) {
			builder := strings.Builder{}
			builder.Grow(len(availableNumbers) * 2)

			result := removeElement(remainder, availableNumbers)

			return strings.Join(convertToStringArray(result), " + ")
		}
	}

	sortedNumbers := sortNumberArray(availableNumbers, DESC)

	fmt.Printf("sorted: [%v]\n", sortedNumbers)

	return calculateSolution(total, sortedNumbers)
}

func calculateSolution(total int, numbers []int) string {
	//for _, number := range numbers {
	//division := total / number
	//}

	return ""
}

func arrayContainsOrLowDifference(number int, array []int) ContainsSolution {
	solution := ContainsSolution{
		Difference: DifferenceInitialValue,
	}

	for i := range array {
		if array[i] == number {
			return ContainsSolution{
				Number:     array[i],
				Difference: 0,
			}
		}

		if number > array[i]-3 && number < array[i]+3 {
			difference := int(math.Abs(float64(number) - float64(array[i])))
			if difference < solution.Difference {
				solution = ContainsSolution{
					Number:     array[i],
					Difference: difference,
				}
			}
		}
	}

	return solution
}

func convertToStringArray(arr []int) []string {
	result := make([]string, len(arr))

	for i, val := range arr {
		result[i] = strconv.Itoa(val)
	}

	return result
}

func removeElement(toRemove int, numbers []int) []int {
	for i := range numbers {
		if numbers[i] == toRemove {
			return append(numbers[:i], numbers[i+1:]...)
		}
	}

	return numbers
}

func arrayContains(number int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == number {
			return true
		}
	}

	return false
}

func sumArray(numbers []int) int {
	var result int

	for _, val := range numbers {
		result += val
	}

	return result
}

func sortNumberArray(numbers []int, asc bool) []int {
	if asc {
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})
	} else {
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] > numbers[j]
		})
	}

	return numbers
}
