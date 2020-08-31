package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Define logical representation for player
type player struct {
	image *ebiten.Image
	x, y  float64
	speed float64
}

// Set screen dimensions
const (
	screenWidth, screenHeight = 640, 480
)

// Define space, ship and player
var (
	err       error
	space     *ebiten.Image
	ship      *ebiten.Image
	playerOne player
)

// Initialize space, ship and player
func init() {
	space, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	ship, _, err = ebitenutil.NewImageFromFile("assets/ship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{ship, screenWidth / 2, screenHeight / 2, 4}
}

// Update player position if moved
func movePlayer() {
    if ebiten.IsKeyPressed(ebiten.KeyUp) {
        playerOne.y -= playerOne.speed
    }
    if ebiten.IsKeyPressed(ebiten.KeyDown) {
        playerOne.y += playerOne.speed
    }
    if ebiten.IsKeyPressed(ebiten.KeyLeft) {
        playerOne.x -= playerOne.speed
    }
    if ebiten.IsKeyPressed(ebiten.KeyRight) {
        playerOne.x += playerOne.speed
    }
}

// Update game state
func update(screen *ebiten.Image) error {
    movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	spaceOp := &ebiten.DrawImageOptions{}
	spaceOp.GeoM.Translate(0, 0)
	screen.DrawImage(space, spaceOp)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.x, playerOne.y)
    screen.DrawImage(playerOne.image, playerOp)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Ebiten Hello World"); err != nil {
		log.Fatal(err)
	}
}
