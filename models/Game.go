package models

type Game struct {
	Players []Player
	Deck    Deck
	Pot     int
}

func (g *Game) New() {
	g.Deck.New()
	g.Deck.Shuffle()
}

func (g *Game) Deal(handSize int) {
	for _, player := range g.Players {
		player.Hand = *g.Deck.Deal(handSize)
	}
}
