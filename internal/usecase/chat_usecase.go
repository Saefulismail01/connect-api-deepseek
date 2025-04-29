package usecase

import "initialProject/internal/domain"

type ChatUsecase struct {
	repo domain.ChatRepository
}

func NewChatUsecase(repo domain.ChatRepository) *ChatUsecase {
	return &ChatUsecase{repo: repo}
}

func (u *ChatUsecase) Chat(request domain.ChatRequest) (domain.ChatResponse, error) {
	return u.repo.SendChat(request)
}
