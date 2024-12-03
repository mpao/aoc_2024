package main

import (
	"aoc24/common"
	"fmt"
	"regexp"
	"testing"
)

func TestPart_One(t *testing.T) {
	s := common.ReadAll("input.txt")
	muls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	var count int
	for _, m := range muls.FindAllString(s, -1) {
		var a, b int
		fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
		count += a * b
	}
	t.Log("Day III, part 1:", count)
}

func TestPart_Two(t *testing.T) {
	s := common.ReadAll("input.txt")
	muls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	mulsIndex := muls.FindAllStringIndex(s, -1)
	dont := regexp.MustCompile(`don't()`).FindAllStringIndex(s, -1)
	do := regexp.MustCompile(`do()`).FindAllStringIndex(s, -1)
	type disable struct {
		state bool
		index int
	}
	var count int
	var cmd disable
	for i, m := range muls.FindAllString(s, -1) {
		var a, b int
		fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
		index := mulsIndex[i][0]
		for _, d := range dont {
			indexDont := d[0]
			if indexDont < index && cmd.index < indexDont {
				cmd = disable{state: true, index: indexDont}
			}
		}
		for _, d := range do {
			indexDo := d[0]
			if indexDo < index && cmd.index < indexDo {
				cmd = disable{state: false, index: indexDo}
			}
		}
		if !cmd.state {
			count += a * b
		}
	}
	t.Log("Day III, part 2:", count)
}
