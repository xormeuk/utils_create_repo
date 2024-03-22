# create_repo

`create_repo` is a command-line tool written in Go, designed to streamline the process of creating new repositories on GitHub. It automates the creation of a remote repository on GitHub and sets up a local repository, including initializing it with git and setting the remote origin.

## Features

- Create a new private repository on GitHub.
- Initialize a local git repository.
- Set the remote origin of the local repository to the newly created GitHub repository.
- Push an initial commit to the remote repository.

## Prerequisites

Before you can use `create_repo`, make sure you have the following installed:

- [Go](https://golang.org/dl/) (1.17 or higher recommended)
- Git
- A GitHub account and a [Personal Access Token](https://github.com/settings/tokens) with repo permissions.

## Installation

### Option 1: From Source

Clone this repository and build the `create_repo` executable:

```bash
git clone https://github.com/xormeuk/create_repo.git
cd create_repo
go build -o create_repo create_repo.go
```

### Option 2: Download the Artifact

For convenience, you can also download the compiled binary directly from our GitHub Actions artifacts:

- Navigate to the Actions tab in this repository.
- Select the latest successful workflow run from the list.
- Scroll down to the Artifacts section at the bottom of the workflow run page.
- Click on the create_repo-binary artifact to download it.

### Usage

To use create_repo, you'll need to set your GitHub Personal Access Token as an environment variable:

```bash
export GITHUB_TOKEN="your_github_token_here"
```

Then, you can create a new repository by running:
```bash
./create_repo <GitHubUsername> <RepositoryName>
```

Replace <GitHubUsername> with your GitHub username and <RepositoryName> with the desired name for your new repository.
