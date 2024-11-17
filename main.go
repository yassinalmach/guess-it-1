package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nbr, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalln(err)
			return
		}
		numbers = append(numbers, nbr)
		avg := calculateAvrege(numbers)
        fmt.Println(avg)
	}

}

func calculateAvrege(numbers []float64) float64 {
	var div float64
	var sum float64
	for _, v := range numbers {
		sum += v
		div++
	}
	return sum / div
}
