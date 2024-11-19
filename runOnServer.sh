#!/bin/bash


watchtower_names_to_watch=""
declare -g watchtower_names_to_watch=""

echo "Ron the bot"

doctl registry login

# Remove all local images
docker image prune --all --force

# Stop all containers
docker stop $(docker ps -q)

# Remove all containers
docker rm -f $(docker ps -aq)

# Pull latest
docker pull registry.digitalocean.com/hexlive/bot:latest


docker run -d \
        --restart unless-stopped \
        --name=gatherPairs \
        -l com.centurylinklabs.watchtower.enable=true \
        --env-file=/opt/bot.env \
        --log-opt max-size=100m \
        registry.digitalocean.com/hexlive/bot:latest /app/cmd/start gatherPairs

#docker run -d \
#        --restart unless-stopped \
#        --name=findSimpleArb \
#        -l com.centurylinklabs.watchtower.enable=true \
#        --env-file=/opt/bot.env \
#        --log-opt max-size=100m \
#        registry.digitalocean.com/hexlive/bot:latest /app/cmd/start findSimpleArb

docker run \
    -d \
    --name watchtower \
    --log-opt max-size=100m \
    --restart unless-stopped \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /home/paul/.docker/config.json:/config.json \
    containrrr/watchtower:1.5.3 \
    --cleanup \
    --include-stopped \
    --include-restarting \
    --interval 30 \
    gatherPairs
