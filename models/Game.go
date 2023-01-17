package models

type Game struct {
	Players []Player
	Blinds  Blinds
	Deck    Deck
	Pot     int
	Flop    []Card
}

func (g *Game) New(payers []Player, blinds Blinds) {
	g.Blinds = blinds
	g.Players = payers
	g.Deck.New()
	g.Deck.Shuffle()
}

func (g *Game) Deal(handSize int) {
	for _, player := range g.Players {
		player.Hand = g.Deck.Deal(handSize)
	}
	g.Pot = g.Blinds.Big.Amount + g.Blinds.Small.Amount
	g.Flop = g.Deck.Deal(3).Cards

	for _, player := range g.Players {
		if player.Name == g.Blinds.Big.PlayerName {
			player.Bet(g.Blinds.Big.Amount, g)
		}
		if player.Name == g.Blinds.Small.PlayerName {
			player.Bet(g.Blinds.Small.Amount, g)
		}
	}
}

// func (g *Game) Deal(handSize int) {
// 	for i := 0; i < handSize; i++ {
// 		for _, player := range g.Players {
// 			if len(g.Deck.Cards) == 0 {
// 				return
// 			}
// 			nextCard := g.Deck.Cards[i]
// 			player.Hand.Cards = append(player.Hand.Cards, nextCard)
// 			g.Deck.Cards = g.Deck.Cards[1:]
// 		}
// 	}
// }

func (g *Game) Bet(playerName string, bet int) {
	for _, player := range g.Players {
		if player.Name == playerName {
			player.Bet(bet, g)
		}
	}
}

func (g *Game) Fold(playerName string) {
	for _, player := range g.Players {
		if player.Name == playerName {
			player.Fold()
		}
	}
}

func (g *Game) Turn() {
	if len(g.Flop) < 5 {
		g.Flop = append(g.Flop, g.Deck.Deal(1).Cards[0])
	}
}

func (g *Game) GameCheck() map[string]int {
	var bestHand Hand
	playersHand := make(map[string]int)
	for _, player := range g.Players {
		tempHand := append(player.Hand.Cards, g.Flop...)

		newHand := Hand{Cards: tempHand}

		if newHand.Compare(bestHand) > 0 {
			bestHand = newHand
		}
		playersHand[player.Name] = newHand.Rank()
	}
	return playersHand
}
