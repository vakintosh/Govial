# Govial

Govial is a DevOps-focused shell prompt written in Go. Inspired by [Jovial](https://github.com/zthxxx/jovial) and [Starship](https://starship.rs/), Govial aims to provide a feature-rich and highly customizable prompt for daily usage.

## Features

- **Cloud Provider and Region Detection**: Automatically displays the current cloud provider and region.
- **Language and Environment Detection**: Detects the programming language in the current folder and shows the version and virtual environment.
- **Terraform Workspace Detection**: Identifies the current Terraform workspace.
- **Dynamic Git Branch**: Displays the current Git branch if the folder is a Git repository.
- **Customizable Prompt Structure**: The prompt structure includes:
  - Current Hour
  - Git Branch (if applicable)
  - Shortened Path
  - Cloud Profile
  - Language Version and Virtual Environment/Workspace
  - Project Initialization

## Prompt Structure

The prompt is structured as follows:

    [Hour] - (git branch) - shortened path - cloud profile - language version + venv/workspace

## Installation

To install Govial, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/vakintosh/govial.git
    ```
2. Navigate to the project directory:
    ```sh
    cd govial
    ```
3. Build the project:
    ```sh
<!-- go mod init govial -->
<!-- go mod tidy -->
<!-- go run . -->
    go build
    ```
4. Add the binary to your shell's PATH.

## Usage

Once installed, you can start using Govial by adding it to your shell configuration file (e.g.,`.zshrc`).

```sh
# Example for .zshrc
export PATH=$PATH:/path/to/govial
```

## Project Initialization
Govial also provides a shortcut to initialize different types of projects. For example, to initialize a Terraform project with the current version, a custom pre-commit hook, and a custom .gitignore, you can use the following command:

```sh
govial init tf
```

This command will set up the necessary files and configurations for a Terraform project. ```
