package models

import (
	"math/rand"
)

type Deck struct {
	Cards []Card
}

var suits = [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
var ranks = [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Deal(handSize int) *Hand {
	cards := d.Cards[:handSize]
	d.Cards = d.Cards[handSize:]
	return &Hand{Cards: cards}
}

func (d *Deck) New() {
	cards := make([]Card, 52)
	for i, suit := range suits {
		for j, rank := range ranks {
			cards[i*13+j] = Card{Suit: suit, Rank: rank}
		}
	}
	d.Cards = cards
}
