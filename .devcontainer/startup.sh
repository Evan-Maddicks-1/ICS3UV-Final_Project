#!/bin/bash

# change welcome message
sudo cp ./.devcontainer/welcome.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt

# change the shell prompt
cp ./.devcontainer/.bashrc ~/.bashrc
# shellcheck disable=SC1090
source ~/.bashrc

# install Go linter
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.2.1
