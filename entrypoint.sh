#!/bin/sh

set -e

sleep 5

#Run migrations
/app/migrate

# Run app
/app/main