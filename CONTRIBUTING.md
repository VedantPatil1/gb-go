# Contributing

## Commit Message Conventions

This project follows [conventional commits](https://www.conventionalcommits.org/) to maintain clean and consistent commit history.

### Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer]
```

### Types

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `build`: Changes that affect the build system or external dependencies
- `ci`: Changes to our CI configuration files and scripts
- `chore`: Other changes that don't modify src or test files
- `revert`: Reverts a previous commit

### Examples

```
feat: add new API endpoint

fix: resolve null pointer exception in login

docs: update installation instructions

style: format code according to project standards

refactor: extract user validation logic to separate function

perf: optimize database query for user list

test: add unit tests for payment processing

build: update dependencies to latest versions

ci: add coverage reporting to CI pipeline

chore: update copyright year in all files

revert: revert commit 42b1b8f
```

## Pre-commit Setup

To ensure commit messages follow conventions, set up pre-commit hooks:

1. Install pre-commit:
```bash
pip install pre-commit
```

2. Install the hooks:
```bash
pre-commit install
```

The pre-commit hook will validate your commit messages and prevent commits that don't follow the conventional commit format.
