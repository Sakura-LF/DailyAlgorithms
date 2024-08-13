package main

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	matrix := GenerateMatrix(3)
	fmt.Println(matrix)
}

func TestGenerateMatrix1(t *testing.T) {
	GenerateMatrix(3)
}
