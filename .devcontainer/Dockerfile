# Use the official Dev Containers Go image (includes Go and common tooling)
FROM mcr.microsoft.com/vscode/devcontainers/go:1

# Install apt packages we need (non-interactive), keep image small
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    build-essential \
    postgresql \
    postgresql-contrib \
 && rm -rf /var/lib/apt/lists/*

# Set up GOPATH/GOBIN and ensure the vscode user owns the directory
ENV GOPATH=/home/vscode/go
ENV GOBIN=/home/vscode/go/bin
ENV PATH=${GOBIN}:/usr/local/go/bin:${PATH}
RUN mkdir -p ${GOPATH} ${GOBIN} && chown -R vscode:vscode /home/vscode

# Switch to vscode user to run go install so binaries end up owned by vscode
USER vscode
WORKDIR /home/vscode

# Install CLI tools used by your project and put them in GOBIN
RUN go install github.com/bootdotdev/bootdev@latest \
 && go install github.com/pressly/goose/v3/cmd/goose@latest \
 && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Keep root for any remaining operations (devcontainer will set remoteUser)
USER root