package models

type Blinds struct {
	Small Blind
	Big   Blind
}

type Blind struct {
	Amount     int
	PlayerName string
}
