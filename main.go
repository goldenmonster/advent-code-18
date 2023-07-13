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
	return x < 0 || x > maxX || y < 0 || y > maxY || z < 0 || z > maxZ || cubic[x][y][z] == 0
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
