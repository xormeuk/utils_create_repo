package main

import (
	"fmt"
	"os"
	"os/exec"
)

func executeCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run create_repo.go <repoName>")
		os.Exit(1)
	}
	repoName := os.Args[1]
	username := os.Getenv("USER")
	token := os.Getenv("GITHUB_TOKEN") // Ensure you have GITHUB_TOKEN set in your environment

	if token == "" {
		fmt.Println("GitHub token not found in environment variable GITHUB_TOKEN")
		os.Exit(1)
	}

	// Create private repo on GitHub
	if err := executeCommand("curl", "-X", "POST", "-H", "Authorization: token "+token, "-d", `{"name":"`+repoName+`","private":true}`, "https://api.github.com/user/repos"); err != nil {
		fmt.Println("Failed to create repository on GitHub:", err)
		os.Exit(1)
	}

	// Create local directory and initialize git
	repoPath := "/home/" + username + "/dev/" + repoName
	if err := os.MkdirAll(repoPath, 0755); err != nil {
		fmt.Println("Failed to create local directory:", err)
		os.Exit(1)
	}

	if err := os.Chdir(repoPath); err != nil {
		fmt.Println("Failed to change directory:", err)
		os.Exit(1)
	}

	if err := executeCommand("git", "init"); err != nil {
		fmt.Println("Failed to initialize git:", err)
		os.Exit(1)
	}
	if err := executeCommand("git", "remote", "add", "origin", "https://github.com/"+username+"/"+repoName+".git"); err != nil {
		fmt.Println("Failed to add remote origin:", err)
		os.Exit(1)
	}

	fmt.Println("Repository", repoName, "created locally and on GitHub")
}
