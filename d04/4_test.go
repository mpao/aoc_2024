package main

import (
	"testing"

	"aoc24/common"
)

func TestPart_One(t *testing.T) {
	var matrix [][]rune
	for _, line := range common.ReadLines("input.txt", 0) {
		matrix = append(matrix, []rune(line))
	}
	var words []string
	// like a wind rose: for every char, compose the word 4 chars long in the 8 directions
	for y, line := range matrix {
		for x, r := range line {
			if x > 2 && y > 2 {
				NW := string([]rune{r, matrix[y-1][x-1], matrix[y-2][x-2], matrix[y-3][x-3]})
				words = append(words, NW)
			}
			if y > 2 {
				N := string([]rune{r, matrix[y-1][x], matrix[y-2][x], matrix[y-3][x]})
				words = append(words, N)
			}
			if y > 2 && x < len(line)-3 {
				NE := string([]rune{r, matrix[y-1][x+1], matrix[y-2][x+2], matrix[y-3][x+3]})
				words = append(words, NE)
			}
			if x < len(line)-3 {
				E := string([]rune{r, matrix[y][x+1], matrix[y][x+2], matrix[y][x+3]})
				words = append(words, E)
			}
			if y < len(matrix)-3 && x < len(line)-3 {
				SE := string([]rune{r, matrix[y+1][x+1], matrix[y+2][x+2], matrix[y+3][x+3]})
				words = append(words, SE)
			}
			if y < len(matrix)-3 {
				S := string([]rune{r, matrix[y+1][x], matrix[y+2][x], matrix[y+3][x]})
				words = append(words, S)
			}
			if y < len(matrix)-3 && x > 2 {
				SW := string([]rune{r, matrix[y+1][x-1], matrix[y+2][x-2], matrix[y+3][x-3]})
				words = append(words, SW)
			}
			if x > 2 {
				W := string([]rune{r, matrix[y][x-1], matrix[y][x-2], matrix[y][x-3]})
				words = append(words, W)
			}
		}
	}
	var count int
	for _, w := range words {
		if w == "XMAS" {
			count++
		}
	}
	t.Log("Day IV, part 1:", count)
}

func TestPart_Two(t *testing.T) {
	var matrix [][]rune
	for _, line := range common.ReadLines("input.txt", 0) {
		matrix = append(matrix, []rune(line))
	}
	var count int
	// find two MAS in the shape of an X
	for y, line := range matrix {
		if y < 1 || y > len(matrix)-2 {
			continue
		}
		for x, r := range line {
			if x < 1 || x > len(line)-2 {
				continue
			}
			if r != 'A' {
				continue
			}
			NW := matrix[y-1][x-1]
			NE := matrix[y-1][x+1]
			SW := matrix[y+1][x-1]
			SE := matrix[y+1][x+1]
			if (NW == 'S' && SE == 'M' && NE == 'S' && SW == 'M') ||
				(NW == 'S' && SE == 'M' && NE == 'M' && SW == 'S') ||
				(NW == 'M' && SE == 'S' && NE == 'S' && SW == 'M') ||
				(NW == 'M' && SE == 'S' && NE == 'M' && SW == 'S') {
				count++
			}
		}
	}
	t.Log("Day IV, part 2:", count)
}
