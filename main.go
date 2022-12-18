package main

import (
	"fmt"

	"github.com/NattapatN/maze-generator/generator"
)

func main() {
	fmt.Println("== Maze Generator ==")

	generator.Initial(40, 30)
	generator.PrintMaze()
}
