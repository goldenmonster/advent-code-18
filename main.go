package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

var cubic [][][]int
var maxX int
var maxY int
var maxZ int
var flag bool

func isWater(x, y, z int) bool {
	return x < 0 || x > maxX || y < 0 || y > maxY || z < 0 || z > maxZ || cubic[x][y][z] == 2
}

func isOpened(x, y, z int) bool {
	return x < 0 || x > maxX || y < 0 || y > maxY || z < 0 || z > maxZ || cubic[x][y][z] == 2
}

func checkAir(x int, y int, z int, visited [][][]int, visitedPositions []Position) {

	if cubic[x][y][z] != 0 {
		return
	}

	if isOpened(x-1, y, z) || isOpened(x+1, y, z) || isOpened(x, y-1, z) || isOpened(x, y+1, z) || isOpened(x, y, z-1) || isOpened(x, y, z+1) {
		cubic[x][y][z] = 2
		for _, pos := range visitedPositions {
			cubic[pos.x][pos.y][pos.z] = 2
		}
		return
	}

	isBackBlocked := cubic[x-1][y][z] == 1 || visited[x-1][y][z] == 1
	isFrontBlocked := cubic[x+1][y][z] == 1 || visited[x+1][y][z] == 1
	isLeftBlocked := cubic[x][y-1][z] == 1 || visited[x][y-1][z] == 1
	isRightBlocked := cubic[x][y+1][z] == 1 || visited[x][y+1][z] == 1
	isUpBlocked := cubic[x][y][z+1] == 1 || visited[x][y][z+1] == 1
	isDownBlocked := cubic[x][y][z-1] == 1 || visited[x][y][z-1] == 1

	if isBackBlocked && isFrontBlocked && isLeftBlocked && isRightBlocked && isUpBlocked && isDownBlocked {
		cubic[x][y][z] = 3
		for _, pos := range visitedPositions {
			cubic[pos.x][pos.y][pos.z] = 3
		}
		return
	}

	visited[x][y][z] = 1
	visitedPositions = append(visitedPositions, Position{x, y, z})

	if cubic[x-1][y][z] == 0 && visited[x-1][y][z] == 0 {
		checkAir(x-1, y, z, visited, visitedPositions)
	}

	if cubic[x+1][y][z] == 0 && visited[x+1][y][z] == 0 {
		checkAir(x+1, y, z, visited, visitedPositions)
	}

	if cubic[x][y+1][z] == 0 && visited[x][y+1][z] == 0 {
		checkAir(x, y+1, z, visited, visitedPositions)
	}

	if cubic[x][y-1][z] == 0 && visited[x][y-1][z] == 0 {
		checkAir(x, y-1, z, visited, visitedPositions)
	}

	if cubic[x][y][z+1] == 0 && visited[x][y][z+1] == 0 {
		checkAir(x, y, z+1, visited, visitedPositions)
	}

	if cubic[x][y][z-1] == 0 && visited[x][y][z-1] == 0 {
		checkAir(x, y, z-1, visited, visitedPositions)
	}
}

func main() {

	var lines []string

	file, err := os.Open("./in.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var lavas []Position

	for _, line := range lines {
		posArr := strings.Split(line, ",")
		x, _ := strconv.Atoi(posArr[0])
		y, _ := strconv.Atoi(posArr[1])
		z, _ := strconv.Atoi(posArr[2])

		lavas = append(lavas, Position{x, y, z})
		if maxX < x {
			maxX = x
		}

		if maxY < y {
			maxY = y
		}

		if maxZ < z {
			maxZ = z
		}
	}

	cubic = make([][][]int, maxX+1)
	visited := make([][][]int, maxX+1)

	for i := 0; i <= maxX; i++ {
		cubic[i] = make([][]int, maxY+1)
		visited[i] = make([][]int, maxY+1)

		for j := 0; j <= maxY; j++ {
			cubic[i][j] = make([]int, maxZ+1)
			visited[i][j] = make([]int, maxZ+1)
		}
	}

	for _, pos := range lavas {
		cubic[pos.x][pos.y][pos.z] = 1
	}

	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if cubic[i][j][0] != 1 {
				cubic[i][j][0] = 2
			}
			if cubic[i][j][maxZ] != 1 {
				cubic[i][j][maxZ] = 2
			}
		}
	}

	for i := 0; i <= maxX; i++ {
		for k := 0; k <= maxZ; k++ {
			if cubic[i][0][k] != 1 {
				cubic[i][0][k] = 2
			}

			if cubic[i][maxY][k] != 1 {
				cubic[i][maxY][k] = 2
			}
		}
	}

	for j := 0; j <= maxY; j++ {
		for k := 0; k <= maxZ; k++ {
			if cubic[0][j][k] != 1 {
				cubic[0][j][k] = 2
			}

			if cubic[maxX][j][k] != 1 {
				cubic[maxX][j][k] = 2
			}
		}
	}

	for i := 1; i < maxX; i++ {
		for j := 1; j < maxY; j++ {
			for k := 1; k < maxZ; k++ {
				checkAir(i, j, k, visited, []Position{})
			}
		}
	}

	cnt := 0

	for _, pos := range lavas {
		if isWater(pos.x-1, pos.y, pos.z) {
			cnt++
		}

		if isWater(pos.x+1, pos.y, pos.z) {
			cnt++
		}

		if isWater(pos.x, pos.y-1, pos.z) {
			cnt++
		}

		if isWater(pos.x, pos.y+1, pos.z) {
			cnt++
		}
		if isWater(pos.x, pos.y, pos.z-1) {
			cnt++
		}

		if isWater(pos.x, pos.y, pos.z+1) {
			cnt++
		}

	}

	fmt.Println(cnt)
}
