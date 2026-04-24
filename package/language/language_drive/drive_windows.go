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

//go:build windows
// +build windows

package language_drive

import (
	"syscall"
	"unsafe"
)

// GetWindowsLanguage retrieves the system locale on Windows using kernel32.dll.
// Returns the locale name (e.g., "en-US") or "unknown" if detection fails.
func GetWindowsLanguage() string {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")       // Load kernel32.dll
	proc := kernel32.NewProc("GetUserDefaultLocaleName") // Get the procedure for locale name

	buf := make([]uint16, 85)                                                   // Windows locale name max length is 85
	ret, _, _ := proc.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf))) // Call the procedure
	if ret == 0 {                                                               // If call failed
		return "unknown"
	}
	return syscall.UTF16ToString(buf) // Convert UTF-16 buffer to string
}

// GetDarwinLanguage retrieves the system locale on macOS.
// It first checks the AppleLocale environment variable, then uses the defaults command.
func GetDarwinLanguage() string {
	return "unknown"
}
