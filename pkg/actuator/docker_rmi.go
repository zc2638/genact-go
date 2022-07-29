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

func NewDockerRmi() Interface {
	return &dockerRmi{}
}

type dockerRmi struct{}

func (a *dockerRmi) Name() string {
	return "docker_rmi"
}

func (a *dockerRmi) Execute(ctx context.Context, manifest map[string][]string) {
	data := manifest["docker"]
	if len(data) == 0 {
		return
	}

	currentData := getCurrentData(data, 20, 100)
	for _, name := range currentData {
		version := utils.RandVersion(false)
		fmt.Printf("Untagged: %s:%s\n", name, version)
		fmt.Printf("Untagged: %s:%s@sha256:%s\n", name, version, utils.RandHashStr(64))

		hashNum := utils.RandInt(3, 15)
		for i := 0; i < hashNum; i++ {
			fmt.Printf("Deleted: sha256:%s\n", utils.RandHashStr(64))
		}

		waitMillSec := utils.RandInt64(500, 5000)
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(waitMillSec) * time.Millisecond):
		}
	}
}
