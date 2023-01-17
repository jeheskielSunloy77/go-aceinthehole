package models

type Player struct {
	Name string
	Hand Hand
	Cash int
}

func (p *Player) Bet(amount int, g *Game) {
	p.Cash -= amount
	g.Pot += amount
}

func (p *Player) Win(amount int, g *Game) {
	p.Cash += amount
	g.Pot -= amount
}

func (p *Player) Fold() {
	p.Hand = Hand{}
}
