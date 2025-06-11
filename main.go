package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

const (
	WindowWidth  int32 = 800
	WindowHeight int32 = 450
)

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, "Raylib - Physics2D")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game := NewCircleGame("Circle Game", 20, true, true)
		for !rl.WindowShouldClose() {
			game.UpdatePhysics()
			game.Draw()

			if rl.IsKeyPressed('R') {
				break
			}
		}
	}

	rl.CloseWindow()
}

func drawAABB(box p2d.AABB) {
	rl.DrawRectangleLines(int32(box.Min.X), int32(box.Min.Y), int32(box.Max.X)-int32(box.Min.X), int32(box.Max.Y)-int32(box.Min.Y), rl.Yellow)
}

func drawVec2(origin p2d.Vec2, v p2d.Vec2, c rl.Color) {
	rl.DrawLine(int32(origin.X), int32(origin.Y), int32(v.X), int32(v.Y), c)
}
