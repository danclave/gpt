package gpt

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Request payload for OpenAI API
type GPTRequest struct {
	Model    string       `json:"model"`
	Messages []GPTMessage `json:"messages"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response payload from OpenAI API
type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func Ask(prompt string) string {
	reqBody := GPTRequest{
		Model: OMNI,
		Messages: []GPTMessage{
			{Role: "user", Content: prompt},
		},
	}
	gptResp := requestGPT(reqBody)

	jsonStr := gptResp.Choices[0].Message.Content
	return jsonStr
}

func AskShort(prompt string) string {
	reqBody := GPTRequest{
		Model: OMNI,
		Messages: []GPTMessage{
			{Role: "system", Content: "Response in a brief and consise manner"},
			{Role: "user", Content: prompt},
		},
	}
	gptResp := requestGPT(reqBody)

	jsonStr := gptResp.Choices[0].Message.Content
	return jsonStr
}

func AskStruct[T any](prompt string, example T) T {
	exampleBytes, _ := json.MarshalIndent(example, "", "  ")
	exampleStr := string(exampleBytes)

	systemMsg := "You are a service that converts unstructured input into structured JSON. " +
		"Use the following example as a template:\n\n" + exampleStr + "\n\nRespond only with valid JSON."

	reqBody := GPTRequest{
		Model: OMNI,
		Messages: []GPTMessage{
			{Role: "system", Content: systemMsg},
			{Role: "user", Content: prompt},
		},
	}
	gptResp := requestGPT(reqBody)

	jsonStr := gptResp.Choices[0].Message.Content

	// Strip code block formatting if present
	re := regexp.MustCompile("(?s)```(?:json)?(.*?)```")
	matches := re.FindStringSubmatch(jsonStr)
	if len(matches) > 1 {
		jsonStr = matches[1]
	}
	jsonStr = strings.TrimSpace(jsonStr)

	var result T
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}
