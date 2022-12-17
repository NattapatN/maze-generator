package main

import (
	"fmt"
	"math/rand"
	"time"
)

var blockY int
var blockX int
var Maze [][][4]int
var wall map[string]int
var visitedMaze [][]bool
var path [][2]int

func Initial(width, hieght int) {
	blockY = hieght
	blockX = width
	wall = make(map[string]int)
	wall["up"] = 0
	wall["left"] = 1
	wall["right"] = 2
	wall["down"] = 3

	initMaze()
	mazeGenerator(0, 0)
}

func GetMaze() [][][4]int {
	return Maze
}

func mazeGenerator(startX, startY int) {
	currentPostion := [2]int{startX, startY}
	path = append(path, currentPostion)
	Maze[0][0][wall["up"]] = 0
	for len(path) > 0 {
		visitedMaze[currentPostion[0]][currentPostion[1]] = true
		availablePath := getAvailableNeighbor(currentPostion[0], currentPostion[1])
		if len(availablePath) == 0 {
			path = path[:len(path)-1]
		} else {
			randomPostion := rand.Intn(len(availablePath))
			path = append(path, availablePath[randomPostion])
			removeWall(currentPostion, availablePath[randomPostion])
		}
		if len(path) > 0 {
			currentPostion = path[len(path)-1]
		}
	}
	Maze[blockY-1][blockX-1][wall["down"]] = 0
	visitedMaze[blockY-1][blockX-1] = true
}

func getAvailableNeighbor(x, y int) [][2]int {
	result := [][2]int{}
	// check left
	if x > 0 && !visitedMaze[x-1][y] {
		result = append(result, [2]int{x - 1, y})
	}
	//check rigth
	if x < blockY-1 && !visitedMaze[x+1][y] {
		result = append(result, [2]int{x + 1, y})
	}
	//check up
	if y > 0 && !visitedMaze[x][y-1] {
		result = append(result, [2]int{x, y - 1})
	}
	//check down
	if y < blockX-1 && !visitedMaze[x][y+1] {
		result = append(result, [2]int{x, y + 1})
	}
	return result
}

func initMaze() {
	for i := 0; i < blockY; i++ {
		Maze = append(Maze, [][4]int{})
		visitedMaze = append(visitedMaze, []bool{})
		for j := 0; j < blockX; j++ {
			Maze[i] = append(Maze[i], [4]int{1, 1, 1, 1})
			visitedMaze[i] = append(visitedMaze[i], false)
		}
	}
}

func removeWall(currentLocation, nextLocation [2]int) {
	//check next postopn
	diffX := nextLocation[0] - currentLocation[0]
	diffY := nextLocation[1] - currentLocation[1]
	switch {
	case diffY == 1:
		Maze[currentLocation[0]][currentLocation[1]][wall["right"]] = 0
		Maze[nextLocation[0]][nextLocation[1]][wall["left"]] = 0
	case diffY == -1:
		Maze[currentLocation[0]][currentLocation[1]][wall["left"]] = 0
		Maze[nextLocation[0]][nextLocation[1]][wall["right"]] = 0
	case diffX == -1:
		Maze[currentLocation[0]][currentLocation[1]][wall["up"]] = 0
		Maze[nextLocation[0]][nextLocation[1]][wall["down"]] = 0
	case diffX == 1:
		Maze[currentLocation[0]][currentLocation[1]][wall["down"]] = 0
		Maze[nextLocation[0]][nextLocation[1]][wall["up"]] = 0
	}
}

func PrintMaze() {
	time.Sleep(time.Second / 10)
	for i := 0; i < blockX; i++ {
		if i == 0 {
			fmt.Print("   ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
	for _, value := range Maze {
		for _, vj := range value {
			if vj[1] == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			if vj[3] == 1 {
				fmt.Print("_")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}
