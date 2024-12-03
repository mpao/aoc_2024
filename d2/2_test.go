package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"testing"

	"aoc24/common"
)

func TestPart_One(t *testing.T) {
	var safecount int
	safe := func(in []int) error {
		for i := 0; i < len(in)-1; i++ {
			diff := float64(in[i] - in[i+1])
			if math.Abs(diff) == 0 || math.Abs(diff) > 3 {
				return errors.New("")
			}
			if (in[0] > in[1]) != (in[i] > in[i+1]) {
				return errors.New("")
			}
		}
		return nil
	}
	for _, v := range common.ReadLines("input.txt", 1000) {
		var report []int
		for _, i := range strings.Split(v, " ") {
			v, _ := strconv.Atoi(i)
			report = append(report, v)
		}
		if err := safe(report); err == nil {
			safecount++
		}
	}
	t.Log("Day II, part 1:", safecount)
}

func TestPart_Two(t *testing.T) {
	var safecount int
	safe := func(in []int) error {
		problemDampener := false
		for i := 0; i < len(in)-1; i++ {
			diff := float64(in[i] - in[i+1])
			if math.Abs(diff) == 0 || math.Abs(diff) > 3 {
				if !problemDampener {
					problemDampener = true
					continue
				}
				return errors.New("")
			}
			if (in[0] > in[1]) != (in[i] > in[i+1]) {
				if !problemDampener {
					problemDampener = true
					continue
				}
				return errors.New("")
			}
		}
		return nil
	}
	for _, v := range common.ReadLines("input.txt", 1000) {
		var report []int
		for _, i := range strings.Split(v, " ") {
			v, _ := strconv.Atoi(i)
			report = append(report, v)
		}
		if err := safe(report); err == nil {
			safecount++
		}
	}
	t.Log("Day II, part 2:", safecount)
}
