#!/bin/bash

# Set variables
APP_NAME="Overwatch"  # Replace with your application name
APP_DIR="/usr/local/bin"  # Where the binary will be moved
CONFIG_DIR="/etc/$APP_NAME"  # Configuration directory
ENVIRONMENT="production"  # Set the environment (can be changed to development, testing, etc.)

# Step 1: Build the Go application
echo "Building the Go application..."
go build -o $APP_NAME

# Step 2: Move the binary to the appropriate directory
echo "Installing the application binary..."
sudo mv $APP_NAME $APP_DIR

# Step 3: Move the correct config file to the config directory
echo "Setting up the configuration file..."
sudo mkdir -p $CONFIG_DIR
sudo cp config.$ENVIRONMENT.yml $CONFIG_DIR/config.yml

# Step 4: Create a systemd service file (optional, for auto startup)
SERVICE_FILE="/etc/systemd/system/$APP_NAME.service"

echo "Creating the systemd service file..."
sudo bash -c "cat > $SERVICE_FILE" << EOL
[Unit]
Description=$APP_NAME Service
After=network.target

[Service]
ExecStart=$APP_DIR/$APP_NAME
WorkingDirectory=$CONFIG_DIR
Restart=always
Environment=ENVIRONMENT=$ENVIRONMENT

[Install]
WantedBy=multi-user.target
EOL

# Step 5: Enable and start the service
echo "Enabling and starting the service..."
sudo systemctl daemon-reload
sudo systemctl enable $APP_NAME
sudo systemctl start $APP_NAME

echo "$APP_NAME setup is complete and running."
