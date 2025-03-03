package ai

import (
	"context"
	
	"log"
	

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	
	
	
)

var aiClient *genai.Client

func InitAI() {
	apiKey := "AIzaSyD4YYxqscqMQIvSThZclMKqfedLHY2cA0c" 
	if apiKey == "" {
		log.Fatal("Missing Gemini AI API Key")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("Failed to initialize AI client:", err)
	}
	aiClient = client
}

func GetAISuggestions(tasks string) (string, error) {
	

	return "Under Progress", nil
}