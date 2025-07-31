#!/bin/bash

#Settings
PROGRAM_NAME="mango"
INSTALL_DIR="/usr/bin"

echo "[*] Fetching latest version tag from GitHub..."

LATEST_TAG=$(curl -s "https://api.github.com/repos/lexionq/mango/releases/latest" | grep '"tag_name":' | head -1 | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo "[!] Could not fetch latest version. Using default v0.3"
    LATEST_TAG="v0.3"
fi


BINARY_URL="https://github.com/lexionq/mango/releases/download/$LATEST_TAG/$PROGRAM_NAME"

echo "[*] Install starting..."

if [ "$EUID" -ne 0 ]; then
    sudo -v || exit 1
fi

if command -v curl >/dev/null 2>&1; then
    sudo curl -L "$BINARY_URL" -o "$INSTALL_DIR/$PROGRAM_NAME"
elif command -v wget >/dev/null 2>&1; then
    sudo wget -O "$INSTALL_DIR/$PROGRAM_NAME" "$BINARY_URL"
else
    echo "[!] curl or wget not found. Please install curl or wget."
    exit 1
fi

sudo chmod +x "$INSTALL_DIR/$PROGRAM_NAME"

echo "[✓] Installing finished."