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
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

const (
	PixelsPerMeter float64 = 200 // 1000 px world is 5 meters accross
	MetersPerPixel float64 = 1.0 / PixelsPerMeter
	worldWidth     float64 = float64(WindowWidth) * MetersPerPixel
	worldHeight    float64 = float64(WindowHeight) * MetersPerPixel
)

type Simulation interface {
	Update(dt float64)
	Draw()
}

// The game core is the core of every simulation
type GameCore struct {
	physicsWorld    *p2d.World
	textColor       color.RGBA
	backgroundColor color.RGBA
	colors          []color.RGBA
	debugMode       bool
	numBodies       int
	elapsedSteps    int
	stopwatchStart  float64
	avgStepTime     float64
}

func (c *GameCore) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(c.backgroundColor)

	for i, body := range c.physicsWorld.Bodies {
		color := c.colors[i]
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

	if c.debugMode {
		performanceString := fmt.Sprintf(
			"Step Time: %f s\nBodies: %d\nFPS: %d",
			c.avgStepTime,
			c.numBodies,
			rl.GetFPS())
		rl.DrawText(performanceString, 10, 10, 20, c.textColor)
	}

	if c.physicsWorld.Paused {
		rl.DrawText("PAUSED", (WindowWidth/2)-65, 10, 30, rl.Red)
		rl.DrawRectangleLinesEx(rl.NewRectangle(0, 0, float32(WindowWidth), float32(WindowHeight)), 5, rl.Red)
	}
	rl.EndDrawing()
}

func (c *GameCore) Update(dt float64) {
	if rl.IsKeyPressed(rl.KeySpace) {
		c.physicsWorld.Paused = !c.physicsWorld.Paused
	}
	c.physicsWorld.UpdatePhysics(dt)

	if c.debugMode {
		c.elapsedSteps += c.physicsWorld.NumSteps()

		currTime := rl.GetTime()
		elapsedTime := currTime - c.stopwatchStart
		if elapsedTime > 0.2 {
			c.avgStepTime = elapsedTime / float64(c.elapsedSteps)

			c.elapsedSteps = 0
			c.stopwatchStart = rl.GetTime()
			c.numBodies = len(c.physicsWorld.Bodies)
		}
	}

	deleteSet := map[int]bool{}
	for i, b := range c.physicsWorld.Bodies {
		if b.Position().Y() < -5 {
			deleteSet[i] = true
		}
	}
	for i := range deleteSet {
		c.physicsWorld.Bodies = slices.Delete(c.physicsWorld.Bodies, i, i+1)
		c.colors = slices.Delete(c.colors, i, i+1)
	}
}

type StackingSim struct {
	GameCore
}

func (s *StackingSim) Update(dt float64) {
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
		s.physicsWorld.AddBody(newBox)
		s.colors = append(s.colors, rl.SkyBlue)
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
		s.physicsWorld.AddBody(newBall)
		s.colors = append(s.colors, rl.Yellow)
	}
	s.GameCore.Update(dt)
}

func NewStackingSim() *StackingSim {
	var bodies []*p2d.Body
	var colors []color.RGBA
	floor, _ := p2d.NewBox(p2d.NewVec2(worldWidth/2, 0.15), p2d.NewVec2(worldWidth-0.15, 0.15), 0, 0.5, 0)
	bodies = append(bodies, floor)
	colors = append(colors, rl.Gray)

	world := p2d.NewWorld(bodies, p2d.NewVec2(worldWidth, worldHeight), 9.8, 50)

	return &StackingSim{
		GameCore{
			&world,
			rl.NewColor(255, 240, 124, 255),
			rl.NewColor(13, 27, 42, 255),
			colors,
			//debug stuff
			true,
			0,
			0,
			rl.GetTime(),
			1.0,
		},
	}
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

func drawVectorArrow(origin, vector p2d.Vec2) {
	rl.DrawLineEx(toRLVec(origin), toRLVec(vector), 2, rl.Red)
	rl.DrawCircleV(toRLVec(vector), 4, rl.Red)
}
