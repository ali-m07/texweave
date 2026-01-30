# Push to GitHub

1. **Create the repo on GitHub**: [New repository](https://github.com/new) — name it `texweave`, leave it empty (no README/license).
2. **Set module path and remote** (if you use a different GitHub username, replace `ali-mansouri`):
   ```bash
   ./scripts/publish.sh ali-mansouri
   ```
   Or manually:
   ```bash
   go mod edit -module github.com/YOUR_USERNAME/texweave
   git remote add origin https://github.com/YOUR_USERNAME/texweave.git
   # if origin already exists: git remote set-url origin https://github.com/YOUR_USERNAME/texweave.git
   ```
3. **Push**:
   ```bash
   git push -u origin main
   ```
4. **Optional**: If you used a different repo path, update the install command in README (e.g. `github.com/YOUR_USERNAME/texweave`).
