#! /bin/bash

docker run \
    --rm \
    -e LOG_LEVEL="debug" \
    -e GITHUB_ACCESS_TOKEN="$GITHUB_ACCESS_TOKEN" \
    -v $(pwd)/config.js:/usr/src/app/config.js \
    renovate/renovate:32.39.0
