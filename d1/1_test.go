package main

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"aoc24/common"
)

func TestPart_One(t *testing.T) {
	left, right := make([]int, 0, 1000), make([]int, 0, 1000)
	for _, v := range common.ReadLines("input.txt", 1000) {
		var a, b int
		fmt.Sscanf(v, "%d   %d", &a, &b)
		left = append(left, a)
		right = append(right, b)
	}
	sort.Ints(left)
	sort.Ints(right)
	var count int
	for i := 0; i < len(left); i++ {
		count += int(math.Abs(float64(left[i] - right[i])))
	}
	t.Log("Day I, part 1:", count)
}

func TestPart_Two(t *testing.T) {
	left, right := make([]int, 0, 1000), make([]int, 0, 1000)
	for _, v := range common.ReadLines("input.txt", 1000) {
		var a, b int
		fmt.Sscanf(v, "%d   %d", &a, &b)
		left = append(left, a)
		right = append(right, b)
	}
	var total int
	for _, l := range left {
		var count int
		for _, r := range right {
			if l == r {
				count++
			}
		}
		total += l * count
	}
	t.Log("Day I, part 2:", total)
}
