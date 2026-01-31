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

Config file location:
- `$XDG_CONFIG_HOME/rssaurus/config.json` (if set)
- otherwise `~/.config/rssaurus/config.json`

## Commands (planned MVP)

- `rssaurus auth login`
- `rssaurus auth whoami`
- `rssaurus feeds`
- `rssaurus items` (unread by default; supports `--feed-id`, `--status`, `--limit`, `--cursor`)
- `rssaurus open <url>`
- `rssaurus read <item-id>` / `rssaurus unread <item-id>` (IDs available via `--json` output)
- `rssaurus mark-read --all` (or `--ids 1,2,3`, optional `--feed-id`)
- `rssaurus save <url>` / `rssaurus unsave <saved-item-id>` (IDs available via `--json` output)

## Development

```bash
go test ./...
```
