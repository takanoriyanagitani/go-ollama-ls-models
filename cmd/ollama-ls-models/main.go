package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	lm "github.com/takanoriyanagitani/go-ollama-ls-models"
)

func main() {
	w := json.NewEncoder(os.Stdout) // Initialize JSON encoder using stdout.

	client, err := lm.EnvToClient()
	if err != nil {
		log.Fatalf("Failed to initialize Ollama client: %v", err)
	}

	ctx := context.Background()

	models, err := client.ListModels(ctx)
	if err != nil {
		log.Fatalf("Failed to retrieve model list: %v", err)
	}

	err = client.FullModelsToJsonWriter(ctx, models, w) // Write full model info as JSONL.
	if err != nil {
		log.Fatalf("Failed to write full models in JSONL format: %v", err)
	}
}
