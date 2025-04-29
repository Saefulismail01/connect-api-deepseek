package http

import (
	"encoding/json"
	"net/http"
	"strings"
	"initialProject/internal/domain"
	"initialProject/internal/usecase"
)

func StartServer(chatUsecase *usecase.ChatUsecase) {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		var req domain.ChatAPIRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return
		}

		if strings.TrimSpace(req.Prompt) == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Prompt tidak boleh kosong"})
			return
		}

		chatReq := domain.ChatRequest{
			Model:    "deepseek-chat",
			Messages: []domain.Message{{Role: "user", Content: req.Prompt}},
			Stream:   false,
		}

		chatResp, err := chatUsecase.Chat(chatReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		if len(chatResp.Choices) > 0 {
			json.NewEncoder(w).Encode(domain.ChatAPIResponse{Response: chatResp.Choices[0].Message.Content})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "No response from AI"})
		}
	})

	http.ListenAndServe(":8080", nil)
}
