// Copyright © 2022 zc2638 <zc2638@qq.com>.
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

package static

import (
	"embed"
	"path/filepath"
	"strings"
)

//go:embed data
var data embed.FS

func Data() (map[string][]string, error) {
	files, err := data.ReadDir("data")
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string, len(files))
	for _, f := range files {
		filename := f.Name()
		currentPath := "data/" + filename
		fileData, err := data.ReadFile(currentPath)
		if err != nil {
			return nil, err
		}

		arr := strings.Split(string(fileData), "\n")
		set := make([]string, 0, len(arr))
		for _, v := range arr {
			if strings.TrimSpace(v) == "" {
				continue
			}
			set = append(set, v)
		}

		ext := filepath.Ext(filename)
		result[strings.TrimSuffix(filename, ext)] = set
	}
	return result, nil
}
