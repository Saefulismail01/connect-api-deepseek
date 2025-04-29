package domain

type ChatRepository interface {
	SendChat(request ChatRequest) (ChatResponse, error)
}
