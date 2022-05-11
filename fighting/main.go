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

	for i := 0; i < 5; i++ {
		println("=====")
		println("Turn ", i+1)
		println("=====")

		// var attackCard *Card
		// var defenseCard *Card

		// pointer to element from array

		var attackCard *Card
		var defenseCard *Card

		if whoseTurn == "player1" {
			attackCard = &player1Cards[rand.Intn(len(player1Cards))]
			defenseCard = &player2Cards[rand.Intn(len(player2Cards))]
		} else if whoseTurn == "player2" {
			attackCard = &player2Cards[rand.Intn(len(player2Cards))]
			defenseCard = &player1Cards[rand.Intn(len(player1Cards))]
		}

		attack(attackCard, defenseCard)

		if whoseTurn == "player1" {
			whoseTurn = "player2"
		} else if whoseTurn == "player2" {
			whoseTurn = "player1"
		}

		printBoard(*player1Cards, *player2Cards)
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
