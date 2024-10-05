package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	linearstats "linearstats/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter two args")
		return
	}

	if string(os.Args[1]) != "data.txt" {
		fmt.Println("File to read : data.txt")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Erreur file open :", err)
		return
	}
	defer file.Close()

	elements := bufio.NewScanner(file)

	Numbers := []int{}

	for elements.Scan() {

		nb, err := strconv.Atoi(elements.Text())
		if err != nil {
			continue
		}

		Numbers = append(Numbers, nb)

	}

	if err := elements.Err(); err != nil {
		return
	}

	a, b := linearstats.LinearRegression(Numbers)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", linearstats.PearsonCorrelationCoefficient(Numbers))
}
