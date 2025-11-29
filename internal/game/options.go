package game

import (
	entityplayer "slahseronline/internal/entity/player"
)

type Option func(g *Game)

func WithPlayer(p *entityplayer.Player) Option {
	return func(g *Game) {
		g.player = p
	}
}

func WithScreenSettings(screenWidth int, screenHeight int) Option {
	return func(g *Game) {
		g.screen = &screen{screenWidth: screenWidth, screenHeight: screenHeight}
	}
}
