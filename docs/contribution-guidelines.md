---
icon: lucide/users
---

# 🤝 Contribution Guidelines

Welcome to **[Green Book**! We're excited about your contribution. This guide will help you understand how to get started and ensure your work aligns with our project standards.

## 🎯 Why Contribute?

- Help shape a growing open-source ecosystem
- Improve the library through peer review
- Expand your own skills while contributing
- Build a network of fellow developers interested in this domain

---

## 1️⃣ Development Setup

> **Note:** This guide is being continuously updated. Please revisit this section as we expand it with detailed setup instructions!

### Prerequisites

To begin contributing, you'll need:

- ✅ Python 3.x (recommended version: 3.10+)
- ✅ git (latest stable version)
- ✅ A code editor (VS Code, Python IDE, or any lightweight editor)
- 💡 Optional but recommended: [`pre-commit`](https://pre-commit.com/) for automated code quality

### Step-by-Step Setup

```bash
# Clone the repository with all branches and submodules
git clone https://github.com/green-book/gb-go.git
cd gb-go

# Make sure you have main (development) and tags available
git checkout main

# Install dependencies via uv
uv venv
source .venv/bin/activate

# Pre-commit config for automatic linting/hints
pre-commit install
```

**Note:** As we expand the project, this setup guide will be updated to include:
- Detailed toolchain requirements
- Environment configuration tutorials
- Platform-specific guidelines

### Repository Structure

| Directory | Purpose |
|-----------|---------|
| `gb/` | Main package source code |
| `docs/` | Documentation and guides |
| `.github/` | CI/CD configuration and reusable actions |
| `tests/` | Test suites and fixtures |
| `scripts/` | Dev scripts, generators |

---

## 2️⃣ Git Commit Guidelines

> **Reference:** [Conventional Commits](https://conventionalcommits.org/)
> **Hook Status:** Enforced via [`pre-commit`](https://github.com/compilerla/conventional-pre-commit) framework (version 4.1.0)

### 📖 The Standard

**Green Book** follows the [Conventional Commits](https://conventionalcommits.org/) specification, a lightweight convention for adding semantic meaning to commit messages. This helps automated tools parse your changelog and generate release notes.

### ✨ Commit Message Format

Each commit message should consist of:
1. **Type**: What kind of change? (`feat`, `fix`, `docs`, etc.)
2. **Scope** (Optional): Where did the change occur? (`core`, `api`, `ui`, etc.)
3. **Description**: A concise space-separated summary

#### Format Pattern

```text
<type>(optional: <scope>): <description>
```

### ✅ Valid Examples

| Type | Example Message | Description |
|------|-----------------|-------------|
| `feat` | `feat(core): add new feature X` | A new feature being added |
| `fix` | `fix(api): resolve authentication bug` | Bug fix or patch |
| `docs` | `docs: update installation guide` | Documentation changes only |
| `refactor` | `refactor(utils): improve performance` | Code refactoring (no behavior change) |
| `style` | `style(core): fix whitespace and formatting` | Formatting/style changes only |
| `perf` | `perf(image): optimize loading speed` | Performance improvements |
| `test` | `test(api): add unit tests for new endpoint` | Adding or updating tests |
| `chore` | `chore(deps): update dependencies to latest versions` | Chore (configuration, dependencies, etc.) |

### ✅ More Complex Examples

```text
feat(core): introduce plugin manager for extensibility
refactor(image): optimize compression efficiency by 40%
fix(ui): resolve dropdown layout on mobile devices
build: prepare v1.2.0 release artifacts
chore(deps): upgrade to Python 3.12 compat
perf(storage): reduce memory footprint of image buffers
```

### ⚠️ What Won't Pass Validation

❌ `add new feature` – Missing type and scope
❌ `fix a bug` – No capitalization, period at end, or clear scope
❌ `docs: update readme` – Lowercase first letter after colon

---

## 🏪 Contributing Workflow

Use these branches for contributions:

```bash
# Create your feature branch
git checkout -b feat/my-feature-name main

# Or create a bugfix branch
git checkout -b fix/issue-some-description main

# Make changes and commit (follow guidelines above!)
git add .
git commit -m "feat(core): implement my new feature"

# Push upstream and open a pull request
git push origin feat/my-feature-name
```

---

## 📋 Pull Request Template

When submitting your PR, include:

- ✅ A clear title summarizing the changes
- ✅ The issue/feature number (e.g., `feat(core): #23`)
- ✅ Testing performed and results
- ✅ Screenshots of UI changes or API documentation

### Pre-submission checklist

```markdown
- [ ] Code follows project conventions
- [ ] New code is properly formatted (pre-commit hooks pass)
- [ ] All tests are passing locally
- [ ] Documentation updated for new features
- [ ] PR linked to an existing issue/PR
```

---

## 🤔 Questions?

Feel free to:

- **Open an issue**: Discuss any questions or problems before they reach us in a PR.
- **Join the discussion thread**: We're active on GitHub Discussions!
- **Email us**: contact@greensignal.com

Thanks for contributing to **Green Book**! 🚀

---

**Generated with ❤️ by the Green Book Team**
[![Conventional Commits](https://conventionalcommits.org/badge.svg)](https://conventionalcommits.org/)
