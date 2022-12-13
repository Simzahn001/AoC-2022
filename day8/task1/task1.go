package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var possibleTrees = 0

	//get the filescanner
	readfile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var grid = new(grid)
	grid.NewGrid(readfile)

	err = readfile.Close()
	if err != nil {
		return
	}

	//check if trees are visible
	trees := []tree{}
	for i := 1; i < len(grid.GetRow(0))-1; i++ {
		for j := 1; j < len(grid.GetColumn(0))-1; j++ {
			trees = append(trees, grid.GetTree(i, j))
		}
	}

	for _, tree := range trees {

		isVisible := false

		//left side
		if tree.height > findMax(grid.GetRow(tree.row)[0:tree.column]) {
			isVisible = true
		}

		//right side
		if tree.height > findMax(grid.GetRow(tree.row)[tree.column+1:]) {
			isVisible = true
		}

		//top
		if tree.height > findMax(grid.GetColumn(tree.column)[0:tree.row]) {
			isVisible = true
		}

		//bottom
		if tree.height > findMax(grid.GetColumn(tree.column)[tree.row:]) {
			isVisible = true
		}

		if !isVisible {
			possibleTrees++
		}

	}

	fmt.Println(grid.getSize() - possibleTrees)

}

type grid struct {
	grid [][]int
}

// NewGrid a constructor to generate a grid from the raw input
func (g *grid) NewGrid(input *os.File) {

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanLines)

	//put trees into array grid
	for filescanner.Scan() {
		row := []int{}

		for _, tree := range strings.Split(filescanner.Text(), "") {
			value, err := strconv.Atoi(tree)
			if err != nil {
				panic(err)
			}
			row = append(row, value)
		}
		g.grid = append(g.grid, row)
	}

}

// GetRow return a specific row with the given index
func (g *grid) GetRow(index int) []int {
	return g.grid[index]
}

// GetColumn returns a specific column with the given index
func (g *grid) GetColumn(index int) []int {
	column := []int{}

	for _, row := range g.grid {
		column = append(column, row[index])
	}

	return column
}

func (g *grid) GetTree(row int, column int) tree {
	return tree{
		height: g.grid[row][column],
		row:    row,
		column: column,
	}
}

// return size of the gird; the total amount of cells
func (g *grid) getSize() int {
	size := 0
	for i := 0; i < len(g.grid); i++ {
		size += len(g.grid[i])
	}
	return size
}

// IsBiggest returs, if the given int i is bigger than any part of the array

type tree struct {
	height int
	row    int
	column int
}

func findMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
