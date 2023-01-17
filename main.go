package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jeheskielSunloy77/go-aceinthehole/models"
)

type NewGameResponse struct {
	Players []models.Player `json:"players"`
	Blinds  models.Blinds   `json:"blinds"`
}

func main() {
	var game models.Game

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		game.New([]models.Player{
			{Name: "Jeheskiel", Cash: 1000},
			{Name: "Sunloy", Cash: 1000},
		}, models.Blinds{
			Small: models.Blind{Amount: 10, PlayerName: "Jeheskiel"},
			Big:   models.Blind{Amount: 20, PlayerName: "Sunloy"},
		})
		game.Deal(2)
		json, err := json.Marshal(map[string]interface{}{
			"message": "New game created",
			"game":    game,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)

	})

	http.HandleFunc("/new-game", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var response NewGameResponse
		if err := decoder.Decode(&response); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		game.New(response.Players, response.Blinds)
		json, err := json.Marshal(map[string]interface{}{
			"message": "New game created",
			"players": game.Players,
			"pot":     game.Pot,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)
	})
	http.HandleFunc("/deal", func(w http.ResponseWriter, r *http.Request) {
		handSizeQuery := r.URL.Query().Get("hand-size")

		handSize, err := strconv.Atoi(handSizeQuery)
		if err != nil {
			handSize = 2
		}
		game.Deal(handSize)
		fmt.Println(game.Players)
		hands := make([][]models.Card, len(game.Players))
		for i, player := range game.Players {
			hands[i] = player.Hand.Cards
		}
		json, err := json.Marshal(map[string]interface{}{
			"message": "Cards dealt",
			"hands":   hands,
			"flop":    game.Flop,
			"pot":     game.Pot,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)

	})

	http.HandleFunc("/turn", func(w http.ResponseWriter, r *http.Request) {
		game.Flop = append(game.Flop, game.Deck.Deal(1).Cards[0])
		json, err := json.Marshal(map[string]interface{}{
			"message": "Turn dealt",
			"flop":    game.Flop,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)
	})

	http.HandleFunc("/game-check", func(w http.ResponseWriter, r *http.Request) {
		playersRanks := game.GameCheck()

		json, err := json.Marshal(map[string]interface{}{
			"message":       "Game check",
			"players-ranks": playersRanks,
		})
		if err != nil {
			w.Write([]byte(`{"message":"Error"}`))
		}
		w.Write(json)
	})
	http.HandleFunc("/bet", func(w http.ResponseWriter, r *http.Request) {
		playerQuery := r.URL.Query().Get("player")
		betQuery := r.URL.Query().Get("bet")
		bet, err := strconv.Atoi(betQuery)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		if playerQuery == "" && betQuery == "" {
			w.Write([]byte("player and bet are required"))
			return
		}
		game.Bet(playerQuery, bet)
		w.Write([]byte(`{"message":"Bet placed", "player":` + playerQuery + `, "bet":` + betQuery + `, "pot":` + strconv.Itoa(game.Pot) + `}`))
	})
	http.HandleFunc("/fold", func(w http.ResponseWriter, r *http.Request) {
		playerQuery := r.URL.Query().Get("player")
		if playerQuery == "" {
			w.Write([]byte("player is required"))
			return
		}
		game.Fold(playerQuery)
		w.Write([]byte(`{"message":"Player folded", "player":` + playerQuery + `}`))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
