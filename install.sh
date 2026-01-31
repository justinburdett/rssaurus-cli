#!/usr/bin/env bash
set -euo pipefail

# RSSaurus CLI installer
# - Downloads the latest GitHub release asset for your OS/arch
# - Installs `rssaurus` into /usr/local/bin (or $INSTALL_DIR)

REPO="RSSaurus/rssaurus-cli"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

need() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "Missing dependency: $1" >&2
    exit 1
  }
}

need curl
need uname
need mktemp

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
  x86_64|amd64) ARCH=amd64 ;;
  arm64|aarch64) ARCH=arm64 ;;
  *)
    echo "Unsupported arch: $ARCH" >&2
    exit 1
    ;;
esac

if [[ "$OS" == "darwin" || "$OS" == "linux" ]]; then
  EXT=tar.gz
elif [[ "$OS" == "msys" || "$OS" == "mingw" || "$OS" == "cygwin" ]]; then
  echo "Windows shell environments are not supported by this installer. Download the .zip from GitHub Releases." >&2
  exit 1
else
  echo "Unsupported OS: $OS" >&2
  exit 1
fi

need tar

echo "Fetching latest release info for $REPO..." >&2
JSON="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest")"
TAG="$(printf '%s' "$JSON" | sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p' | head -n1)"
if [[ -z "$TAG" ]]; then
  echo "Could not determine latest tag." >&2
  exit 1
fi

ASSET="rssaurus_${TAG#v}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${TAG}/${ASSET}"

TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

echo "Downloading $URL..." >&2
curl -fsSL -o "$TMPDIR/$ASSET" "$URL"

tar -xzf "$TMPDIR/$ASSET" -C "$TMPDIR"

if [[ ! -f "$TMPDIR/rssaurus" ]]; then
  echo "Archive did not contain expected binary 'rssaurus'" >&2
  exit 1
fi

chmod +x "$TMPDIR/rssaurus"

# If INSTALL_DIR requires sudo, let the user know.
mkdir -p "$INSTALL_DIR"

if [[ ! -w "$INSTALL_DIR" ]]; then
  echo "Installing to $INSTALL_DIR requires sudo..." >&2
  sudo install -m 0755 "$TMPDIR/rssaurus" "$INSTALL_DIR/rssaurus"
else
  install -m 0755 "$TMPDIR/rssaurus" "$INSTALL_DIR/rssaurus"
fi

echo "Installed: $INSTALL_DIR/rssaurus" >&2

# Basic smoke check
"$INSTALL_DIR/rssaurus" --help >/dev/null 2>&1 || true

echo "Done. Try: rssaurus --help" >&2
