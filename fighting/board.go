package main

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
