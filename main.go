package main

import (
	"os"
	"strconv"
	"fmt"
	"math"
	"errors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need a size argument")
		os.Exit(1)
	}

	arg := os.Args[1]
	size, err := strconv.Atoi(arg)

	pals := make(map[int]int)

	if err != nil {
		fmt.Println("Size must be an integer")
		fmt.Println(err)
		os.Exit(1)
	}

	for sqr := 1; sqr < size; sqr++ {
		for cube := 1; cube < size; cube++ {
			current := int(math.Pow(float64(sqr), 2) + math.Pow(float64(cube), 3))
			str := strconv.Itoa(current)

			if isPalindrome(str) {
				pals[current]++
			}

			if pals[current] == 4 {
				fmt.Println("Found 4 ways with", str)
				result, err := checkPalindromes(pals)
				if err == nil {
					fmt.Println("Finished with", result)
					os.Exit(0)
				}
			}
		}
	}
}

func checkPalindromes(pals map[int]int) (int, error) {
	sum := 0
	numVals := 0
	for i, count := range pals {
		if count == 4 {
			sum += i
			numVals++
		}
	}

	if numVals == 5 {
		return sum, nil
	}
	return 0, errors.New("don't have 5 palindromes yet")
}
