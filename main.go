package main

import (
	"bytes"
	_ "embed"
	"flappyGopher/player"
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
	// put this outside of draw loop
	// raw, err := os.ReadFile("flappyGopher/src/gopher.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	buf := bytes.NewBuffer(gophFile)
	im, _, err := image.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}

	goph = ebiten.NewImageFromImage(im)

}

// Game implements ebiten.Game interface.
type Game struct {
	player player.Player
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if g.player.Direction {
		g.player.Ypos += 3
	} else {
		g.player.Ypos -= 3
	}

	spaceDown := ebiten.IsKeyPressed(ebiten.KeySpace)
	if spaceDown {
		g.player.Jump()
	}

	_, height := g.Layout(0, 0)

	if g.player.Ypos+(0.5*float64(goph.Bounds().Size().Y)) >= float64(height) {
		g.player.Direction = false
	} else if g.player.Ypos <= 0 {
		g.player.Direction = true
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
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
	screen.DrawImage(goph, &op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 1024
}

func main() {
	game := &Game{
		player: player.Player{
			Xpos: 100,
			Ypos: 100,
		},
	}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Flappy Gopher")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
