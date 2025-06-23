package main

// This file holds all the code for various game setups, as well
// as any useful helper functions for interfacing between the
// physics library and raylib.

// The main file should just create one of these games and run it

import (
	"errors"
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

const (
	PixelsPerMeter float64 = 200 // 1000 px world is 5 meters accross
	MetersPerPixel float64 = 1.0 / PixelsPerMeter
	worldWidth     float64 = float64(WindowWidth) * MetersPerPixel
	worldHeight    float64 = float64(WindowHeight) * MetersPerPixel
)

var (
	textColor       = rl.NewColor(255, 240, 124, 255)
	playerColor     = rl.NewColor(224, 225, 221, 255)
	objectColor     = rl.NewColor(119, 181, 169, 255)
	backgroundColor = rl.NewColor(13, 27, 42, 255)
)

type Game interface {
	Update(dt float64)
	Draw()
}

type BoxesAndBallsGame struct {
	physicsWorld p2d.World
}

func NewParticleGame(particles int) BoxesAndBallsGame {
	bodies := make([]*p2d.Body, particles)
	for i := range particles {
		p, _ := p2d.NewBall(getRandomPosition(), 0.05, 1, 0.1)
		p.ApplyForce(getRandomVector(-1, 1))
		bodies[i] = p
	}
	floor, _ := p2d.NewBox(p2d.NewVec2(worldWidth/2, 0.025), p2d.NewVec2(worldWidth, 0.05), 0, 1, 0)
	ceil, _ := p2d.NewBox(p2d.NewVec2(worldWidth/2, worldHeight-0.025), p2d.NewVec2(worldWidth, 0.05), 0, 1, 0)
	wall1, _ := p2d.NewBox(p2d.NewVec2(0.025, worldHeight/2), p2d.NewVec2(0.05, worldHeight), 0, 1, 0)
	wall2, _ := p2d.NewBox(p2d.NewVec2(worldWidth-0.025, worldHeight/2), p2d.NewVec2(0.05, worldHeight), 0, 1, 0)
	bodies = append(bodies, floor)
	bodies = append(bodies, ceil)
	bodies = append(bodies, wall1)
	bodies = append(bodies, wall2)

	game := BoxesAndBallsGame{
		p2d.NewWorld(bodies, p2d.NewVec2(worldWidth, worldHeight), 0, 1),
	}

	return game
}

func NewPegGame() BoxesAndBallsGame {
	var bodies []*p2d.Body
	player, _ := p2d.NewBall(p2d.NewVec2(getRandomFloat(0.5, worldWidth-0.5), worldHeight-0.2), 0.1, 1, 0.1)
	bodies = append(bodies, player)
	floor, _ := p2d.NewBox(p2d.NewVec2(worldWidth/2, 0.05), p2d.NewVec2(worldWidth, 0.1), 0, 0.2, 0)
	bodies = append(bodies, floor)
	wall1, _ := p2d.NewBox(p2d.NewVec2(0.3, worldHeight/2+0.1), p2d.NewVec2(0.03, 3.5), 0.11, 1, 0)
	bodies = append(bodies, wall1)
	wall2, _ := p2d.NewBox(p2d.NewVec2(worldWidth-0.3, worldHeight/2+0.1), p2d.NewVec2(0.03, 3.5), -0.11, 1, 0)
	bodies = append(bodies, wall2)

	slots := 8
	sep := 2.5 / float64(slots)
	for i := range slots + 1 {
		slot, _ := p2d.NewBox(p2d.NewVec2(0.5+(sep*float64(i)), 0.25), p2d.NewVec2(0.03, 0.3), 0, 0.6, 0)
		bodies = append(bodies, slot)
		print(i)
	}

	cols := 6
	rows := 8
	xsep := 2.8 / float64(cols)
	ysep := 3 / float64(rows)
	for j := range rows {
		for i := range cols - (j+1)%2 {
			peg, _ := p2d.NewBall(p2d.NewVec2(0.6+(xsep*float64(i)+(xsep/2.0*float64((j+1)%2))), 0.8+(ysep*float64(j))), 0.03, 0.7, 0)
			bodies = append(bodies, peg)
		}
	}

	newGame := BoxesAndBallsGame{
		p2d.NewWorld(bodies, p2d.NewVec2(worldWidth, worldHeight), 2, 1),
	}

	rl.SetWindowTitle("Let's go gambling!")
	return newGame
}


func NewBoxesAndBallGame(numBodies int) BoxesAndBallsGame {
	bodies := make([]*p2d.Body, numBodies)
	for i := range numBodies {
		var err error
		var body *p2d.Body
		mass := 0.5
		if rand.IntN(2) == 0 && i > 0 {
			mass = 0
		}
		if rand.IntN(2) == 0 {
			body, err = p2d.NewBox(
				getRandomPosition(),               // Position
				getRandomVector(.3, .5),           // Size
				getRandomFloat(-math.Pi, math.Pi), // Rotation
				1,                                 // Resitution
				mass,                              // Mass
			)
		} else {
			body, err = p2d.NewBall(
				getRandomPosition(),       // Position
				getRandomFloat(0.2, 0.25), // Size
				1,                         // Resitution
				mass,                      // Mass
			)
		}

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		bodies[i] = body
	}

	newGame := BoxesAndBallsGame{
		p2d.NewWorld(bodies, p2d.NewVec2(worldWidth, worldHeight), 0, 1),
	}

	rl.SetWindowTitle("Raylib - Boxs and Ball")

	return newGame
}

func (g BoxesAndBallsGame) Update(dt float64) {
	p := g.physicsWorld.Bodies[0]
	if rl.IsKeyDown(rl.KeyA) {
		p.ApplyForce(p2d.NewVec2(-2, 0))
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.ApplyForce(p2d.NewVec2(2, 0))
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.ApplyForce(p2d.NewVec2(0, -2))
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.ApplyForce(p2d.NewVec2(0, 2))
	}
	if rl.IsKeyDown(rl.KeyQ) {
		p.ApplyTorque(0.5)
	}
	if rl.IsKeyDown(rl.KeyE) {
		p.ApplyTorque(-0.5)
	}

	g.physicsWorld.UpdatePhysics(dt)

}

func (g BoxesAndBallsGame) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(backgroundColor)

	for i, body := range g.physicsWorld.Bodies {
		color := objectColor
		if body.Mass() == 0 {
			color = rl.Red
		}
		if body.Shape() == p2d.Polygon {
			err := drawPolygon(body.Vertices(), color)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else if body.Shape() == p2d.Ball {
			rl.DrawCircleV(toRLVec(body.Position()),
				float32(body.Radius()*PixelsPerMeter),
				color)
		}
		if i == 0 {
			rl.DrawCircleV(toRLVec(body.Position()), 4, playerColor)
		}
		// showVector := body.Velocity()
		// rl.DrawLineEx(toRLVec(body.Position()), toRLVec(body.Position().Add(showVector)), 2, textColor)
	}

	mestr := fmt.Sprintf("%.1f, %.1f", toP2dVec(rl.GetMousePosition()).X(), toP2dVec(rl.GetMousePosition()).Y())
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), WindowWidth-50, 10, 24, textColor)
	rl.DrawText(mestr, 10, 10, 24, textColor)
	rl.EndDrawing()
}

