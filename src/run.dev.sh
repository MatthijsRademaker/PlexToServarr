#!/bin/bash

# Set variables
APP_NAME="Overwatch"  # Replace with your application name
APP_DIR="/usr/local/bin"  # Where the binary will be moved
ENV_FILE=".env.development"  # Path to your .env file

# Step 1: Load environment variables from the .env file
if [ -f $ENV_FILE ]; then
    echo "Loading environment variables from $ENV_FILE..."
    set -a  # Automatically export all variables
    source $ENV_FILE
    set +a
else
    echo "Error: $ENV_FILE file not found."
    exit 1
fi

# Step 2: Run the Go application (environment variables are already set)
go run .
