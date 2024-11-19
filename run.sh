#!/bin/bash

echo "Run"

docker run -e POSTGRESS_HOST='host.docker.internal' bot /app/cmd/start gatherPairs