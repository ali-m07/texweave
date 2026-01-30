# Publishing the Go package

For Go, the **repo is the package**. There is no separate publish step (like `npm publish`).

## 1. It's already installable

Anyone can run:

```bash
go install github.com/ali-m07/texweave/cmd/texweave@latest
```

Because the repo is public, `@latest` uses the latest commit on `main`.

## 2. Create a release (recommended)

Releases give users a **version** to pin (e.g. `@v1.0.0` instead of `@latest`).

From the repo root:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Then create a GitHub Release (optional but nice):

- Go to **Releases** → **Create a new release**
- Choose tag `v1.0.0`
- Title: `v1.0.0`
- Description: short changelog or "Initial release"
- Publish

Or with GitHub CLI:

```bash
gh release create v1.0.0 --title "v1.0.0" --notes "Initial release"
```

Users can then install a specific version:

```bash
go install github.com/ali-m07/texweave/cmd/texweave@v1.0.0
```

## 3. pkg.go.dev (docs)

[pkg.go.dev](https://pkg.go.dev) indexes public Go modules when someone visits the module URL or runs `go get`. After you push a tag (e.g. `v1.0.0`), visit:

**https://pkg.go.dev/github.com/ali-m07/texweave**

If it's not there yet, click "Request" to trigger indexing. After that, the README badge and docs will work.

## 4. GitHub "Packages" tab

The **Packages** tab on GitHub is for npm, Docker, Maven, etc. Go modules are **not** published there. Your Go package is the repo itself; `go install` and `go get` use the repo directly.

## Summary

| Goal                    | What to do                                      |
|-------------------------|--------------------------------------------------|
| Others install CLI      | They run `go install ...@latest` (already works) |
| Pin to a version        | Create tag + release (e.g. `v1.0.0`)            |
| Docs on pkg.go.dev      | Push a tag, then open the module URL on pkg.go.dev |
| No extra "publish" step  | Repo = package for Go                           |
