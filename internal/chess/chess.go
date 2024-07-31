package chess

import "github.com/Iyed-M/go-chess/internal/chess/state"

type game struct {
	state     *state.State
	whiteName string
	blackName string
}

type GameConfig struct {
	whiteName string
	blackName string
}

func startGame(cfg GameConfig) *game {
	g := &game{
		whiteName: "Player White",
		blackName: "Player Black",
	}
	if cfg.blackName != "" {
		g.blackName = cfg.blackName
	}

	if cfg.whiteName != "" {
		g.whiteName = cfg.whiteName
	}

	return g
}
