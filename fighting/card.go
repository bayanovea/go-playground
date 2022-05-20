package main

import "math/rand"

type Card struct {
	attack int
	health int
	status string
}

func fillCards(cards *[5]Card) {
	for i := 0; i < len(cards); i++ {
		cards[i].attack = rand.Intn(10) + 1
		cards[i].health = rand.Intn(10) + 1
		cards[i].status = "alive"
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
