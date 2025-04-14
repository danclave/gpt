package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func requestGPT(reqBody GPTRequest) GPTResponse {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY not set")
	}
	reqBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBytes))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var gptResp GPTResponse
	if err := json.Unmarshal(bodyBytes, &gptResp); err != nil {
		panic(err)
	}
	return gptResp
}
