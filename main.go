package main

import (
	"fmt"

	"github.com/NattapatN/maze-generator/mazegenerator"
)

func main() {
	fmt.Println("== Maze Generator ==")

	mazegenerator.Initial(40, 30)
	mazegenerator.PrintMaze()
}
