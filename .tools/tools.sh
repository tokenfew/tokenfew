#!/bin/bash

# Copyright 2025 TOKENFEW, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Exit immediately if a command exits with a non-zero status
set -e && lsb_release -a

# Get the system's architecture
architecture=$(dpkg --print-architecture)

# Get the Ubuntu version code name
ubuntu_code=$(lsb_release -c -s)

# Get the system's architecture name
arch=$(uname -m)

# Current script directory
current_dir=$(cd "$(dirname "$0")" && pwd)

# Temporary directory for the installation files
temp_dir=$(mktemp -d)

# Function to deploy the initial environment
on_init(){
    echo "Deploying the initial environment..."

    # Update the package list and install necessary tools
    sudo apt -y update && sudo apt install -y dpkg-dev gpg

    # If Go is not installed, install it
    if [ ! -d "/usr/local/go/bin/" ]; then
        # Install Go if not already installed
        golang_version="1.24.5"
        # Download and install Go
        sudo wget -q https://golang.google.cn/dl/go"${golang_version}".linux-"${architecture}".tar.gz && sudo tar -C /usr/local -xzf go"${golang_version}".linux-"${architecture}".tar.gz
        # Touch the Go profile script to ensure it exists
        touch /etc/profile.d/tokenfew_golang.sh
        # Add Go to the PATH environment variable
        sudo sh -c 'echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile.d/tokenfew_golang.sh'
        # Source the Go profile script to apply changes
        source /etc/profile.d/tokenfew_golang.sh
        # Clean up the downloaded Go tarball
        sudo rm -rf go"${golang_version}".linux-"${architecture}".tar.gz
    fi
}

# Function to publish the package to the repository
on_publish(){
    echo "Publishing the package to the repository..."

    # Extract the version number from the version.go file
    version=$(grep 'Number = .*"' "../package/version/version.go" | awk -F'"' '{print $2}')

    # Generate a timestamp for the package version
    datetime=$(date +%Y%m%d%H%M%S)

    # Log the version and timestamp for debugging purposes
    echo "Version: $version Timestamp: $datetime"
}

# Main script logic
case "$1" in
    "init")
        # Deploy the initial environment
        on_init
        ;;
    "publish")
        # Publish the package to the repository
        on_publish
        ;;
    *)
        echo "Usage: $0 {init}"
        exit 1
        ;;
esac