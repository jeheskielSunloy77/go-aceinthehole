package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

type Rank int

const (
	Ace Rank = iota + 1 // Add 1 to start ranks from 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Rank Rank
	Suit Suit
}

type Hand struct {
	Card1 Card
	Card2 Card
}

type Player struct {
	Name string
	Hand Hand
	Cash int
}

type Game struct {
	Players []Player
	Pot     int
	Deck    []Card
	Board   []Card
}

func (g *Game) CreateDeck() {
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; rank++ {
			g.Deck = append(g.Deck, Card{rank, suit})
		}
	}
}

func (g *Game) ShuffleDeck() {
	for i := range g.Deck {
		j := i + rand.Intn(len(g.Deck)-i)
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	}
}

func (g *Game) AddPlayer(player Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) Deal() {
	for i := 0; i < 2; i++ {
		for j := range g.Players {
			g.Players[j].Hand = Hand{
				Card1: g.Deck[0],
				Card2: g.Deck[1],
			}
			g.Deck = g.Deck[2:]
		}
	}
}

func (g *Game) Flop() {
	for i := 0; i < 3; i++ {
		g.Board = append(g.Board, g.Deck[0])
		g.Deck = g.Deck[1:]
	}
}

func (g *Game) Turn() {
	g.Board = append(g.Board, g.Deck[0])
	g.Deck = g.Deck[1:]
}

func (g *Game) River() {
	g.Board = append(g.Board, g.Deck[0])
	g.Deck = g.Deck[1:]
}

func (g *Game) Showdown() {
	sort.Sort(PlayersSortByRank(g.Players))

	for i, player := range g.Players {
		fmt.Printf("Player %s have a %v %v and %v %v \n Player Ranking is %d\n", player.Name, player.Hand.Card1.Rank, player.Hand.Card1.Suit, player.Hand.Card2.Rank, player.Hand.Card2.Suit, i+1)
	}

}

type PlayersSortByRank []Player

func (a PlayersSortByRank) Len() int      { return len(a) }
func (a PlayersSortByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PlayersSortByRank) Less(i, j int) bool {
	return a[i].Hand.Card1.Rank < a[j].Hand.Card1.Rank || a[i].Hand.Card1.Rank < a[j].Hand.Card2.Rank || a[i].Hand.Card2.Rank < a[j].Hand.Card2.Rank || a[i].Hand.Card2.Rank < a[j].Hand.Card1.Rank

}

func main() {
	var game Game

	game.CreateDeck()
	game.ShuffleDeck()
	game.AddPlayer(Player{Name: "John"})
	game.AddPlayer(Player{Name: "Danny"})
	game.AddPlayer(Player{Name: "Mike"})

	game.Deal()
	game.Flop()
	game.Turn()
	game.River()
	game.Showdown()

}
