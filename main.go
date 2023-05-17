package main

import (
	"fmt"
	"os"

	"github.com/artaasadi/test-github-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v38/github"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var (
	client *github.Client
)

const (
	filepath = "config.yaml"
)

func main() {
	r := gin.New()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/github/payload", utils.ConsumeEvent)
		//v1.GET("/github/pullrequests/:owner/:repo", apis.GetPullRequests)
		//v1.GET("/github/pullrequests/:owner/:repo/:page", apis.GetPullRequestsPaginated)
	}
	var cfg utils.Config
	err := readFile(&cfg, filepath)
	if err != nil {
		readEnv(&cfg)
	}
	client = utils.InitGitHubClient(cfg)
	r.Run(fmt.Sprintf(":%v", 8000))
}

func readFile(cfg *utils.Config, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func readEnv(cfg *utils.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
}
