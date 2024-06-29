package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const N = 20

func main() {
	rl.InitWindow(1024, 720, "Game of Life")
	defer rl.CloseWindow()

	cellSize := int32(25)
	grid := make([][]int, N)

	for i := range grid {
		grid[i] = make([]int, N)

		for j := range grid[i] {
			if rand.Float32() > 0.3 {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
		}
	}

	rl.SetTargetFPS(10)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.LightGray)

		if isGameOver(&grid) {
			textX := 0 + (1024 - rl.MeasureText("Game over!", 15)) / 2
			textY := int32(0 + (720 - 15) / 2)
			rl.DrawText("Game over!", textX, textY, 15, rl.Red)
		} else {
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
		}

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

		rl.DrawFPS(720, 0)

		rl.EndDrawing()
	}
}

func cellNeighboursCount(gridRef *[][]int, i, j int) int {
	count := 0
	grid := (*gridRef)

	// above left
	if i > 0 && j > 0 && grid[i - 1][j - 1] == 1 {
		count += 1
	}
	// above
	if i > 0 && grid[i - 1][j] == 1 {
		count += 1
	}
	// above right
	if i > 0 && j < N - 1 && grid[i - 1][j + 1] == 1 {
		count += 1
	}

	// left
	if j > 0 && grid[i][j - 1] == 1 {
		count += 1
	}
	//right
	if j < N - 1 && grid[i][j + 1] == 1 {
		count += 1
	}

	// below left
	if i < N - 1 && j > 0 && grid[i + 1][j - 1] == 1 {
		count += 1
	}
	// below
	if i < N - 1 && grid[i + 1][j] == 1 {
		count += 1
	}
	// below right
	if i < N - 1 && j < N - 1 && grid[i + 1][j + 1] == 1 {
		count += 1
	}


	return count
}

func isGameOver(grid *[][]int) bool {
	isOver := true

	for _, row := range (*grid) {
		for _, cell := range row {
			if cell == 1 {
				isOver = false
				break
			}
		}
	}

	return isOver
}
