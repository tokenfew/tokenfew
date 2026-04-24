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

package language_drive

import (
	"os"
	"runtime"
	"strings"
)

// GetSystemLanguage detects and returns the system language code (e.g., "en", "zh").
// It checks the OS type and delegates to the appropriate function for locale detection.
func GetSystemLanguage() string {
	rawLocale := ""       // Variable to store the raw locale string
	switch runtime.GOOS { // Check the operating system
	case "windows":
		rawLocale = GetWindowsLanguage() // Get Windows locale
	case "darwin":
		rawLocale = GetDarwinLanguage() // Get macOS locale
	default:
		rawLocale = GetUnixLanguage() // Get Unix/Linux locale
	}

	if rawLocale == "unknown" { // Fallback to English if unknown
		rawLocale = "en"
	}

	return extractLanguageCode(rawLocale) // Extract and return the language code
}

// GetUnixLanguage retrieves the system locale on Unix/Linux systems.
// It reads the LANG environment variable and returns the locale code.
func GetUnixLanguage() string {
	lang := os.Getenv("LANG") // Get LANG env var
	if lang == "" {
		return "unknown"
	}
	parts := strings.Split(lang, ".") // Remove encoding part if present
	return parts[0]
}

// extractLanguageCode extracts the language code from a locale string.
// For example, "en-US" -> "en".
func extractLanguageCode(locale string) string {
	locale = strings.TrimSpace(locale) // Remove leading/trailing spaces
	if locale == "" || locale == "unknown" {
		return "unknown"
	}

	locale = strings.ReplaceAll(locale, "_", "-") // Replace underscores with hyphens

	if idx := strings.Index(locale, "."); idx != -1 { // Remove encoding if present
		locale = locale[:idx]
	}

	parts := strings.Split(locale, "-") // Split by hyphen
	if len(parts) > 0 {
		return strings.ToLower(parts[0]) // Return the language part in lowercase
	}

	return "unknown"
}
