package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lwbuchanan/Physics2D/physics"
)

type Game struct {
	me  *physics.Circle
	obj []*physics.Circle
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	g.me.Move(physics.Vec2[float64]{X: float64(x), Y: float64(y)})
	return nil
}

func drawCircle(image *ebiten.Image, x_center int, y_center int, radius int) {
	x := 0
	y := radius
	d := 3 - 2*radius

	for x <= y {
		image.Set(x_center+x, y_center+y, color.White)
		image.Set(x_center-x, y_center+y, color.White)
		image.Set(x_center+x, y_center-y, color.White)
		image.Set(x_center-x, y_center-y, color.White)
		image.Set(x_center+y, y_center+x, color.White)
		image.Set(x_center-y, y_center+x, color.White)
		image.Set(x_center+y, y_center-x, color.White)
		image.Set(x_center-y, y_center-x, color.White)

		if d < 0 {
			d += 4*x + 6
		} else {
			d += 4*(x-y) + 10
			y -= 1
		}
		x += 1
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 20, 0xff})
	ebitenutil.DebugPrint(screen, "Hello, World!")

	var meCir *physics.Circle = g.me
	drawCircle(screen, int(meCir.Pos.X), int(meCir.Pos.Y), int(meCir.Rad))

	for i := 0; i < len(g.obj); i++ {
		var cir *physics.Circle = g.obj[i]
		drawCircle(screen, int(cir.Pos.X), int(cir.Pos.Y), int(cir.Rad))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	var c1 *physics.Circle = physics.NewCircle(physics.Vec2[float64]{X: 100., Y: 100.}, 25)
	var obj []*physics.Circle
	//obj = append(obj, c1)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{c1, obj}); err != nil {
		log.Fatal(err)
	}
}
