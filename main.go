package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	p2d "github.com/lwbuchanan/Physics2D/physics"
)

func main() {
	rl.InitWindow(800, 450, "Physics2D")
	defer rl.CloseWindow()

	circle := p2d.NewCircle(p2d.Vec2{X: 400, Y: 225}, 20)
	circle2 := p2d.NewCircle(p2d.Vec2{X: 200, Y: 100}, 20)
	status := ""
	c1status := ""
	c2status := ""

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		circle.Move(p2d.Vec2(rl.GetMousePosition()))

		if collides, collision := circle.Collides(circle2); collides {
			status = fmt.Sprintf("Normal: %.2v, Depth: %.2f", collision.Normal, collision.Depth)
		} else {
			status = "nothing going on"
		}

		c1status = fmt.Sprintf("C1: %d, %d", int32(circle.Pos.X), int32(circle.Pos.Y))
		c2status = fmt.Sprintf("C2: %d, %d", int32(circle2.Pos.X), int32(circle2.Pos.Y))

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawText(status, 10, 10, 20, rl.White)
		rl.DrawText(c1status, 10, 30, 20, rl.White)
		rl.DrawText(c2status, 10, 50, 20, rl.White)

		rl.DrawCircle(int32(circle2.Pos.X), int32(circle2.Pos.Y), float32(circle2.Rad), rl.White)
		rl.DrawCircle(int32(circle.Pos.X), int32(circle.Pos.Y), float32(circle.Rad), rl.Blue)
		// drawAABB(circle2.Box)
		// drawAABB(circle.Box)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func drawAABB(box p2d.AABB) {
	rl.DrawRectangleLines(int32(box.Min.X), int32(box.Min.Y), int32(box.Max.X)-int32(box.Min.X), int32(box.Max.Y)-int32(box.Min.Y), rl.Yellow)
	drawVec2(box.Min, rl.Blue)
	drawVec2(box.Max, rl.Red)
}

func drawVec2(v p2d.Vec2, c rl.Color) {
	rl.DrawLine(0, 0, int32(v.X), int32(v.Y), c)
}
