package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth    int32   = 1000
	WindowHeight   int32   = 600
	PixelsPerMeter float64 = 10
	MetersPerPixel float64 = 1.0 / PixelsPerMeter
)

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, "Raylib")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game := NewCircleGame("Circle Game", 4, true)

		previousTime := float64(rl.GetTime())
		dt := float64(0.0)

		for !rl.WindowShouldClose() {
			dt = float64(rl.GetTime()) - previousTime
			previousTime = float64(rl.GetTime())

			game.UpdatePhysics(dt)
			game.Draw()

			if rl.IsKeyPressed('R') {
				break
			}
		}
	}

	rl.CloseWindow()
}
