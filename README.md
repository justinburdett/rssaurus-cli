# rssaurus-cli

A command-line client for [RSSaurus](https://rssaurus.com).

## Install

(TODO) GitHub Releases will provide prebuilt binaries for macOS/Linux/Windows.

## Quick start

1. Create an API token: https://rssaurus.com/api_tokens/new
2. Login:

```bash
rssaurus auth login
```

## Configuration

For v1 we keep a single host/token.

- Default host: `https://rssaurus.com`
- Override host via `RSSAURUS_HOST` or `--host`
- Provide token via `RSSAURUS_TOKEN` or store it locally via `rssaurus auth login`

## Commands (planned MVP)

- `rssaurus auth login`
- `rssaurus auth whoami`
- `rssaurus feeds`
- `rssaurus items` (unread by default; `--json` for scripting)
- `rssaurus read <id>` / `rssaurus unread <id>`
- `rssaurus save <url>` / `rssaurus unsave <id>`

## Development

```bash
go test ./...
```
