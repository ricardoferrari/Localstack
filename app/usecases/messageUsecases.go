package betUseCase

import (
	betModels "github.com/ricardoferrari/localstack/models"
	messageRepository "github.com/ricardoferrari/localstack/repositories"
)

type MessageUseCaseInterface interface {
	GetMessages() []betModels.Message
	AddMessage(message betModels.Message)
}

type MessageUseCase struct {
	messageRepository messageRepository.MessageRepositoryInterface
}

func (messageUseCase *MessageUseCase) GetMessages() []betModels.Message {
	return messageUseCase.messageRepository.GetMessages()
}

func (messageUseCase *MessageUseCase) AddMessage(message betModels.Message) {
	messageUseCase.messageRepository.AddMessage(message)
}

func NewMessageUseCase(repository messageRepository.MessageRepositoryInterface) MessageUseCaseInterface {
	return &MessageUseCase{messageRepository: repository}
}
