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
	player *Player
}

type Player struct {
	playerPosition         Vector
	playerVelocity         Vector
	playerRotation         float64
	playerRotationVelocity float64
}

func NewPlayer() *Player {
	return &Player{
		playerPosition:         Vector{X: 0, Y: 0},
		playerVelocity:         Vector{X: 0, Y: 0},
		playerRotation:         0,
		playerRotationVelocity: 0,
	}
}

func (p *Player) Update() {
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

	p.playerVelocity.X += delta.X
	p.playerVelocity.Y += delta.Y

	p.playerPosition.X += p.playerVelocity.X
	p.playerPosition.Y += p.playerVelocity.Y

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		p.playerRotationVelocity -= 0.5
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		p.playerRotationVelocity += 0.5
	}

	p.playerRotation += p.playerRotationVelocity
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	width := assets.PlayerSprite.Bounds().Dx()
	height := assets.PlayerSprite.Bounds().Dx()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	if p.playerPosition.X != 0 && p.playerPosition.Y != 0 {
		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Rotate(p.playerRotation * math.Pi / 180.0)
		op.GeoM.Translate(halfW, halfH)
	}
	op.GeoM.Translate(p.playerPosition.X, p.playerPosition.Y)
	screen.DrawImage(assets.PlayerSprite, op)
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	p := &Player{}
	g := &Game{
		player: p,
	}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
