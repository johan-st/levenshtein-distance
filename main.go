package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	for i := 0; i < 1000000; i++ {
		levenshtein([]rune("jihgfedcba"), []rune("abcdefg"))
	}

	// fmt.Printf("distance was %v.\n", dist)
}

func levenshtein(a []rune, b []rune) int {
	// fmt.Printf("checking distance between %c and %c.\n", a, b)
	matrix := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		matrix[i] = make([]int, len(b)+1)

		for j := range matrix[i] {
			matrix[i][j] = 0

		}
	}

	for i := 0; i < len(matrix); i++ {
		matrix[i][0] = i
	}

	for i := 0; i < len(matrix[0]); i++ {
		matrix[0][i] = i
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			substitutionCost := 0
			if a[i-1] != b[j-1] {
				substitutionCost = 1
			}
			matrix[i][j] = min3(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+substitutionCost)
		}
	}
	// printMatrix(matrix)
	return matrix[len(matrix)-1][len(matrix[0])-1]
}

func min3(x int, y int, z int) int {
	if x < y && x < z {
		return x
	} else if y < x && y < z {
		return y
	} else {
		return z
	}
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("row %v -> \t", i)
		for j := range matrix[i] {
			fmt.Printf("%v\t", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-----------------------\n")

}
