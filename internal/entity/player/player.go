package entityplayer

import (
	"log"
	"slahseronline/internal/asset"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var playerSprite *ebiten.Image

func New() *Player {
	playerSprite, _, err := ebitenutil.NewImageFromFileSystem(asset.AssetsFS, "assets/images/player/player.png")
	if err != nil {
		log.Fatal(err)
	}

	p := &Player{
		img: playerSprite,
	}
	return p
}

type Player struct {
	img *ebiten.Image
}
