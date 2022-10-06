package main

import (
	"fmt"
	"stablematching/algorithmgs"
)

func main() {
	MP := [][]int{
		{0, 1},
		{0, 1, 2},
		{0, 1, 2, 3},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4, 5},
	}
	WP := [][]int{
		{1, 0, 2, 3, 4},
		{2, 1, 0, 3, 4},
		{2, 4, 1, 0, 3},
		{3, 2, 4, 1, 0},
		{3, 2, 1, 4, 0},
	}
	result := algorithmgs.FindStableMatching(MP, WP)
	fmt.Println(result)
}
