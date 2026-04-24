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

package language_message

// English language translations
var en = map[int]string{
	-1:    "Not logged in account",
	0:     "request successful",
	10000: "request error",
	10001: "illegal request",
}

// GetEnLanguage retrieves the English language string for a given code.
func GetEnLanguage(code int) string {
	return en[code]
}
