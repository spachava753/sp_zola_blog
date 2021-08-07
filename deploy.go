package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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

	commit := os.Getenv("CF_PAGES_COMMIT_SHA")

	if commit == "" {
		fmt.Println("env CF_PAGES_COMMIT_SHA is not set!")
		return
	}

	commit = commit[:7]

	url := fmt.Sprintf("https://%s.sp-zola-blog.pages.dev", commit)

	fmt.Printf("building with base url %s\n", url)

	if err := runCommand("zola", "build", "--base-rul", url); err != nil {
		fmt.Printf("found error while building site for branch build: %s\n", err)
		return
	}
}
