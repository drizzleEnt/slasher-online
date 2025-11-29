package game

import (
	"fmt"
	"log"
	entityplayer "slahseronline/internal/entity/player"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	maxAngle = 256
	maxLean  = 16
)

var _ ebiten.Game = (*Game)(nil)

func New(opts ...Option) *Game {

	g := &Game{}

	for _, opt := range opts {
		opt(g)
	}

	if g.screen == nil {
		g.screen = &screen{
			screenWidth:  800,
			screenHeight: 500,
		}
	}

	if g.player == nil {
		log.Fatalf("player not inited")
	}

	ebiten.SetWindowSize(g.screen.screenWidth, g.screen.screenHeight)
	ebiten.SetWindowTitle("Slasher Online")

	return g
}

type screen struct {
	screenWidth  int
	screenHeight int
}

type Game struct {
	screen *screen
	player *entityplayer.Player
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	
	return
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return g.screen.screenWidth, g.screen.screenHeight
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("pressed space")
	}
	return nil
}
