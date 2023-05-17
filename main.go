package main

import (
	"context"
	"crypto/rsa"
	"log"

	"github.com/artaasadi/test-github-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v38/github"
)

var (
	privateKey *rsa.PrivateKey
	client     *github.Client
)

const (
	repoOwner = "artaasadi"
	repoName  = "trader-bot"
)

func main() {
	r := gin.New()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/github/payload", utils.ConsumeEvent)
		//v1.GET("/github/pullrequests/:owner/:repo", apis.GetPullRequests)
		//v1.GET("/github/pullrequests/:owner/:repo/:page", apis.GetPullRequestsPaginated)
	}
	client = utils.InitGitHubClient()
	//r.Run(fmt.Sprintf(":%v", 8080))
}

func readContent(fileRoute string) string {
	// Get the contents of a file in the repository
	fileContent, _, _, err := client.Repositories.GetContents(context.Background(), repoOwner, repoName, fileRoute, nil)
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
