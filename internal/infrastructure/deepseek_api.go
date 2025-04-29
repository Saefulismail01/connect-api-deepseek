package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"initialProject/internal/domain"
	"io/ioutil"
	"net/http"
)

type DeepseekAPI struct {
	apiKey string
}

func NewDeepseekAPI(apiKey string) *DeepseekAPI {
	return &DeepseekAPI{apiKey: apiKey}
}

func (d *DeepseekAPI) SendChat(request domain.ChatRequest) (domain.ChatResponse, error) {
	url := "https://api.deepseek.com/chat/completions"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return domain.ChatResponse{}, fmt.Errorf("gagal encode JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return domain.ChatResponse{}, fmt.Errorf("gagal membuat request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+d.apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.ChatResponse{}, fmt.Errorf("gagal melakukan request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.ChatResponse{}, fmt.Errorf("gagal membaca response: %v", err)
	}

	var chatResp domain.ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return domain.ChatResponse{}, fmt.Errorf("gagal parsing response JSON: %v", err)
	}

	return chatResp, nil
}
