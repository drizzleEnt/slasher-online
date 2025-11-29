package main

import (
	"log"
	entityplayer "slahseronline/internal/entity/player"
	"slahseronline/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	player := entityplayer.New()

	game := game.New(
		game.WithScreenSettings(800, 500),
		game.WithPlayer(player),
	)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
