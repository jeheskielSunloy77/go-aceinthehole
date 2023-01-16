package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jeheskielSunloy77/go-aceinthehole/models"
)

func main() {
	var game models.Game

	http.HandleFunc("/new-game", func(w http.ResponseWriter, r *http.Request) {
		playersQuery := r.URL.Query().Get("players")
		players, err := strconv.Atoi(playersQuery)
		if err != nil {
			w.Write([]byte("Invalid number of players"))
			return
		}
		game.Players = make([]models.Player, players)
		game.New()
		w.Write([]byte(`{"message":"New game created", "players":` + playersQuery + `}`))
	})
	http.HandleFunc("/deal", func(w http.ResponseWriter, r *http.Request) {
		handSizeQuery := r.URL.Query().Get("hand-size")

		handSize, err := strconv.Atoi(handSizeQuery)
		if err != nil {
			w.Write([]byte("Invalid hand size"))
			return
		}
		game.Deal(handSize)
		hands := make([][]models.Card, len(game.Players))
		for i, player := range game.Players {
			hands[i] = player.Hand.Cards
		}
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(hands)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonResp)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
