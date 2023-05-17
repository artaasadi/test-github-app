package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/artaasadi/test-github-app/utils"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

const (
	appID          = 334984       // Replace with your app ID
	installation   = 37605233     // Replace with your installation ID
	repoOwner      = "artaasadi"  // Replace with the owner of the repository
	repoName       = "trader-bot" // Replace with the name of the repository
	privateKeyFile = "arta-test-github-app.2023-05-17.private-key.pem"
)

var (
	privateKey *rsa.PrivateKey
)

func main() {
	r := gin.New()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/github/payload", utils.ConsumeEvent)
		//v1.GET("/github/pullrequests/:owner/:repo", apis.GetPullRequests)
		//v1.GET("/github/pullrequests/:owner/:repo/:page", apis.GetPullRequestsPaginated)
	}
	utils.InitGitHubClient()
	r.Run(fmt.Sprintf(":%v", 8080))
}

func readContent(fileRoute string) string {
	// Load the private key
	keyBytes, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}
	// privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	// if err != nil {
	// 	log.Fatalf("Failed to parse private key: %v", err)
	// }

	// Create a GitHub client using the app's private key
	tr, err := ghinstallation.NewAppsTransport(http.DefaultTransport, appID, keyBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	client := github.NewClient(&http.Client{Transport: tr})

	// Get the installation token for the repository
	token, _, err := client.Apps.CreateInstallationToken(context.Background(), installation, nil)
	if err != nil {
		log.Fatalf("Failed to create installation token: %v", err)
	}

	// Create a GitHub client using the installation token
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.GetToken()},
	))
	client = github.NewClient(tc)

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
