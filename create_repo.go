package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// RepoRequest represents the JSON payload for creating a GitHub repository.
type RepoRequest struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func createGitHubRepo(token, username, repoName string) error {
	apiUrl := "https://api.github.com/user/repos"
	repo := RepoRequest{Name: repoName, Private: true}
	repoJson, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(repoJson))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to create repository: %s", string(body))
	}

	return nil
}

func executeGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: create_repo <GitHubUsername> <repoName>")
		os.Exit(1)
	}
	username := os.Args[1]
	repoName := os.Args[2]
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		fmt.Println("GitHub token not found in environment variable GITHUB_TOKEN")
		os.Exit(1)
	}

	// Create private repo on GitHub
	if err := createGitHubRepo(token, username, repoName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repoPath := "./" + repoName
	if err := os.MkdirAll(repoPath, 0755); err != nil {
		fmt.Println("Failed to create local directory:", err)
		os.Exit(1)
	}

	if err := os.Chdir(repoPath); err != nil {
		fmt.Println("Failed to change directory:", err)
		os.Exit(1)
	}

	if _, err := executeGitCommand("init"); err != nil {
		fmt.Println("Failed to initialize git:", err)
		os.Exit(1)
	}
	if _, err := executeGitCommand("remote", "add", "origin", "git@github.com:"+username+"/"+repoName+".git"); err != nil {
		fmt.Println("Failed to add remote origin:", err)
		os.Exit(1)
	}

	// Creating an initial commit
	if _, err := executeGitCommand("commit", "--allow-empty", "-m", "Initial commit"); err != nil {
		fmt.Println("Note: Failed to create initial commit. This can be normal if there are no files.")
	}

	// Dynamically identifying the default branch name
	defaultBranch, err := executeGitCommand("branch", "--show-current")
	if err != nil {
		fmt.Println("Failed to identify the default branch name:", err)
		os.Exit(1)
	}
	defaultBranch = strings.TrimSpace(defaultBranch)

	// Setting upstream branch
	if _, err := executeGitCommand("push", "--set-upstream", "origin", defaultBranch); err != nil {
		fmt.Println("Failed to set upstream branch:", err)
		os.Exit(1)
	}

	fmt.Println("Repository", repoName, "created locally and on GitHub using SSH URL, with upstream set")
}
