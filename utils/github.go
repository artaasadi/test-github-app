package utils

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func InitGitHubClient(cfg Config) *github.Client {
	// Load the private key
	keyBytes, err := ioutil.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}

	// Create a GitHub client using the app's private key
	tr, err := ghinstallation.NewAppsTransport(http.DefaultTransport, int64(cfg.AppID), keyBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	client := github.NewClient(&http.Client{Transport: tr})

	// Get the installation token for the repository
	token, _, err := client.Apps.CreateInstallationToken(context.Background(), int64(cfg.InstallationID), nil)
	if err != nil {
		log.Fatalf("Failed to create installation token: %v", err)
	}

	// Create a GitHub client using the installation token
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.GetToken()},
	))
	client = github.NewClient(tc)
	return client
}
