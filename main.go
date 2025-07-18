package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth  int32 = 1400
	WindowHeight int32 = 800
)

// Creates a raylib window and starts the main game loop
func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowHighdpi)
	rl.InitWindow(WindowWidth, WindowHeight, "Raylib")
	defer rl.CloseWindow()
	rl.SetTargetFPS(120)

	var sim Simulation = createSim()
	var lastTickTime float64 = rl.GetTime()
	var deltaTime float64 = 0
	for !rl.WindowShouldClose() {

		// We can always get a new game instance to hot-reload
		if rl.IsKeyPressed('R') {
			sim = createSim()
		}

		deltaTime = rl.GetTime() - lastTickTime
		lastTickTime = rl.GetTime()
		sim.Update(deltaTime)
		sim.Draw()
	}
}

func createSim() Simulation {
	return NewStackingSim()
}
