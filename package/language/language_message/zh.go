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

// Chinese language translations
var zh = map[int]string{
	-1:    "未登录账号",
	0:     "请求成功",
	10000: "请求错误",
	10001: "非法请求",
}

// GetZhLanguage retrieves the Chinese language string for a given code.
func GetZhLanguage(code int) string {
	return zh[code]
}
