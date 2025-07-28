#!/bin/bash

#Settings
PROGRAM_NAME="mango"
VERSION="v0.2"
INSTALL_DIR="/usr/bin"
BINARY_URL="https://github.com/lexionq/mango/releases/download/$VERSION/$PROGRAM_NAME"

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