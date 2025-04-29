package main

import (
	"initialProject/internal/config"
	"initialProject/internal/delivery/http"
	"initialProject/internal/delivery/cli"
	"initialProject/internal/usecase"
	"initialProject/internal/infrastructure"
)

func main() {
	config.LoadEnv()
	apiKey := config.GetDeepseekAPIKey()

	repo := infrastructure.NewDeepseekAPI(apiKey)
	chatUsecase := usecase.NewChatUsecase(repo)

	go http.StartServer(chatUsecase)
	cli.StartCLI(chatUsecase)
}
