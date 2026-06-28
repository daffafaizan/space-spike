package main

import (
	"log"
	"math"

	"github.com/daffafaizan/space-spike/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition         Vector
	playerVelocity         Vector
	playerRotation         float64
	playerRotationVelocity float64
}

func (g *Game) Update() error {
	speed := float64(60 / ebiten.TPS())
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	g.playerVelocity.X += delta.X
	g.playerVelocity.Y += delta.Y

	g.playerPosition.X += g.playerVelocity.X
	g.playerPosition.Y += g.playerVelocity.Y

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.playerRotationVelocity -= 0.5
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.playerRotationVelocity += 0.5
	}

	g.playerRotation += g.playerRotationVelocity

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	width := assets.PlayerSprite.Bounds().Dx()
	height := assets.PlayerSprite.Bounds().Dx()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	if g.playerPosition.X != 0 && g.playerPosition.Y != 0 {
		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Rotate(g.playerRotation * math.Pi / 180.0)
		op.GeoM.Translate(halfW, halfH)
	}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(assets.PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerPosition:         Vector{X: 100, Y: 100},
		playerVelocity:         Vector{X: 0, Y: 0},
		playerRotation:         0,
		playerRotationVelocity: 0,
	}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
