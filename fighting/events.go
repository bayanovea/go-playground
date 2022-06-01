package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameEventType uint8

const (
	_                              GameEventType = iota
	GameEventType_BuffSpecificRace               = 1
	GameEventType_Disease                        = 2
)

type GameEvent struct {
	game  Game
	_type GameEventType
}

func startGameEventsLoop() {
	var gameEventsChannel chan GameEvent = make(chan GameEvent)
	_handleGameEvents(gameEventsChannel)

	for {
		time.Sleep(time.Second * 1)

		randomGameEventType := GameEventType(rand.Intn(2) + 1)
		activeGames := getActiveGames()

		for _, game := range activeGames {
			newGameEvent := GameEvent{game: game, _type: randomGameEventType}
			gameEventsChannel <- newGameEvent

			fmt.Println("New event", newGameEvent._type, "for game", game.name, "dispatched")
		}
	}
}

func _handleGameEvents(gameEventsChannel chan GameEvent) {
	go func() {
		for {
			gameEvent := <-gameEventsChannel

			switch gameEvent._type {
			case GameEventType_BuffSpecificRace:
				handleBuffSpecificRaceEvent()
			case GameEventType_Disease:
				handleDiseaseEvent(&gameEvent.game)
			}
		}
	}()
}

func handleBuffSpecificRaceEvent() {
	fmt.Println("handleBuffSpecificRaceEvent")
}

func handleDiseaseEvent(game *Game) {
	fmt.Println("handleDiseaseEvent")
}
