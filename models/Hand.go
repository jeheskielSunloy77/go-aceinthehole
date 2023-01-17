package models

import "sort"

type Hand struct {
	Cards []Card
}

func (h *Hand) Rank() int {
	rankCounts := make(map[int]int)
	for _, card := range h.Cards {
		rankCounts[card.Rank]++
	}
	suitCounts := make(map[string]int)
	for _, card := range h.Cards {
		suitCounts[card.Suit]++
	}

	isRoyalFlush := false
	for _, count := range suitCounts {
		if count == 5 {
			isRoyalFlush = true
			for _, card := range h.Cards {
				if card.Rank < 10 {
					isRoyalFlush = false
					break
				}
			}
		}
	}
	if isRoyalFlush {
		return 1
	}

	isStraightFlush := false
	for _, count := range suitCounts {
		if count >= 5 {
			isStraightFlush = true
			for i := 0; i <= len(h.Cards)-5; i++ {
				if h.Cards[i].Suit != h.Cards[i+1].Suit || h.Cards[i].Rank+1 != h.Cards[i+1].Rank {
					isStraightFlush = false
					break
				}
			}
		}
	}
	if isStraightFlush {
		return 2
	}

	for _, count := range rankCounts {
		if count == 4 {
			return 3
		}
	}

	isThreeOfAKind := false
	isPair := false
	for _, count := range rankCounts {
		if count == 3 {
			isThreeOfAKind = true
		} else if count == 2 {
			isPair = true
		}
	}
	if isThreeOfAKind && isPair {
		return 4
	}

	for _, count := range suitCounts {
		if count >= 5 {
			return 5
		}
	}

	isStraight := false
	for i := 0; i <= len(h.Cards)-5; i++ {
		if h.Cards[i].Rank+1 == h.Cards[i+1].Rank {
			isStraight = true
		} else {
			isStraight = false
			break
		}
	}
	if isStraight {
		return 6
	}

	for _, count := range rankCounts {
		if count == 3 {
			return 7
		}
	}

	isTwoPair := false
	for _, count := range rankCounts {
		if count == 2 {
			isTwoPair = true
		}
	}
	if isTwoPair {
		return 8
	}

	for _, count := range rankCounts {
		if count == 2 {
			return 9
		}
	}

	return 10
}

func (h *Hand) TieBreaker() []int {
	tieBreaker := make([]int, 0)

	for _, card := range h.Cards {
		tieBreaker = append(tieBreaker, card.Rank)
	}

	sort.Slice(tieBreaker, func(i, j int) bool {
		return tieBreaker[i] > tieBreaker[j]
	})

	return tieBreaker
}

func (h *Hand) Compare(otherHand Hand) int {
	hRank := h.Rank()
	otherHRank := otherHand.Rank()
	if hRank > otherHRank {
		return 1
	} else if hRank < otherHRank {
		return -1
	} else {
		hTieBreaker := h.TieBreaker()
		otherHTieBreaker := otherHand.TieBreaker()
		for i := 0; i < len(hTieBreaker); i++ {
			if hTieBreaker[i] > otherHTieBreaker[i] {
				return 1
			} else if hTieBreaker[i] < otherHTieBreaker[i] {
				return -1
			}
		}
		return 0
	}
}
