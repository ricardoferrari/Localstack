package betRepository

import (
	betModels "github.com/ricardoferrari/localstack/models"
)

type MessageRepositoryInterface interface {
	GetMessages() []betModels.Message
	AddMessage(message betModels.Message)
}

type MessageRepository struct {
	messages []betModels.Message
}

func (messageRepository *MessageRepository) GetMessages() []betModels.Message {
	return messageRepository.messages
}

func (messageRepository *MessageRepository) AddMessage(message betModels.Message) {
	messageRepository.messages = append(messageRepository.messages, message)
}

func NewMessageRepository() MessageRepositoryInterface {
	return &MessageRepository{}
}
