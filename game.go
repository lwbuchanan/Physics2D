package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics2d"
)

type CircleGame struct {
	Name       string
	numCirlces int
	circles    []*p2d.Circle
	me         *p2d.Circle
}

func NewCircleGame(name string, numCircles int, varySize bool, hasPlayer bool) CircleGame {
	newGame := CircleGame{
		Name:       name,
		numCirlces: numCircles,
		circles:    make([]*p2d.Circle, numCircles),
		me:         nil,
	}

	for i := range numCircles {
		if varySize {
			newGame.circles[i] = p2d.NewCircle(p2d.Vec2{X: float32(rand.Int31n(WindowWidth)), Y: float32(rand.Int31n(WindowHeight))}, float32(rand.Int31n(30)+10))
		} else {
			newGame.circles[i] = p2d.NewCircle(p2d.Vec2{X: float32(rand.Int31n(WindowWidth)), Y: float32(rand.Int31n(WindowHeight))}, 20)
		}
	}
	if hasPlayer {
		newGame.me = p2d.NewCircle(p2d.Vec2{X: float32(WindowWidth / 2), Y: float32(WindowHeight / 2)}, 25)
		newGame.circles[0] = newGame.me
	}

	rl.SetWindowTitle(name)

	return newGame
}

func (g CircleGame) UpdatePhysics() {

	if g.me != nil {
		g.me.SetPos(p2d.Vec2(rl.GetMousePosition()))
	}

	for i := range len(g.circles) {
		for j := range len(g.circles) {
			if i == j {
				break
			}
			if collides, collision := g.circles[i].Collides(g.circles[j]); collides {
				collision.Resolve()
			}
		}
	}
}

func (g CircleGame) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	for i := range len(g.circles) {
		rl.DrawCircle(int32(g.circles[i].Pos.X), int32(g.circles[i].Pos.Y), float32(g.circles[i].Rad), rl.White)
	}

	if g.me != nil {
		rl.DrawCircle(int32(g.me.Pos.X), int32(g.me.Pos.Y), float32(g.me.Rad), rl.Blue)
	}

	rl.EndDrawing()
}
