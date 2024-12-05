package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"

	"aoc24/common"
)

var rules []string
var updates [][]int

func init() {
	var empty bool
	for _, line := range common.ReadLines("input.txt", 0) {
		if line == "" {
			empty = true
		}
		if empty && line != "" {
			var tmp []int
			for _, v := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(v)
				tmp = append(tmp, n)
			}
			updates = append(updates, tmp)
		} else {
			rules = append(rules, line)
		}
	}
}

func TestPart_One(t *testing.T) {
	var count int
	for _, up := range updates {
		valid := true
		for i := 0; i < len(up)-1; i++ {
			s := fmt.Sprintf("%d|%d", up[i], up[i+1])
			if !slices.Contains(rules, s) {
				valid = false
				break
			}
		}
		if valid {
			count += up[len(up)/2]
		}
	}
	t.Log("Day V, part 1:", count)
}

func TestPart_Two(t *testing.T) {
	var incorrectlyOrderedUpdates [][]int
	for _, up := range updates {
		for i := 0; i < len(up)-1; i++ {
			s := fmt.Sprintf("%d|%d", up[i], up[i+1])
			if !slices.Contains(rules, s) {
				incorrectlyOrderedUpdates = append(incorrectlyOrderedUpdates, up)
				break
			}
		}
	}
	var count int
	for _, v := range incorrectlyOrderedUpdates {
		slices.SortStableFunc(v, func(a, b int) int {
			s := fmt.Sprintf("%d|%d", a, b)
			if slices.Contains(rules, s) {
				return -1
			}
			return 1
		})
		count += v[len(v)/2]
	}
	t.Log("Day V, part 2:", count)
}
