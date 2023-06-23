package recursion

import (
	"strings"
)

type Point struct {
	X int
	Y int
}

func SolveMaze(maze *[]string, start *Point) []*Point {
	path := []*Point{}
	seen := [][]bool{}

	m := *maze

	for i := 0; i < len(m); i++ {
		seen = append(seen, []bool{})
		for j := 0; j < len(m[0]); j++ {
			seen[i] = append(seen[i], false)
		}
	}

	if success, path := walk(maze, start, seen, path); success {
		m := *maze
		for i := 0; i < len(m); i++ {
			row := strings.Split(m[i], "")
			newRow := ""
			for j := 0; j < len(m[0]); j++ {
				if isInPath(i, j, path) {
					newRow = newRow + "*"
				} else {
					newRow = newRow + row[j]
				}
			}
			m[i] = newRow
		}

		return path
	}

	return nil
}

func walk(maze *[]string, cur *Point, seen [][]bool, path []*Point) (bool, []*Point) {
	m := *maze

	if cur.Y < 0 || cur.Y == len(m) ||
		cur.X < 0 || cur.X >= len(m[0]) {
		return false, path
	}

	row := strings.Split(m[cur.Y], "")

	if row[cur.X] == "#" {
		return false, path
	} else if row[cur.X] == "E" {
		return true, path
	} else if seen[cur.Y][cur.X] {
		return false, path
	}

	seen[cur.Y][cur.X] = true
	path = append(path, cur)

	if n, np := walk(maze, &Point{X: cur.X, Y: cur.Y + 1}, seen, path); n {
		return n, np
	} else if s, sp := walk(maze, &Point{X: cur.X, Y: cur.Y - 1}, seen, path); s {
		return s, sp
	} else if e, ep := walk(maze, &Point{X: cur.X + 1, Y: cur.Y}, seen, path); e {
		return e, ep
	} else if w, wp := walk(maze, &Point{X: cur.X - 1, Y: cur.Y}, seen, path); w {
		return w, wp
	}

	return false, path
}

func isInPath(y int, x int, path []*Point) bool {
	for _, p := range path {
		if p.X == x && p.Y == y {
			return true
		}
	}

	return false
}
