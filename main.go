package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth  int32 = 1000
	WindowHeight int32 = 600
)

// Creates a raylib window and starts the main game loop
func main() {
	rl.InitWindow(WindowWidth, WindowHeight, "Raylib")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var game Game = createGame()
	var lastTickTime float64 = rl.GetTime()
	var deltaTime float64 = 0
	for !rl.WindowShouldClose() {

		// We can always get a new game instance to hot-reload
		if rl.IsKeyPressed('R') {
			game = createGame()
		}

		deltaTime = rl.GetTime() - lastTickTime
		lastTickTime = rl.GetTime()
		game.Update(deltaTime)
		game.Draw()
	}
}

func createGame() Game {
	return NewBoxesAndBallGame(30)
}
