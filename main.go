package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed src/gopher.jpg
var gophFile []byte

var gopherPic image.Image

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
	gopherPic = im
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
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
	goph := ebiten.NewImageFromImage(gopherPic)
	screen.DrawImage(goph, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Flappy Gopher")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
