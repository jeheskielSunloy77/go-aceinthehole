package main

import (
	"fmt"
	"math/rand"
)

var suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var ranks = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

const (
	Hearts = iota
	Diamonds
	Clubs
	Spades
)

const (
	Ace = iota
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
	Rank int
	Suit int
}

type Hand struct {
	Card1 Card
	Card2 Card
}

type Player struct {
	Name string
	Hand Hand
	Cash int
	Seat int
}

func (p *Player) ReadHand() {
	fmt.Printf("Hand = %s of %s and %s of %s\n", ranks[p.Hand.Card1.Rank], suits[p.Hand.Card1.Suit], ranks[p.Hand.Card2.Rank], suits[p.Hand.Card2.Suit])
}
func (p *Player) ReadStats() {
	fmt.Printf("\n=====================================%s's turn=====================================\n", p.Name)
	fmt.Printf("Cash = %d\n", p.Cash)
	fmt.Printf("Hand = %s of %s and %s of %s\n", ranks[p.Hand.Card1.Rank], suits[p.Hand.Card1.Suit], ranks[p.Hand.Card2.Rank], suits[p.Hand.Card2.Suit])
}

type Round int

const (
	preFlop Round = iota
	flop
	turn
	river
)

var rounds = []string{"Pre Flop", "Flop", "Turn", "River"}

type Bet struct {
	Amount  int
	Player  *Player
	IsAllIn bool
	Round   Round
	IsBlind bool
}

type Blind struct {
	Amount int
	Player *Player
}

type Game struct {
	Players         []Player
	InActivePlayers []Player
	Bets            []Bet
	Deck            []Card
	Board           []Card
	SmallBlind      Blind
	BigBlind        Blind
	DealerPos       int
	BuyIn           int
}

func (g *Game) AddPlayer(username string) {
	g.Players = append(g.Players, Player{
		Name: username,
		Cash: g.BuyIn,
		Seat: len(g.Players),
	})
}

