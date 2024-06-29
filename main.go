package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 1024
const ScreenHeight = 720
const N = 20

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Game of Life")
	defer rl.CloseWindow()

	cellSize := int32(25)
	grid := make([][]int, N)
	prevGrid := make([][]int, N)

	for i := range grid {
		grid[i] = make([]int, N)
		prevGrid[i] = make([]int, N)

		for j := range grid[i] {
			if rand.Float32() > 0.9 {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
			prevGrid[i][j] = 0
		}
	}

	rl.SetTargetFPS(5)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.LightGray)

		for i := range grid {
			for j, cell := range grid[i] {
				var color rl.Color
				if cell == 1 {
					color = rl.DarkGray
				} else {
					color = rl.White
				}

				x := int32(i) * cellSize
				y := int32(j) * cellSize
				rl.DrawRectangle(x, y, cellSize, cellSize, color)

				textX := x + (cellSize - rl.MeasureText(fmt.Sprint(cell), 15)) / 2
				textY := y + (cellSize - 15) / 2
				rl.DrawText(fmt.Sprint(cell), textX, textY, 15, rl.Red)
			}
		}

		if isGameOver(&prevGrid, &grid) {
			color := rl.NewColor(80, 80, 80, 180)
			rl.DrawRectangle(0, 0, ScreenWidth, ScreenHeight, color)

			textX := 0 + (1024 - rl.MeasureText("Game over!", 15)) / 2
			textY := int32(0 + (720 - 15) / 2)
			rl.DrawText("Game over!", textX, textY, 15, rl.Red)
		} else {
			prevGrid = deepCopy(&grid)

			for i := range grid {
				for j, cell := range grid[i] {
					neighboursCount := cellNeighboursCount(&grid, i, j)
					if cell == 0 && neighboursCount == 3 {
						grid[i][j] = 1
					}

					if cell == 1 && (neighboursCount < 2 || neighboursCount > 3) {
						grid[i][j] = 0
					}
				}
			}
		}


		rl.DrawFPS(720, 0)

		rl.EndDrawing()
	}
}

func cellNeighboursCount(gridRef *[][]int, i, j int) int {
	grid := (*gridRef)

	return grid[mathRemainder(i - 1)][mathRemainder(j - 1)] +
		grid[mathRemainder(i - 1)][j] +
		grid[mathRemainder(i - 1)][mathRemainder(j + 1)] +
		grid[i][mathRemainder(j - 1)] +
		grid[i][mathRemainder(j + 1)] +
		grid[mathRemainder(i + 1)][mathRemainder(j - 1)] +
		grid[mathRemainder(i + 1)][j] +
		grid[mathRemainder(i + 1)][mathRemainder(j + 1)]
}

func mathRemainder(divident int) int {
	divisor := N
	return ((divident % divisor) + divisor) % divisor
}

func isGameOver(prevGrid, grid *[][]int) bool {
	isOver := true

	for _, row := range (*grid) {
		for _, cell := range row {
			if cell == 1 {
				isOver = false
				break
			}
		}
	}

	return isOver || areGridsEqual(prevGrid, grid)
}

func areGridsEqual(prevGrid, grid *[][]int) bool {
	isEqual := true

	for i, row := range (*grid) {
		for j, cell := range row {
			if cell != (*prevGrid)[i][j] {
				isEqual = false
				break
			}
		}
	}

	return isEqual
}

func deepCopy(grid *[][]int) [][]int {
	copy := make([][]int, len(*grid))

	for i, row := range (*grid) {
		copy[i] = make([]int, len(row))

		for j, cell := range row {
			copy[i][j] = cell
		}
	}

	return copy
}
