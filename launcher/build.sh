#!/bin/sh
# Builds the Windows, macOS and Linux apps into ./dist (needs Go 1.22 or newer).
# The page inside the apps is ../app/index.html, so the web app and the apps always match.
set -e
cd "$(dirname "$0")"
cp ../app/index.html index.html
mkdir -p dist
export CGO_ENABLED=0 GOFLAGS=-trimpath
GOOS=linux   GOARCH=amd64 go build -ldflags "-s -w" -o dist/MouseKeyboardTest .
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H=windowsgui" -o "dist/Mouse and Keyboard Test.exe" .
GOOS=darwin  GOARCH=amd64 go build -ldflags "-s -w" -o dist/mkt-x86_64 .
GOOS=darwin  GOARCH=arm64 go build -ldflags "-s -w" -o dist/mkt-arm64 .
rm -f index.html
echo "Built into $(pwd)/dist"
