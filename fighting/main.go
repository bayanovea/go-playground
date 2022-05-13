package main

import "math/rand"

type Card struct {
	attack int
	health int
	status string
}

func main() {
	var player1Cards [5]Card
	var player2Cards [5]Card

	fillCards(&player1Cards)
	fillCards(&player2Cards)

	printBoard(player1Cards, player2Cards)

	play(&player1Cards, &player2Cards)
}

func fillCards(cards *[5]Card) {
	for i := 0; i < len(cards); i++ {
		cards[i].attack = rand.Intn(10) + 1
		cards[i].health = rand.Intn(10) + 1
		cards[i].status = "alive"
	}
}

func play(player1Cards *[5]Card, player2Cards *[5]Card) {
	var whoseTurn = "player1"
	var turnNumber int = 1

	for {
		println("=====")
		println("Turn ", turnNumber)
		println("=====")

		var attackCard *Card
		var defenseCard *Card

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
			println(winner, " won!")
			break
		}

		turnNumber++
	}
}

func isAllCardsDead(cards *[5]Card) bool {
	for i := 0; i < len(cards); i++ {
		if cards[i].status == "alive" {
			return false
		}
	}

	return true
}

func findRandomAliveCard(cards *[5]Card) *Card {
	var aliveCardIndexes = make([]int, 0)

	for i := 0; i < len(cards); i++ {
		if cards[i].status == "alive" {
			aliveCardIndexes = append(aliveCardIndexes, i)
		}
	}

	var randomCardIndex int = aliveCardIndexes[rand.Intn(len(aliveCardIndexes))]
	var randomCard *Card = &cards[randomCardIndex]

	return randomCard
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

func printBoard(player1Cards [5]Card, player2Cards [5]Card) {
	println("Player 1")
	for i := 0; i < len(player1Cards); i++ {
		println("Card ", i+1, ": (", player1Cards[i].attack, ", ", player1Cards[i].health, ")", player1Cards[i].status)
	}

	println("")

	println("Player 2")
	for i := 0; i < len(player2Cards); i++ {
		println("Card ", i+1, ": (", player2Cards[i].attack, ", ", player2Cards[i].health, ")", player2Cards[i].status)
	}
}
