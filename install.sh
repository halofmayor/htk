#!/bin/bash
set -e

# -----------------------
# Project configurations
# -----------------------
GITHUB_USER="halofmayor"
REPO="htk"
VERSION="v1.1-alpha"

# Detect OS
OS=$(uname -s)
ARCH=$(uname -m)

if [[ "$OS" == "Linux" ]]; then
    OS_NAME="linux"
elif [[ "$OS" == "Darwin" ]]; then
    OS_NAME="macos"
else
    echo "System not supported. Only Linux and macOS"
    exit 1
fi

# Detect architecture
if [[ "$ARCH" == "x86_64" || "$ARCH" == "amd64" ]]; then
    ARCH_NAME="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
    ARCH_NAME="arm64"
else
    echo "Architecture not supported: $ARCH"
    exit 1
fi

# ---------------------
# Binary selection
# ---------------------
# Certifique-se de que os binários na Release tenham estes nomes:
# htk-linux-amd64, htk-macos-amd64, htk-macos-arm64
BIN_NAME="htk-$OS_NAME-$ARCH_NAME"
URL="https://github.com/$GITHUB_USER/$REPO/releases/download/$VERSION/$BIN_NAME"

# ---------------------
# Download and install
# ---------------------
echo "Downloading HTK from $URL ..."
curl -L -o /tmp/htk "$URL"

# Verifica se o download produziu um binário válido
if ! file /tmp/htk | grep -qE 'executable'; then
    echo "Downloaded file is not a valid binary. Please check the Release."
    exit 1
fi

chmod +x /tmp/htk
sudo mv /tmp/htk /usr/local/bin/htk

echo "--------------"
echo "HTK installed successfully!"
