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
	"fmt"
	"time"

	"github.com/zc2638/genact-go/pkg/utils"
)

func NewCargo() Interface {
	return &cargo{}
}

type cargo struct{}

func (a *cargo) Name() string {
	return "cargo"
}

func (a *cargo) Execute(ctx context.Context, manifest map[string][]string) {
	data := manifest["packages"]
	if len(data) == 0 {
		return
	}

	currentData := getCurrentData(data, 10, 100)
	currentFullData := make([]string, 0, len(currentData))
	for _, name := range currentData {
		version := utils.RandVersion(true)
		currentFullData = append(currentFullData, fmt.Sprintf("%s v%s", name, version))
	}

	now := time.Now()
	for _, stage := range []string{"Downloading", "Compiling"} {
		for _, v := range currentFullData {
			// TODO stage add color
			fmt.Printf("%s %s\n", stage, v)

			wait := utils.RandInt(100, 2000)
			select {
			case <-ctx.Done():
			case <-time.After(time.Duration(wait) * time.Millisecond):
			}
		}
	}

	duration := float64(time.Since(now)) / float64(time.Second)
	fmt.Printf("%s release [optimized] target(s) in %.2f secs\n", "Finished", duration)
}
