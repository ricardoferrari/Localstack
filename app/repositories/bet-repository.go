package betRepository

import (
	betModels "github.com/ricardoferrari/localstack/models"
)

type BetRepositoryInterface interface {
	GetBets() []betModels.Bet
}

type BetRepository struct {
	bets []betModels.Bet
}

func (betRepository *BetRepository) GetBets() []betModels.Bet {
	return betRepository.bets
}

func NewBetRepository() BetRepositoryInterface {
	return &BetRepository{}
}
