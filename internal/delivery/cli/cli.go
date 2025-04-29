package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"initialProject/internal/domain"
	"initialProject/internal/usecase"
)

func StartCLI(chatUsecase *usecase.ChatUsecase) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Chat dengan Deepseek AI. Ketik pesan dan tekan Enter. (Ctrl+C untuk keluar)")
	for {
		fmt.Print("Anda: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Gagal membaca input: %v", err)
			continue
		}
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			continue
		}
		chatReq := domain.ChatRequest{
			Model:    "deepseek-chat",
			Messages: []domain.Message{{Role: "user", Content: userInput}},
			Stream:   false,
		}
		chatResp, err := chatUsecase.Chat(chatReq)
		if err != nil {
			log.Printf("Gagal request ke AI: %v", err)
			continue
		}
		if len(chatResp.Choices) > 0 {
			fmt.Println("AI:", chatResp.Choices[0].Message.Content)
		} else {
			fmt.Println("Tidak ada balasan dari AI.")
		}
	}
}
