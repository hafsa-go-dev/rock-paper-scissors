package game

import "math/rand"

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var WinList = []string{"Great", "It's your lucky day!", "Nice!", "You won!", "Gz beast, what's next?"}

var DrawList = []string{"Uh...", "Try again!", "Great minds think alike", "How drawful!", "??????"}

var LoseList = []string{"Better luck next time!", "It's not your day, huh?", "Oof...", "You lose!", "Well, then..."}

type Round struct {
	Result         string `json:"result"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {

	computerChoice := ""
	result := ""
	roundResult := ""

	computerValue := rand.Intn(3)

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw!"
		result = DrawList[rand.Intn(6)]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		result = WinList[rand.Intn(6)]
	} else {
		roundResult = "Computer wins!"
		result = LoseList[rand.Intn(6)]
	}

	var round = Round{
		Result:         result,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}

	return round
}
