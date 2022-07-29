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

package actuator

import (
	"context"

	"github.com/99nil/gopkg/sets"

	"github.com/zc2638/genact-go/pkg/utils"
)

type Interface interface {
	Name() string
	Execute(ctx context.Context, manifest map[string][]string)
}

func getCurrentData(data []string, min, max int) []string {
	set := sets.NewString(data...)
	if min < 0 {
		return set.List()
	}

	nameNum := utils.RandInt(min, max)
	currentSet := make([]string, 0, nameNum)

	var count int
	for v := range set {
		count++
		if count > nameNum {
			break
		}
		currentSet = append(currentSet, v)
	}
	return currentSet
}

func getCurrentSet(data []string, min, max int) sets.String {
	set := sets.NewString(data...)
	if min < 0 {
		return set
	}

	nameNum := utils.RandInt(min, max)
	currentSet := sets.NewString()

	for v := range set {
		if currentSet.Len() >= nameNum {
			break
		}
		currentSet.Add(v)
	}
	return currentSet
}

var extensions = []string{
	"gif", "mkv", "webm", "mp4", "html", "php", "md", "png", "jpg", "opus", "ogg", "mp3", "flac",
	"iso", "zip", "rar", "tar.gz", "tar.bz2", "tar.xz", "tar.zstd", "deb", "rpm", "exe",
}