func (g *Game) RemovePlayer(playerName string) {
	for i := range g.Players {
		if g.Players[i].Name == playerName {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
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

func (g *Game) AddBet(bet Bet) {
	g.Bets = append(g.Bets, bet)
}

func (g *Game) Deal() {
	numPlayers := len(g.Players)
	var smallBlindPosition, bigBlindPosition int
	var smallBlindPlayer, bigBlindPlayer *Player

	for i := 0; i < 2; i++ {
		for j := range g.Players {
			g.Players[j].Hand = Hand{
				Card1: g.Deck[0],
				Card2: g.Deck[1],
			}
			g.Deck = g.Deck[2:]
		}
	}

	if g.DealerPos == 0 {
		smallBlindPosition = 0
		bigBlindPosition = 1

		g.SmallBlind.Amount = 1
		g.BigBlind.Amount = 2
	} else {
		smallBlindPosition = (g.DealerPos + 1) % numPlayers
		bigBlindPosition = (g.DealerPos + 2) % numPlayers

		g.SmallBlind.Amount = g.SmallBlind.Amount * 2
		g.BigBlind.Amount = g.BigBlind.Amount * 2
	}

	for i := 0; i < numPlayers; i++ {
		if g.Players[i].Seat == smallBlindPosition {
			smallBlindPlayer = &g.Players[i]
		} else if g.Players[i].Seat == bigBlindPosition {
			bigBlindPlayer = &g.Players[i]
		}
	}

	g.SmallBlind.Player = smallBlindPlayer
	g.BigBlind.Player = bigBlindPlayer

	g.AddBet(Bet{Amount: g.SmallBlind.Amount, Player: g.SmallBlind.Player})
	g.SmallBlind.Player.Cash -= g.SmallBlind.Amount
	g.AddBet(Bet{Amount: g.BigBlind.Amount, Player: g.BigBlind.Player})
	g.BigBlind.Player.Cash -= g.BigBlind.Amount

	g.DealerPos = (g.DealerPos + 1) % numPlayers

	fmt.Printf("Small blind ammount is %d on %s\n", g.SmallBlind.Amount, g.SmallBlind.Player.Name)
	fmt.Printf("Big blind ammount is %d on %s\n", g.BigBlind.Amount, g.BigBlind.Player.Name)

}

func (g *Game) FoldPlayer(playerIndex int) {
	g.InActivePlayers = append(g.InActivePlayers, g.Players[playerIndex])
	g.Players = append(g.Players[:playerIndex], g.Players[playerIndex+1:]...)
}

func (g *Game) BetPlayer(playerIndex int, bet Bet) {
	g.AddBet(Bet{Amount: bet.Amount, Player: &g.Players[playerIndex], IsAllIn: bet.IsAllIn, Round: bet.Round})
	g.Players[playerIndex].Cash -= bet.Amount
}

func (g *Game) ReadBoard(round Round) {
	fmt.Printf("\n===============================%s's Board===============================\n", rounds[round])
	for i := range g.Board {
		fmt.Printf("%s of %s\n", ranks[g.Board[i].Rank], suits[g.Board[i].Suit])
	}
	fmt.Println("========================================================================")
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

func main() {
	game := Game{BuyIn: 30}
	var doneAddingPlayers bool
	var input string

	for !doneAddingPlayers {
		fmt.Print("Enter player name name: ")
		fmt.Scanln(&input)
		game.AddPlayer(input)
		fmt.Print("Add another player? (y/n)")
		fmt.Scanln(&input)
		if input != "y" {
			if len(game.Players) < 2 {
				fmt.Println("You need at least 2 players to play")
				continue
			}
			doneAddingPlayers = true
		}

	}

	game.CreateDeck()

	fmt.Println("Shuffling the deck...")
	game.ShuffleDeck()

	fmt.Println("Dealing the cards...")
	game.Deal()
	game.Flop()
	game.ReadBoard(flop)

	for loopCount := 0; loopCount < 1; loopCount++ {
		var betToCall Bet
		for i := range game.Players {
			if betToCall.Player == &game.Players[i] {
				continue
			}
			game.Players[i].ReadStats()

			for j := range game.Bets {
				if game.Bets[j].Round == flop {
					if game.Bets[j].Player != &game.Players[i] {
						if game.Players[i].Cash >= game.Bets[j].Amount {
							betToCall = game.Bets[j]
						} else {
							fmt.Println("You don't have enough cash to call")
							game.FoldPlayer(i)
						}
					}
				}
			}

			if betToCall.Amount > 0 {
				fmt.Printf("\nThe bet to call is %d by %s\n", betToCall.Amount, betToCall.Player.Name)
				fmt.Print("Fold (f), bet (b) or call (c)? ")
			} else {
				fmt.Print("Fold (f), bet (b) or check (c)? ")
			}
			fmt.Scanln(&input)
			if input == "f" {
				game.FoldPlayer(i)
			} else if input == "b" {
				var ammount int
				fmt.Print("How much would you like to bet? ")
				fmt.Scanln(&ammount)
				game.BetPlayer(i, Bet{Amount: ammount, Round: flop})
				loopCount -= 1
			} else if input == "c" {
				if betToCall.Amount > 0 {
					game.BetPlayer(i, Bet{Amount: betToCall.Amount, Round: flop})
				} else {
					fmt.Println("Checking")
				}
			} else {
				panic("Invalid input")
			}
		}
	}
	game.Turn()
	game.ReadBoard(turn)
	for loopCount := 0; loopCount < 1; loopCount++ {
		var betToCall Bet
		for i := range game.Players {
			if betToCall.Player == &game.Players[i] {
				continue
			}
			game.Players[i].ReadStats()

			for j := range game.Bets {
				if game.Bets[j].Round == turn {
					if game.Bets[j].Player != &game.Players[i] {
						if game.Players[i].Cash >= game.Bets[j].Amount {
							betToCall = game.Bets[j]
						} else {
							fmt.Println("You don't have enough cash to call")
							game.FoldPlayer(i)
						}
					}
				}
			}

			if betToCall.Amount > 0 {
				fmt.Printf("\nThe bet to call is %d by %s\n", betToCall.Amount, betToCall.Player.Name)
				fmt.Print("Fold (f), bet (b) or call (c)? ")
			} else {
				fmt.Print("Fold (f), bet (b) or check (c)? ")
			}
			fmt.Scanln(&input)
			if input == "f" {
				game.FoldPlayer(i)
			} else if input == "b" {
				var ammount int
				fmt.Print("How much would you like to bet? ")
				fmt.Scanln(&ammount)
				game.BetPlayer(i, Bet{Amount: ammount, Round: turn})
				loopCount -= 1
			} else if input == "c" {
				if betToCall.Amount > 0 {
					game.BetPlayer(i, Bet{Amount: betToCall.Amount, Round: turn})
				} else {
					fmt.Println("Checking")
				}
			} else {
				panic("Invalid input")
			}
		}
	}

	game.Turn()
	game.ReadBoard(river)
	for loopCount := 0; loopCount < 1; loopCount++ {
		var betToCall Bet
		for i := range game.Players {
			if betToCall.Player == &game.Players[i] {
				continue
			}
			game.Players[i].ReadStats()

			for j := range game.Bets {
				if game.Bets[j].Round == river {
					if game.Bets[j].Player != &game.Players[i] {
						if game.Players[i].Cash >= game.Bets[j].Amount {
							betToCall = game.Bets[j]
						} else {
							fmt.Println("You don't have enough cash to call")
							game.FoldPlayer(i)
						}
					}
				}
			}

			if betToCall.Amount > 0 {
				fmt.Printf("\nThe bet to call is %d by %s\n", betToCall.Amount, betToCall.Player.Name)
				fmt.Print("Fold (f), bet (b) or call (c)? ")
			} else {
				fmt.Print("Fold (f), bet (b) or check (c)? ")
			}
			fmt.Scanln(&input)
			if input == "f" {
				game.FoldPlayer(i)
			} else if input == "b" {
				var ammount int
				fmt.Print("How much would you like to bet? ")
				fmt.Scanln(&ammount)
				game.BetPlayer(i, Bet{Amount: ammount, Round: river})
				loopCount -= 1
			} else if input == "c" {
				if betToCall.Amount > 0 {
					game.BetPlayer(i, Bet{Amount: betToCall.Amount, Round: river})
				} else {
					fmt.Println("Checking")
				}
			} else {
				panic("Invalid input")
			}
		}
	}

}
