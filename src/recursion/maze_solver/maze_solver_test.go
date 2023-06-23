package recursion

import (
	"fmt"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"########E#",
		"# #####  #",
		"# #     ##",
		"#   ######",
		"##    ## #",
		"####S    #",
	}

	solution := "{4 5}{4 4}{3 4}{3 3}{3 2}{4 2}{5 2}{6 2}{7 2}{7 1}{8 1}"

	start := Point{X: 4, Y: 5}
	path := SolveMaze(&maze, &start)
	attempt := ""

	if path == nil {
		t.Error("Failed to solve the maze")
	} else {
		for _, p := range path {
			attempt += fmt.Sprint(*p)
		}
	}

	if attempt != solution {
		t.Errorf("Solution is %v, but got %v", solution, attempt)
	}
}
