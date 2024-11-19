#!/bin/bash

echo "Building"

# Remove all local images
docker image prune --all --force

# Stop all containers
docker stop $(docker ps -q)

# Remove all containers
docker rm -f $(docker ps -aq)


docker build --platform=linux/amd64 . -t bot

doctl registry login
docker tag bot:latest registry.digitalocean.com/hexlive/bot:latest
docker image push registry.digitalocean.com/hexlive/bot:latest