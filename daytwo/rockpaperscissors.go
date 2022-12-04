package daytwo

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type play int64
type gameDecision int64

const (
	theirRock play = iota
	theirPaper
	theirScissors
	myRock
	myPaper
	myScissors
)
const (
	lose gameDecision = iota
	draw
	win
)

type game struct {
	theirPlay, myPlay play
	decision gameDecision
}

func Play() {
	games, err := readInputFromFile()
	if err != nil {
		log.Fatalln(err)
	}

	games = setMyMoves(games)

	fmt.Println("Total score: ", getTotalScore(games))
	// fmt.Printf("games: %v\n error: %#v", games, err)
}

func setMyMoves(games []game) []game {
	newGames := []game{}
	for _, v := range games {
		newGames = append(newGames, game{
			theirPlay: v.theirPlay,
			myPlay: calculateMyPlay(v),
			decision: v.decision,
		})
	}
	return newGames
}

func getTotalScore(games []game) int {
	sum := 0
	for _, v := range games {
		sum += v.score()
	}
	return sum
}

func readInputFromFile() ([]game, error) {
	b, err := os.ReadFile("daytwo/input.txt")
	if err != nil {
		return nil, err
	}
	gameStrings := strings.Split(string(b), "\n")
	
	var games []game
	for _, v := range gameStrings {
		sArr := strings.Split(v, " ")
		games = append(games, game{
			theirPlay: convertStringToPlay(sArr[0]),
			decision: convertStringToGame(sArr[1]),
		})
	}

	return games, nil
}

func convertStringToGame(s string) gameDecision {
	switch s {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	default:
		panic("invalid game decision")
	}
}

func convertStringToPlay(s string) play {
	switch s {
	case "A":
		return theirRock
	case "B":
		return theirPaper
	case "C":
		return theirScissors
	default:
		panic("invalid play string")
	}
}

func (p play) String() string {
	switch p {
	case theirRock, myRock:
		return "Rock"
	case theirPaper, myPaper:
		return "Paper"
	case theirScissors, myScissors:
		return "Scissors"
	default:
		panic("invalid play")
	}
}

func (d gameDecision) String() string {
	switch d {
	case lose:
		return "LOSE"
	case draw:
		return "DRAW"
	case win:
		return "WIN"
	default:
		panic("invalid game decision")
	}
}

func (g game) String() string {
	return fmt.Sprintf("Game:\n their play %s\n game decision %s\n Score: %d\n", g.theirPlay.String(), g.decision.String(), g.score())
}

func Print(games []game) {
	for _, v := range games {
		fmt.Println(v.String())
	}
}

func (p play) score() int {
	switch p {
	case myRock:
		return 1
	case myPaper:
		return 2
	case myScissors:
		return 3
	default:
		return 2000
	}
}

func (g game) score() int {
	playScore := g.myPlay.score()

	switch g.decision {
	case win:
		return 6 + playScore
	case draw:
		return 3 + playScore
	case lose:
		return 0 + playScore
	default:
		return 1000
	}
}

func calculateMyPlay(g game) play {
	if g.decision == lose {
		switch g.theirPlay {
		case theirScissors:
			return myPaper
		case theirRock:
			return myScissors
		default:
			return myRock
		}
	}
	if g.decision == draw {
		switch g.theirPlay {
		case theirRock:
			return myRock
		case theirPaper:
			return myPaper
		default:
			return myScissors
		}
	}
	if g.decision == win {
		switch g.theirPlay {
		case theirRock:
			return myPaper
		case theirPaper:
			return myScissors
		default:
			return myRock
		}
	}
	panic("can't compute my play")
}