#!/bin/bash

# Wait for container dependencies to initialize
#sleep 5 # uncomment if needed

# Navigate to the working directory and start Delve to debug the main.go package
cd /var/www/go/ecommerce && \
dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient &>> /var/log/delve/ecommerce.log
