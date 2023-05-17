package utils

import (
	"context"
	"log"

	"github.com/google/go-github/v38/github"
)

func readContent(client *github.Client, fileRoute string, cfg Config) string {
	// Get the contents of a file in the repository
	fileContent, _, _, err := client.Repositories.GetContents(context.Background(), cfg.RepoOwner, cfg.RepoName, fileRoute, nil)
	if err != nil {
		log.Fatalf("Failed to get file contents: %v", err)
	}

	// Print the contents of the file
	content, err := fileContent.GetContent()
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	return content
}
