package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	numbers := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nbr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("The input: %v is not a number.\n", scanner.Text())
			continue
		}
		numbers = append(numbers, nbr)
		min, max := guessNextNbr(numbers)

		fmt.Println(min, max)
	}

}

func guessNextNbr(numbers []int) (int, int) {
	var min, max int
	min = averege(numbers) - stdDev(numbers)
	max = averege(numbers) + stdDev(numbers)
	return min, max
}

func stdDev(numbers []int) int {
	avg := averege(numbers)
	var sum int
	for _, nbr := range numbers {
		diff := nbr - avg
		sum += diff * diff
	}

	variance := sum / len(numbers)
	return int(math.Sqrt(float64(variance)))
}

func averege(numbers []int) int {
	var sum int
	for _, nbr := range numbers {
		sum += nbr
	}
	return sum / len(numbers)
}
