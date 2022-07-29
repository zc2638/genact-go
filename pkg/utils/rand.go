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

package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/99nil/gopkg/sets"
)

func RandInputStr(data []string) string {
	if len(data) == 0 {
		return ""
	}
	set := sets.NewString(data...)
	for v := range set {
		return v
	}
	return ""
}

func RandInt(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func RandInt64(min, max int64) int64 {
	if min == max {
		return min
	}
	return rand.Int63n(max-min) + min
}

var hexStr = "0123456789abcdef"

func RandHashStr(length int) string {
	b := []byte(hexStr)
	result := make([]byte, 0, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

var stageSet = sets.NewString("alpha", "beta", "rc", "stable")

func RandVersion(only bool) string {
	major := RandInt(0, 9)
	minor := RandInt(0, 30)
	patch := RandInt(0, 20)
	version := fmt.Sprintf("%d.%d.%d", major, minor, patch)

	if !only {
		var stage string
		for v := range stageSet {
			stage = v
			break
		}
		version += "-" + stage

		var stageNumStr string
		stageType := RandInt(0, 1)
		if stageType > 0 {
			stageNum := RandInt(0, 10)
			stageNumStr = "." + strconv.Itoa(stageNum)
		}
		version += stageNumStr

		hasMarkSymbol := RandInt(0, 1)
		if hasMarkSymbol > 0 {
			version = "v" + version
		}
	}

	return version
}
