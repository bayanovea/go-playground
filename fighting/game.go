package main

import (
	"math/rand"
	"time"
)

type Game struct {
	status  string
	name    string
	players []Player
}

var games []Game

func addNewGame(name string) *Game {
	var newGame = Game{name: name, status: "new"}
	games = append(games, newGame)

	return &newGame
}

func playGame(game *Game, player1Cards *[5]Card, player2Cards *[5]Card, ch chan string) {
	game.status = "active"

	var whoseTurn = "player1"
	var turnNumber int = 1

	for {
		println("=====")
		println("Game", "("+game.name+")", ",", "Turn", turnNumber)
		println("=====")

		var attackCard *Card
		var defenseCard *Card

		prepareToFight()

		if whoseTurn == "player1" {
			attackCard = findRandomAliveCard(player1Cards)
			defenseCard = findRandomAliveCard(player2Cards)
		} else if whoseTurn == "player2" {
			attackCard = findRandomAliveCard(player2Cards)
			defenseCard = findRandomAliveCard(player1Cards)
		}

		attack(attackCard, defenseCard)

		if whoseTurn == "player1" {
			whoseTurn = "player2"
		} else if whoseTurn == "player2" {
			whoseTurn = "player1"
		}

		printBoard(*player1Cards, *player2Cards)

		var canFinishGame bool
		var winner string
		canFinishGame, winner = handleAfterTurn(player1Cards, player2Cards)

		if canFinishGame {
			game.status = "finished"
			println(winner, " won!")
			break
		}

		turnNumber++
	}
}

func attack(offensiveCard *Card, defensiveCard *Card) {
	defensiveCard.health -= offensiveCard.attack
	offensiveCard.health -= defensiveCard.attack

	if offensiveCard.health <= 0 {
		offensiveCard.status = "dead"
	}
	if defensiveCard.health <= 0 {
		defensiveCard.status = "dead"
	}
}

func handleAfterTurn(player1Cards *[5]Card, player2Cards *[5]Card) (canFinishGame bool, winner string) {
	var canFinishGameVar bool = false
	var winnerVar string

	var isAllPlayer1CardsDead bool = isAllCardsDead(player1Cards)
	var isAllPlayer2CardsDead bool = isAllCardsDead(player2Cards)

	if isAllPlayer1CardsDead || isAllPlayer2CardsDead {
		canFinishGameVar = true

		if isAllPlayer1CardsDead {
			winnerVar = "Player 2"
		}
		if isAllPlayer2CardsDead {
			winnerVar = "Player 1"
		}
	}

	return canFinishGameVar, winnerVar
}

func prepareToFight() {
	println("Troops are preparing to fight! ...")

	var randomSleepDurationInSec = rand.Intn(3) + 1
	time.Sleep(time.Duration(randomSleepDurationInSec) * time.Second)
}

func getActiveGames() []Game {
	var activeGames []Game

	for _, game := range games {
		if game.status == "active" {
			activeGames = append(activeGames, game)
		}
	}

	return activeGames
}
