# Pre-Commit Pulumi

This repository contains a pre-commit plugin designed to prevent accidental commits of Pulumi state files into your version control system. By leveraging this plugin, you can ensure that sensitive information contained within Pulumi state files remains secure and out of your codebase.

## Prerequisites

1. [Pre-Commit](https://pre-commit.com/)
1. [Go](https://go.dev/)
1. [Git](https://git-scm.com/)

## Installation

1. Add Pre-commit Hook Config File

    In the root of your repository, create or update a `.pre-commit-config.yaml` file with the following content:

    ```yaml
    default_install_hook_types: 
    # - pre-push # if you want to run before a push
    - pre-commit
    repos:
    -   repo: https://github.com/zbuchheit/pre-commit-pulumi
        rev: v0.0.1
        hooks:
        -   id: pulumi-state-check
            stages: [pre-commit] #add pre-push if desired
    ```

1. Run `pre-commit install`
    ```bash
    $ pre-commit install
    pre-commit installed at .git/hooks/pre-commit
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

> [!NOTE] 
> If you commit often and only desire to check on push you can change add `pre-push` to `default_install_hook_types` and `stages`

## Contributing

We welcome contributions and improvements to this plugin! Please feel free to fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the [Apache 2.0 License](LICENSE) - see the LICENSE file for details.

