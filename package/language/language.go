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

package language

import (
	"github.com/tokenfew/tokenfew/package/language/language_drive"
	"github.com/tokenfew/tokenfew/package/language/language_message"
)

// Get is a global instance of language used for accessing language settings.
var Get = &Language{}

// Language struct holds internationalization settings for the application.
type Language struct {
	Language string
}

// New creates and returns a new language instance with default language set to English.
func New() *Language {
	return &Language{
		Language: "en", // Default language is English
	}
}

// GetSystemLanguage retrieves the system's default language setting.
func (l *Language) GetSystemLanguage() string {
	return language_drive.GetSystemLanguage()
}

// SetLanguage sets the language for the I18n instance based on the provided acceptLanguage string.
func (l *Language) SetLanguage(acceptLanguage string) *Language {
	// Set the language based on the first two characters of the acceptLanguage string.
	if len(acceptLanguage) >= 2 {
		l.Language = acceptLanguage[:2]
	}

	// Return the I18n instance for method chaining.
	return l
}

// Lang retrieves the localized string for a given code based on the current language setting.
func (l *Language) Lang(code int) string {
	// Retrieve the localized string based on the current language setting.
	string := language_message.GetEnLanguage(code)

	// Retrieve the localized string for the current language setting.
	if l.Language == "zh" {
		// Get the Chinese language string.
		string = language_message.GetZhLanguage(code)
	}

	// Return the localized string.
	return string
}
