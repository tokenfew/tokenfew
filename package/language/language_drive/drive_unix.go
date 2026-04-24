// Copyright 2025 TOKENFEW, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows
// +build !windows

package language_drive

import (
	"os"
	"os/exec"
	"strings"
)

// GetWindowsLanguage retrieves the system locale on Windows using kernel32.dll.
// Returns the locale name (e.g., "en-US") or "unknown" if detection fails.
func GetWindowsLanguage() string {
	return "unknown"
}

// GetDarwinLanguage retrieves the system locale on macOS.
// It first checks the AppleLocale environment variable, then uses the defaults command.
func GetDarwinLanguage() string {
	if locale := os.Getenv("AppleLocale"); locale != "" { // Check AppleLocale env var
		return locale
	}

	out, err := exec.Command("defaults", "read", "NSGlobalDomain", "AppleLanguages").Output() // Run defaults command
	if err == nil {
		str := string(out)
		lines := strings.Split(str, "\n")
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed != "(" && trimmed != ")" && trimmed != "" {
				return strings.Trim(trimmed, "\",") // Return the first valid language
			}
		}
	}

	return GetUnixLanguage() // Fallback to Unix method
}
