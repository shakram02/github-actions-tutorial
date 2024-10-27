

This repository demonstrates a simple Go web server deployed automatically using GitHub Actions. The server runs on a remote server, managed by `systemd` to ensure it stays running and restarts automatically on updates.

## Features

- **Go Server**: A lightweight HTTP server written in Go.
- **GitHub Actions CI/CD**: Automates deployment to a remote server on each push to the `main` branch.
- **systemd**: Manages the server process, restarts it on failure, and ensures it runs on boot.

## Prerequisites

- **Server** with SSH access, Go installed, and `systemd` enabled (most Linux distributions have this by default).
- **GitHub Secrets** for secure SSH access (see below).

## Setup Instructions

```
# Modify the system service file
sudo nano /etc/systemd/system/go-demo-server.service

sudo systemctl daemon-reload
sudo systemctl start go-demo-server
sudo systemctl enable go-demo-server
```