type StackingGame struct {
	physicsWorld *p2d.World
	colors []color.RGBA
}

func NewStackingGame() *StackingGame {
	var bodies []*p2d.Body
	var colors []color.RGBA
	floor, _ := p2d.NewBox(p2d.NewVec2(worldWidth/2, 0.15), p2d.NewVec2(worldWidth - 0.15, 0.15), 0, 0.5, 0)
	bodies = append(bodies, floor)
	colors = append(colors, rl.Gray)
	

	world := p2d.NewWorld(bodies, p2d.NewVec2(worldWidth, worldHeight), 9.8, 100)
	newGame := StackingGame{
		&world,
		colors,
	}
	return &newGame
}

func (g* StackingGame) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)

	for i, body := range g.physicsWorld.Bodies {
		color := g.colors[i]
		if body.Shape() == p2d.Polygon {
			err := drawPolygon(body.Vertices(), color)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else if body.Shape() == p2d.Ball {
			rl.DrawCircleV(toRLVec(body.Position()),
				float32(body.Radius()*PixelsPerMeter),
				color)
		}
	}
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), WindowWidth-50, 10, 24, textColor)
	rl.EndDrawing()
}

func (g* StackingGame) Update(dt float64) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		newBox, err := p2d.NewBox(
			toP2dVec(rl.GetMousePosition()), 
			getRandomVector(0.3, 0.4),
			0,
			0.5,
			1,
		)
		if err != nil {
			fmt.Println(err.Error())
			panic(-1)
		}
		g.physicsWorld.AddBody(newBox)
		g.colors = append(g.colors, rl.SkyBlue)
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		newBall, err := p2d.NewBall(
			toP2dVec(rl.GetMousePosition()), 
			getRandomFloat(0.15, 0.2),
			0.5,
			1,
		)
		if err != nil {
			println(err.Error())
			panic(-1)
		}
		g.physicsWorld.AddBody(newBall)
		g.colors = append(g.colors, rl.Yellow)
	}
	g.physicsWorld.UpdatePhysics(dt)
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

//func drawConnectedVertices(vertices []p2d.Vec2, thickness float32, color color.RGBA) error {
//	numVertices := len(vertices)
//	if numVertices < 2 {
//		return errors.New("drawConnectedVertices: not enough vertices")
//	}
//	for i := range numVertices {
//		rl.DrawLineEx(
//			toRLVec(vertices[i]),
//			toRLVec(vertices[(i+1)%numVertices]),
//			thickness,
//			color)
//	}
//	return nil
//}

func drawPolygon(vertices []p2d.Vec2, color color.RGBA) error {
	numVertices := len(vertices)
	if numVertices < 3 {
		return errors.New("drawPolygon: not enough vertices")
	}
	rlPoints := make([]rl.Vector2, numVertices)
	for i, v := range vertices {
		// raylib expects counter-clockwise vertex list
		rlPoints[numVertices-1-i] = toRLVec(v)
	}
	rl.DrawTriangleFan(rlPoints, color)
	return nil
}
