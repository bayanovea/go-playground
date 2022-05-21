package main

func main() {
	var game1 = addNewGame()
	var game2 = addNewGame()

	var player1 = Player{hash: "player1", name: "Player 1"}
	var player2 = Player{hash: "player2", name: "Player 2"}

	addPlayerToGame(&game1, player1)
	addPlayerToGame(&game1, player2)

	addPlayerToGame(&game1, player1)
	addPlayerToGame(&game2, player2)

	var game1player1Cards [5]Card
	var game1player2Cards [5]Card
	var game2player1Cards [5]Card
	var game2player2Cards [5]Card

	fillCards(&game1player1Cards)
	fillCards(&game1player2Cards)
	fillCards(&game2player1Cards)
	fillCards(&game2player2Cards)

	printBoard(game1player1Cards, game1player2Cards)
	go playGame(game1, &game1player1Cards, &game1player2Cards)

	go printBoard(game2player1Cards, game2player2Cards)
	playGame(game2, &game2player1Cards, &game2player2Cards)
}
