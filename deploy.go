package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type Deployments struct {
	Results []struct {
		URL string `json:"url"`
	} `json:"result"`
}

func getLatestDeployment(accountId, email, apiToken string) (string, error) {
	var url string
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/pages/projects/sp-zola-blog/deployments",
			accountId),
		nil)
	if err != nil {
		return url, fmt.Errorf("found error while creating request for fetching deployments: %w", err)
	}

	req.Header.Add("X-Auth-Email", email)
	req.Header.Add("X-Auth-Key", apiToken)

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		return url, fmt.Errorf("could not complete request: %w", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return url, fmt.Errorf("could not read response body: %w", err)
	}
	var deployments Deployments
	if err := json.Unmarshal(respBody, &deployments); err != nil {
		return url, fmt.Errorf("could not unmarshal deployments: %w", err)
	}
	url = deployments.Results[0].URL
	return url, nil
}

func main() {
	if os.Getenv("CF_PAGES") == "1" {
		fmt.Println("Deploy site in Cloudflare pages")
	} else {
		fmt.Printf("Serving site locally using zola...\n\n")
		if err := runCommand("zola", "serve"); err != nil {
			fmt.Printf("found error while serving site: %s\n", err)
			return
		}
		return
	}

	branch := os.Getenv("CF_PAGES_BRANCH")

	if branch == "" {
		fmt.Println("env CF_PAGES_BRANCH is not set!")
		return
	}

	if branch == "main" {
		fmt.Println("deploying to production")
		if err := runCommand("zola", "build"); err != nil {
			fmt.Printf("found error while building site for production: %s\n", err)
			return
		}
		return
	}

	fmt.Println("deploying to branch build")

	accountId := os.Getenv("ACCOUNT_ID")
	if accountId == "" {
		fmt.Printf("ACCOUNT_ID env is missing")
		return
	}

	email := os.Getenv("EMAIL")
	if email == "" {
		fmt.Printf("EMAIL env is missing")
		return
	}

	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		fmt.Printf("API_TOKEN env is missing")
		return
	}

	url, err := getLatestDeployment(accountId, email, apiToken)
	if err != nil {
		fmt.Printf("could not get latest deployment url: %s\n", err)
		return
	}

	fmt.Printf("building with base url %s\n", url)

	if err := runCommand("zola", "build", "--base-url", url); err != nil {
		fmt.Printf("found error while building site for branch build: %s\n", err)
		return
	}
}
