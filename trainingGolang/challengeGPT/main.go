package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trainingGolang/challengeGPT/models"
)

// Replace "YOUR_OPENAI_API_KEY" with your actual API key
const apiKey = "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"

// const apiKey = "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"
const endpoint = "https://api.openai.com/v1/chat/completions"

func main() {
	var question string
	fmt.Printf("Pertanyaan: ")
	fmt.Scan(&question)

	// Example conversation
	conversation := []models.Message{
		{"system", "You are a helpful assistant."},
		{"user", "Who is the president of USA"},
		{"system", "The president of USA is Joe Biden"},
		{"user", question}, // Add more messages as needed
	}

	// Create a request
	requestData := models.GPTRequest{
		Model:     "gpt-3.5-turbo",
		Messages:  conversation,
		MaxTokens: 150,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding request:", err)
		return
	}

	// Make the HTTP request
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	// defer response.Body.Close()

	// Read and parse the response
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var responseData models.GPTResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the generated message
	fmt.Println("Assistant:", response) //responseData.Choices[0].Message.Content)
}
