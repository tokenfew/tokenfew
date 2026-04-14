#!/bin/bash

# Copyright 2025 GEEKROS, Inc.
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

# Print the current working directory
echo "Current working directory: $PWD"

# Get the directory of the current script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "Script directory: $SCRIPT_DIR"