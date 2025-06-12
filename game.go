package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

const (
	worldWidth  float64 = float64(WindowWidth) * MetersPerPixel
	worldHeight float64 = float64(WindowHeight) * MetersPerPixel
)

type CircleGame struct {
	Name         string
	physicsWorld p2d.World
	player       *p2d.Body
}

func NewCircleGame(name string, numCircles int, hasPlayer bool) CircleGame {

	circles := make([]*p2d.Body, numCircles)

	for i := range numCircles {
		circles[i] = p2d.NewBall(
			p2d.NewVec2(
				rand.Float64()*(worldWidth-0),
				rand.Float64()*(worldHeight-0)),
			1+rand.Float64()*(3), 1, 5)
	}

	var me *p2d.Body = nil
	if hasPlayer {
		me = p2d.NewBall(p2d.NewVec2(worldWidth/2, worldHeight/2), 2, 1, 5)
		circles[0] = me
	}

	newGame := CircleGame{
		Name: name,
		physicsWorld: p2d.NewWorld(numCircles, circles,
			p2d.NewVec2(
				worldWidth,
				worldHeight,
			)),
		player: me,
	}

	rl.SetWindowTitle("Raylib - " + name)

	return newGame
}

func (g CircleGame) UpdatePhysics(dt float64) {

	if g.player != nil {
		mousePos := toP2dVec(rl.GetMousePosition())
		g.player.Position = (mousePos)
	}

	g.physicsWorld.UpdatePhysics(dt)

}

func (g CircleGame) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.NewColor(13, 27, 42, 255))

	for i := range len(g.physicsWorld.Bodies) {
		circle := g.physicsWorld.Bodies[i]
		rl.DrawCircleV(toRLVec(circle.Position),
			float32(circle.Radius*PixelsPerMeter),
			rl.NewColor(119, 141, 169, 255))
	}

	if g.player != nil {
		rl.DrawCircleV(toRLVec(g.player.Position),
			float32(g.player.Radius*PixelsPerMeter),
			rl.NewColor(224, 225, 221, 255))
	}

	mestr := fmt.Sprintf("%.1f, %.1f", g.player.Position.X(), g.player.Position.Y())
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), WindowWidth-50, 10, 24, rl.NewColor(255, 240, 124, 255))
	rl.DrawText(mestr, 10, 10, 24, rl.NewColor(255, 240, 124, 255))
	rl.EndDrawing()
}

func toRLVec(v p2d.Vec2) rl.Vector2 {
	return rl.Vector2{
		X: float32(v.X() * PixelsPerMeter),
		Y: float32((worldHeight - v.Y()) * PixelsPerMeter),
	}
}

func toP2dVec(v rl.Vector2) p2d.Vec2 {
	return p2d.NewVec2(
		float64(v.X)*MetersPerPixel,
		(float64(WindowHeight)-float64(v.Y))*MetersPerPixel,
	)
}
