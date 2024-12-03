package common

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func ReadLines(filename string, lines int) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("file not found")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([]string, 0, lines)
	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}
	return slices.Clip(out)
}

func ReadAll(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("file not found")
	}
	return string(b)
}
