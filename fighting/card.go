package main

import "math/rand"

type CardRace uint8

const (
	Common CardRace = iota
	Murloc          = 1
	Elf             = 2
	Orc             = 3
)

type Card struct {
	attack int
	health int
	status string
	race   int
}

func fillCards(cards *[5]Card) {
	for i := 0; i < len(cards); i++ {
		card := Card{rand.Intn(10) + 1, rand.Intn(10) + 1, "alive", rand.Intn(4)}
		cards[i] = card
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
	aliveCardIndexes := make([]int, 0)

	for i := 0; i < len(cards); i++ {
		if cards[i].status == "alive" {
			aliveCardIndexes = append(aliveCardIndexes, i)
		}
	}

	randomCardIndex := aliveCardIndexes[rand.Intn(len(aliveCardIndexes))]
	randomCard := &cards[randomCardIndex]

	return randomCard
}

func getRandomCardRace() CardRace {
	// TODO: Remove hardcode 4
	return CardRace(rand.Intn(4))
}
