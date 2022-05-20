package main

type Player struct {
	hash string
	name string
}

func addPlayerToGame(game *Game, player Player) {
	game.players = append(game.players, player)
}
