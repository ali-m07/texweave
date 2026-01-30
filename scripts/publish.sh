#!/bin/sh
# Replace YOUR_USERNAME with your GitHub username, then run:
#   ./scripts/publish.sh YOUR_USERNAME
set -e
USER="${1:?Usage: ./scripts/publish.sh YOUR_GITHUB_USERNAME}"
MODULE="github.com/${USER}/texweave"
cd "$(dirname "$0")/.."
go mod edit -module "$MODULE"
git add go.mod
git commit -m "chore: set module path for $MODULE" || true
git remote add origin "https://github.com/${USER}/texweave.git" 2>/dev/null || git remote set-url origin "https://github.com/${USER}/texweave.git"
git branch -M main
echo "Push with: git push -u origin main"
