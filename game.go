package main

import (
	"fmt"
	"math/rand"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

const (
	worldWidth  float32 = float32(WindowWidth) * MetersPerPixel
	worldHeight float32 = float32(WindowHeight) * MetersPerPixel
)

type CircleGame struct {
	Name         string
	physicsWorld p2d.World
	player       *p2d.Circle
}

func NewCircleGame(name string, numCircles int, hasPlayer bool) CircleGame {

	circles := make([]*p2d.Circle, numCircles)

	for i := range numCircles {
		circles[i] = p2d.NewCircle(
			p2d.NewVec2(
				rand.Float32()*(worldWidth-0),
				rand.Float32()*(worldHeight-0)),
			1+rand.Float32()*(3), 5)
	}

	var me *p2d.Circle = nil
	if hasPlayer {
		me = p2d.NewCircle(p2d.NewVec2(worldWidth/2, worldHeight/2), 2, 5)
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

func (g CircleGame) UpdatePhysics(dt float32) {

	g.physicsWorld.UpdatePhysics(dt)

	if g.player != nil {
		mousePos := toP2dVec(rl.GetMousePosition())
		g.player.SetPos(mousePos)
	}
}

func (g CircleGame) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.NewColor(13, 27, 42, 255))

	for i := range len(g.physicsWorld.Circles) {
		circle := g.physicsWorld.Circles[i]
		rl.DrawCircleV(toRLVec(circle.Pos),
			circle.Rad*PixelsPerMeter,
			rl.NewColor(119, 141, 169, 255))
	}

	if g.player != nil {
		rl.DrawCircleV(toRLVec(g.player.Pos),
			g.player.Rad*PixelsPerMeter,
			rl.NewColor(224, 225, 221, 255))
	}

	mestr := fmt.Sprintf("%.1f, %.1f", g.player.Pos.X(), g.player.Pos.Y())
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), WindowWidth-50, 10, 24, rl.NewColor(255, 240, 124, 255))
	rl.DrawText(mestr, 10, 10, 24, rl.NewColor(255, 240, 124, 255))
	rl.EndDrawing()
}

func toRLVec(v p2d.Vec2) rl.Vector2 {
	return rl.Vector2{
		X: v.X() * PixelsPerMeter,
		Y: (worldHeight - v.Y()) * PixelsPerMeter,
	}
}

func toP2dVec(v rl.Vector2) p2d.Vec2 {
	return p2d.NewVec2(
		v.X*MetersPerPixel,
		(float32(WindowHeight)-v.Y)*MetersPerPixel,
	)
}
