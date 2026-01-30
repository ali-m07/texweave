# Push to GitHub

1. Create a new repository on GitHub (e.g. `texweave` under your account or org).
2. From the project root, run:
   ```bash
   ./scripts/publish.sh YOUR_GITHUB_USERNAME
   ```
   This sets the Go module path and `origin` remote.
3. Push:
   ```bash
   git push -u origin main
   ```
4. Update the install path in README if you use a different repo path (e.g. `github.com/YOUR_USERNAME/texweave`).
