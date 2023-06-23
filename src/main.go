package main

import (
	"fmt"

	recursion "github.com/deankinane/algo-course/src/recursion/maze_solver"
)

func main() {
	maze := []string{
		"########E#",
		"# #####  #",
		"# #     ##",
		"#   ######",
		"##    ## #",
		"####S    #",
	}

	start := recursion.Point{X: 4, Y: 5}

	if path := recursion.SolveMaze(&maze, &start); path != nil {
		for i := 0; i < len(maze); i++ {
			fmt.Println(maze[i])
		}
	} else {
		fmt.Print("Failed to solve the maze")
	}
}
