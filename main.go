package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth    int32   = 1000
	WindowHeight   int32   = 600
	PixelsPerMeter float32 = 10
	MetersPerPixel float32 = 1.0 / PixelsPerMeter
)

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, "Raylib")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game := NewCircleGame("Circle Game", 4, true)

		previousTime := float32(rl.GetTime())
		dt := float32(0.0)

		for !rl.WindowShouldClose() {
			dt = float32(rl.GetTime()) - previousTime
			previousTime = float32(rl.GetTime())

			game.UpdatePhysics(dt)
			game.Draw()

			if rl.IsKeyPressed('R') {
				break
			}
		}
	}

	rl.CloseWindow()
}
