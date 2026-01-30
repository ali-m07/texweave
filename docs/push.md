# Push to GitHub

1. **Create the repo on GitHub**: [New repository](https://github.com/new) — name it `texweave`, leave it empty (no README/license).
2. **Set module path and remote** (repo under [ali-m07](https://github.com/ali-m07)):
   ```bash
   ./scripts/publish.sh ali-m07
   ```
   Or manually:
   ```bash
   go mod edit -module github.com/ali-m07/texweave
   git remote add origin https://github.com/ali-m07/texweave.git
   # if origin already exists: git remote set-url origin https://github.com/ali-m07/texweave.git
   ```
3. **Push**:
   ```bash
   git push -u origin main
   ```
4. Install path: `github.com/ali-m07/texweave` (already set in README).
