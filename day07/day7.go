package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func (h *Hand) score() int {
	cards := [14]int{}
	for _, card := range h.cards {
		cards[cardVals[string(card)]] += 1
	}
	numCards := [6]int{}
	for _, count := range cards {
		numCards[count] += 1
	}
	if numCards[5] == 1 {
		return int(FiveOfAKind)
	} else if numCards[4] == 1 {
		return int(FourOfAKind)
	} else if numCards[3] == 1 && numCards[2] == 1 {
		return int(FullHouse)
	} else if numCards[3] == 1 {
		return int(ThreeOfAKind)
	} else if numCards[2] == 2 {
		return int(TwoPair)
	} else if numCards[2] == 1 {
		return int(OnePair)
	} else {
		return int(HighCard)
	}

}

func (h *Hand) scoreWithJokers() int {
	cards := [14]int{}
	for _, card := range h.cards {
		cards[jokerVals[string(card)]] += 1
	}
	numCards := [6]int{}

	for rank, count := range cards {
		if rank > 0 {
			numCards[count] += 1
		}
	}
	if cards[0] > 0 {
		for i := len(numCards) - 1; i >= 0; i -= 1 {
			if numCards[i] > 0 {
				numCards[i] -= 1
				numCards[i+cards[0]] = 1
				break
			}
		}
	}
	if numCards[5] == 1 {
		return int(FiveOfAKind)
	} else if numCards[4] == 1 {
		return int(FourOfAKind)
	} else if numCards[3] == 1 && numCards[2] == 1 {
		return int(FullHouse)
	} else if numCards[3] == 1 {
		return int(ThreeOfAKind)
	} else if numCards[2] == 2 {
		return int(TwoPair)
	} else if numCards[2] == 1 {
		return int(OnePair)
	} else {
		return int(HighCard)
	}

}

var cardVals = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

var jokerVals = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func FindWinnings(input []string) int {
	games := []Hand{}
	for _, game := range input {
		split := strings.Fields(game)
		bidVal, _ := strconv.Atoi(split[1])
		games = append(games, Hand{split[0], bidVal})
	}
	sort.Slice(games, func(i, j int) bool {
		iScore := games[i].score()
		jScore := games[j].score()
		if iScore == jScore {
			for place := range games[i].cards {
				if games[i].cards[place] != games[j].cards[place] {
					return cardVals[string(games[i].cards[place])] < cardVals[string(games[j].cards[place])]
				}
			}
			return false
		} else {
			return iScore < jScore
		}
	})
	total := 0
	for rank, game := range games {
		total += (rank + 1) * game.bid
	}
	return total
}

func FindWinningsWithJokers(input []string) int {
	games := []Hand{}
	for _, game := range input {
		split := strings.Fields(game)
		bidVal, _ := strconv.Atoi(split[1])
		games = append(games, Hand{split[0], bidVal})
	}
	sort.SliceStable(games, func(i, j int) bool {
		iScore := games[i].scoreWithJokers()
		jScore := games[j].scoreWithJokers()
		if iScore == jScore {
			for place := range games[i].cards {
				if games[i].cards[place] != games[j].cards[place] {
					return jokerVals[string(games[i].cards[place])] < jokerVals[string(games[j].cards[place])]
				}
			}
			return false
		} else {
			return iScore < jScore
		}
	})
	total := 0
	for rank, game := range games {
		total += (rank + 1) * game.bid
	}
	return total
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("Winnings:", FindWinnings(input))
	fmt.Println("Joker Winnings:", FindWinningsWithJokers(input))
}
