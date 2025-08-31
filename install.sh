#!/bin/bash
set -e

# -----------------------
# Project configurations
# -----------------------
GITHUB_USER="halofmayor"
REPO="htk"
VERSION="v1.0-alpha"

# Detects OS
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

# Detects architecture
if [[ "$ARCH" == "x86_64" || "$ARCH" == "amd64" ]]; then
    ARCH_NAME="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
    ARCH_NAME="arm64"
else
    echo "Architecture not supported: $ARCH"
    exit 1
fi

# Binary name 
BIN_NAME="htk-$OS_NAME-$ARCH_NAME"

# Binary URL
URL="https://github.com/$GITHUB_USER/$REPO/releases/download/$VERSION/$BIN_NAME"


# ---------------------
# Download and install
# ---------------------
echo "Downloading HTK..."
curl -L -o /tmp/htk "$URL"

chmod +x /tmp/htk

sudo mv /tmp/htk /usr/local/bin/htk

echo "--------------"
echo "HTK installed"
