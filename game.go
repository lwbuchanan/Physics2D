package main

// This file holds all the code for various game setups, as well
// as any useful helper functions for interfacing between the
// physics library and raylib.

// The main file should just create one of these games and run it

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

const (
	PixelsPerMeter float64 = 10
	MetersPerPixel float64 = 1.0 / PixelsPerMeter
	worldWidth     float64 = float64(WindowWidth) * MetersPerPixel
	worldHeight    float64 = float64(WindowHeight) * MetersPerPixel
)

var (
	textColor       = rl.NewColor(255, 240, 124, 255)
	playerColor     = rl.NewColor(224, 225, 221, 255)
	objectColor     = rl.NewColor(119, 141, 169, 255)
	backgroundColor = rl.NewColor(13, 27, 42, 255)
)

type Game interface {
	Update(dt float64)
	Draw()
}

type CircleGame struct {
	Name         string
	physicsWorld p2d.World
	player       *p2d.Body
}

func NewCircleGame(name string, numCircles int, hasPlayer bool) CircleGame {

	circles := make([]*p2d.Body, numCircles)

	for i := range numCircles {
		ball, err := p2d.NewBall(
			p2d.NewVec2(
				rand.Float64()*(worldWidth-0),
				rand.Float64()*(worldHeight-0)),
			1+rand.Float64()*(3), 1, 5)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		circles[i] = ball
	}

	var player *p2d.Body = nil
	if hasPlayer {
		me, err := p2d.NewBall(p2d.NewVec2(worldWidth/2, worldHeight/2), 2, 1, 5)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			circles[0] = me
			player = me
		}
	}

	newGame := CircleGame{
		Name: name,
		physicsWorld: p2d.NewWorld(circles,
			p2d.NewVec2(
				worldWidth,
				worldHeight,
			),
			5),
		player: player,
	}

	rl.SetWindowTitle("Raylib - " + name)

	return newGame
}

func (g CircleGame) Update(dt float64) {

	if g.player != nil {
		mousePos := toP2dVec(rl.GetMousePosition())
		g.player.MoveTo(mousePos)
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

type BoxGame struct {
	physicsWorld p2d.World
	hasPlayer    bool
}

func NewBoxGame(numBoxes int, hasPlayer bool) BoxGame {
	boxes := make([]*p2d.Body, numBoxes)
	for i := range numBoxes {

		box, err := p2d.NewBox(
			getRandomPosition(),   // Position
			getRandomVector(2, 8), // Size
			// getRandomFloat(-math.Pi, math.Pi), // Rotational Velocity
			0,
			1, // Resitution
			5) // Mass
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		boxes[i] = box
	}

	newGame := BoxGame{
		p2d.NewWorld(boxes, p2d.NewVec2(worldWidth, worldHeight), 0),
		hasPlayer,
	}

	return newGame
}

func (g BoxGame) Update(dt float64) {

	if g.hasPlayer {
		mousePos := toP2dVec(rl.GetMousePosition())
		g.physicsWorld.Bodies[0].MoveTo(mousePos)
	}

	g.physicsWorld.UpdatePhysics(dt)

}

func (g BoxGame) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(backgroundColor)

	for i := range g.physicsWorld.Bodies {
		box := g.physicsWorld.Bodies[i]
		color := objectColor
		if i == 0 && g.hasPlayer {
			color = playerColor
		}
		err := drawConnectedVertices(box.Vertices(), 2, color)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	mestr := fmt.Sprintf("%.1f, %.1f", toP2dVec(rl.GetMousePosition()).X(), toP2dVec(rl.GetMousePosition()).Y())
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), WindowWidth-50, 10, 24, textColor)
	rl.DrawText(mestr, 10, 10, 24, textColor)
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

func getRandomPosition() p2d.Vec2 {
	return p2d.NewVec2(rand.Float64()*(worldWidth-0), rand.Float64()*(worldHeight-0))
}

func getRandomVector(min float64, max float64) p2d.Vec2 {
	return p2d.NewVec2(min+rand.Float64()*(max-min), min+rand.Float64()*(max-min))
}

func getRandomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func drawConnectedVertices(vertices []p2d.Vec2, thickness float32, color color.RGBA) error {
	numVertices := len(vertices)
	if numVertices < 2 {
		return errors.New("drawConnectedVertices(): not enough vertices")
	}
	for i := range numVertices {
		rl.DrawLineEx(
			toRLVec(vertices[i]),
			toRLVec(vertices[(i+1)%numVertices]),
			thickness,
			color)
	}
	return nil
}
