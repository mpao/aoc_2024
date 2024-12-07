package main

import (
	"fmt"
	"slices"
	"testing"

	"aoc24/common"
)

var grid [][]rune

func init() {
	for _, line := range common.ReadLines("input.txt", 0) {
		grid = append(grid, []rune(line))
	}
}

type guard struct {
	x    int
	y    int
	dir  rune
	path []string
}

func (g *guard) position() string {
	return fmt.Sprintf("%d %d %s", g.x, g.y, string(g.dir))
}

func (g *guard) move() {
	switch g.dir {
	case '^': // 94
		if g.y > 0 && grid[g.y-1][g.x] != '#' {
			g.y--
		} else {
			g.dir = '>'
		}
	case '>': // 62
		if g.x < len(grid[0]) && grid[g.y][g.x+1] != '#' {
			g.x++
		} else {
			g.dir = 'v'
		}
	case 'v': // 118
		if g.y < len(grid) && grid[g.y+1][g.x] != '#' {
			g.y++
		} else {
			g.dir = '<'
		}
	case '<': // 60
		if g.x > 0 && grid[g.y][g.x-1] != '#' {
			g.x--
		} else {
			g.dir = '^'
		}
	}
	g.path = append(g.path, g.position())
}

func (g *guard) exit() bool {
	switch g.dir {
	case '^':
		return g.y == 0
	case 'v':
		return g.y == len(grid)-1
	case '>':
		return g.x == len(grid[0])-1
	case '<':
		return g.x == 0
	default:
		return true
	}
}

func TestPart_One(t *testing.T) {
	var g guard
	for y, row := range grid {
		for x, r := range row {
			if r == '^' || r == '<' || r == '>' || r == 'v' {
				p := fmt.Sprintf("%d %d %s", x, y, string(r))
				g = guard{x, y, r, []string{p}}
			}
		}
	}
	positions := make(map[string]struct{}, 0)
	for !g.exit() {
		positions[fmt.Sprintf("%d %d", g.x, g.y)] = struct{}{}
		g.move()
	}
	t.Log("Day VI, part 1:", len(positions)+1) // moves + initial position
}

func TestPart_Two(t *testing.T) {
	var g guard
	// find start position
	for y, row := range grid {
		for x, r := range row {
			if r == '^' || r == '<' || r == '>' || r == 'v' {
				p := fmt.Sprintf("%d %d %s", x, y, string(r))
				g = guard{x, y, r, []string{p}}
			}
		}
	}
	startX, startY, startR := g.x, g.y, g.dir
	// find the exit path
	for !g.exit() {
		g.move()
	}
	path := g.path
	count := make(map[string]struct{})
	for i := 1; i < len(path); i++ {
		y, x := block(path[i])
		grid[y][x] = '#'
		p := fmt.Sprintf("%d %d %s", startX, startY, string(startR))
		g = guard{startX, startY, startR, []string{p}}
		for !g.exit() {
			g.move()
			if len(g.path) > 3 { // 4 is the minimum path size for a loop
				lastMoves := g.path[len(g.path)-2:]
				sub := g.path[:len(g.path)-2]
				if slices.Contains(sub, lastMoves[0]) {
					i := slices.Index(sub, lastMoves[0])
					if i+1 < len(sub)-1 && sub[i+1] == lastMoves[1] {
						// loop!
						k := fmt.Sprintf("%d %d", x, y)
						count[k] = struct{}{}
						break
					}
				}
			}
		}
		grid[y][x] = '.'
	}
	t.Log("Day VI, part 2:", len(count))
}

func block(in string) (int, int) {
	var x, y int
	var r rune
	fmt.Sscanf(in, "%d %d %c", &x, &y, &r)
	return y, x
}
