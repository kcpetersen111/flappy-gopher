package main

import (
	"bytes"
	_ "embed"
	"flappyGopher/entity"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed src/gopher.jpg
var gophFile []byte

var goph *ebiten.Image

func init() {
	buf := bytes.NewBuffer(gophFile)
	im, _, err := image.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}

	goph = ebiten.NewImageFromImage(im)

}

type Game struct {
	player  entity.Player
	running bool
}

// Update proceeds the game state.
func (g *Game) Update() error {
	if !g.running {
		panic("Game over man")
	}

	g.player.Ypos += 3

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.player.Jump()
	}

	_, height := g.Layout(0, 0)

	if g.player.Ypos+(0.5*float64(goph.Bounds().Size().Y)) >= float64(height) {
		g.running = false
	}

	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// screen background
	screen.Fill(color.RGBA{
		R: 100,
		B: 255,
		G: 0,
		A: 255,
	})

	op := ebiten.DrawImageOptions{}
	//half image size, in real project dont scale image each draw
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(g.player.Xpos, g.player.Ypos)

	//draw character
	screen.DrawImage(goph, &op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 1024
}

func main() {
	game := &Game{
		player: entity.Player{
			Xpos: 100,
			Ypos: 100,
		},
		running: true,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Flappy Gopher")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
