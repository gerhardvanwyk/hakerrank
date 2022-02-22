package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// Expected 2d array with 1 line N (N = square matrix size), then values.
func TestDiagonalDiff(t *testing.T) {

	lines := strings.Split(in, "\n")
	var arr [][]int32
	var err error
	for i, l := range lines {
		if i == 0 {
			arr[0][0], err = strconv.ParseInt(l, 10, 32)
			if err != nil {
				return "Input Failed", err
			}
		}
	}

	n := arr[0][0]
	var posDiag = int32(0)
	var negDiag = int32(0)
	// Write your code here
	for i := 1; i < int(n+1); i++ {
		for j := 0; j < int(n); j++ {
			posDiag += arr[i+1][j+1]
			negDiag += arr[i+1][n-1]
		}
	}

	fmt.Println(math.Abs(float64(posDiag) - float64(negDiag)))
}
