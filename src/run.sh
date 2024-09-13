#!/bin/bash

# Set variables
APP_NAME="Overwatch"  # Replace with your application name
APP_DIR="/usr/local/bin"  # Where the binary will be moved
ENV_FILE=".env"  # Path to your .env file

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

# Step 2: Build the Go application
echo "Building the Go application..."
go build -o $APP_NAME

# Step 3: Move the binary to the appropriate directory
echo "Installing the application binary..."
sudo mv $APP_NAME $APP_DIR

# Step 4: Run the application (environment variables are already set)
echo "Starting the application using environment variables..."
$APP_DIR/$APP_NAME
