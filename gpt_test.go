package gpt

import (
	"fmt"
	"testing"
)

func Integration_TestGptIntegration(t *testing.T) {
	prompt := `Respond ONLY with JSON. Provide a person with name, age, and job.`

	reqBody := GPTRequest{
		Model: "gpt-4",
		Messages: []GPTMessage{
			{Role: "system", Content: "You are a service that returns JSON only."},
			{Role: "user", Content: prompt},
		},
	}
	gptResp := requestGPT(reqBody)

	jsonStr := gptResp.Choices[0].Message.Content
	fmt.Print(jsonStr)
}

func Integration_TestAskJson(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Job  string `json:"job"`
	}

	example := Person{}
	prompt := `Bob, 26 years old, construction worker`

	result := AskStruct(prompt, example)

	person := result

	if person.Name == "" || person.Job == "" || person.Age == 0 {
		t.Errorf("Incomplete or empty fields in result: %+v", person)
	} else {
		t.Logf("Received person: %+v", person)
	}
}
