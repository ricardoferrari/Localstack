package betModels

type Game struct {
	Id           int
	Championship string
	Stage        string
	Team1        string
	Team2        string
	Score1       int
	Score2       int
}

type Bet struct {
	Id         int
	Player     string
	Games      []Game
	TotalScore int
}
