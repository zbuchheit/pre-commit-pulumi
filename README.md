# Pre-Commit Pulumi

This repository contains a pre-commit plugin designed to prevent accidental commits of Pulumi state files into your version control system. By leveraging this plugin, you can ensure that sensitive information contained within Pulumi state files remains secure and out of your codebase.

## Prerequisites

1. [Pre-Commit](https://pre-commit.com/)
1. [Go](https://go.dev/)
1. [Git](https://git-scm.com/)

## Installation

Add Pre-commit Hook

In the root of your repository, create or update a .pre-commit-config.yaml file with the following content:
```yaml
todo:
```

## Usage

Usage
Once installed, the pre-commit hook will automatically run each time you attempt to commit changes. If a Pulumi state file is detected, the commit will be blocked, and an error message will be displayed.

To manually run the pre-commit checks on all files in the repository, you can use:

```zsh
pre-commit run --all-files
```
To run the checks on specific files, use:
```zsh
pre-commit run pulumi-state-check --files path/to/file.json
```

## Contributing

We welcome contributions and improvements to this plugin! Please feel free to fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the [Apache 2.0 License](LICENSE) - see the LICENSE file for details.

