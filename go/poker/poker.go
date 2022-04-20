package poker

import (
	"fmt"
	"sort"
	"strings"
)

var ranks map[string]int = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14}
var suits map[string]int = map[string]int{"♢": 1, "♡": 2, "♤": 3, "♧": 4}

func scoreKickers(ranking []int) int {
	f := 1
	score := 0
	for _, r := range ranking {
		score += f * r
		f *= 100
	}
	return score
}

func BestHand(hands []string) ([]string, error) {
	hs := []struct {
		category, score int
		hand            string
	}{}
	for _, hand := range hands {
		cards := strings.Split(hand, " ")
		if len(cards) != 5 {
			return []string{}, fmt.Errorf("Invalid hand %s", hand)
		}
		ranksMap := map[string]int{}
		suitsMap := map[string]int{}
		ranking := []int{}
		for _, card := range cards {
			if len(card) < 4 || len(card) > 5 {
				return []string{}, fmt.Errorf("Invalid card %s", card)
			}
			rank := card[0 : len(card)-3]
			r, found := ranks[rank]
			if !found {
				return []string{}, fmt.Errorf("Invalid rank %s", rank)
			}
			ranksMap[rank]++
			ranking = append(ranking, r)
			suit := string(card[len(card)-3:])
			_, found = suits[suit]
			if !found {
				return []string{}, fmt.Errorf("Invalid suit %s", suit)
			}
			suitsMap[suit]++
		}
		sort.Slice(ranking, func(p, q int) bool {
			return ranking[p] < ranking[q]
		})
		score := 0
		category := 0
		switch len(ranksMap) {
		case 5: //highest card, straight, flush or straight flush
			category = 0
			if ranking[0] == ranking[4]-4 || (ranking[0] == ranks["2"] && ranking[4] == ranks["A"] && ranking[3] == ranks["5"]) {
				if ranking[0] == ranks["2"] && ranking[4] == ranks["A"] {
					ranking = append([]int{1}, ranking[0:4]...)
				}
				category = 4
				if len(suitsMap) == 1 {
					category = 8
				}
				score += ranking[4]
				break
			}
			if len(suitsMap) == 1 {
				category = 5
			}
			score += scoreKickers(ranking)
		case 2: //4 of a kind or full house
			category = 6
			for k, r := range ranksMap {
				switch r {
				case 4:
					category = 7
					score += ranks[k] * 100
				case 3:
					score += ranks[k] * 10000
				case 2:
					score += ranks[k] * 100
				default:
					score += ranks[k]
				}
			}
		case 3: //2 pairs or 3 of a kind
			highPair := 0
			for k, m := range ranksMap {
				switch m {
				case 2:
					if highPair < ranks[k] {
						highPair = ranks[k]
					}
					category = 2
				case 3:
					category = 3
					score = ranks[k] * 10000000000
					break
				}
			}
			if category == 2 {
				score = highPair * 10000000000
			}
			score += scoreKickers(ranking)
		case 4: //pair
			category = 1
			for k, m := range ranksMap {
				if m == 2 {
					score = ranks[k] * 10000000000
					break
				}
			}
			score += scoreKickers(ranking)
		}
		hs = append(hs, struct {
			category, score int
			hand            string
		}{category, score, hand})
	}
	if len(hands) == 1 {
		return hands, nil
	}
	out := []string{}
	highCategory := 0
	highScore := 0
	sort.Slice(hs, func(p, q int) bool {
		return hs[p].category > hs[q].category ||
			hs[p].category == hs[q].category && hs[p].score > hs[q].score
	})
	for _, h := range hs {
		if highCategory > h.category || highCategory == h.category && highScore > h.score {
			break
		}
		highCategory = h.category
		highScore = h.score
		out = append(out, h.hand)
	}
	return out, nil
}

        

          
