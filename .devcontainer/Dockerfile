FROM mcr.microsoft.com/devcontainers/base:ubuntu-24.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    wget \
    ca-certificates \
    build-essential \
    sudo \
    mysql-client \
    && apt-get clean && rm -rf /var/lib/apt/lists/*
# Ensure user vscode exists on all platforms
RUN id -u vscode 2>/dev/null || \
    (useradd -m -s /bin/bash vscode && echo "vscode ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers)

WORKDIR /workspace
