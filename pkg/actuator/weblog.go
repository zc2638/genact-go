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
	"net/http"
	"time"

	"github.com/zc2638/genact-go/pkg/utils"
)

var httpCodes = []int{
	http.StatusOK,
	http.StatusCreated,
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
}

func NewWeblog() Interface {
	return &weblog{}
}

type weblog struct{}

func (a *weblog) Name() string {
	return "weblog"
}

func (a *weblog) Execute(ctx context.Context, manifest map[string][]string) {
	data := manifest["packages"]
	if len(data) == 0 {
		return
	}
	currentData := getCurrentData(data, -1, 0)

	numLines := utils.RandInt(50, 200)
	burstMode := false
	countBurstLines := 0

	for i := 0; i < numLines; i++ {
		name := currentData[utils.RandInt(0, len(currentData)-1)]
		currentPath := name + extensions[utils.RandInt(0, len(extensions)-1)]
		dateTime := time.Now().Format(time.RFC1123Z)
		method := http.MethodGet
		httpCode := httpCodes[utils.RandInt(0, len(httpCodes)-1)]
		size := utils.RandInt(99, 5000000)
		referrer := "-"
		userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
		line := fmt.Sprintf("[%s] \"%s %s HTTP/1.0\" %d %d \"%s\" \"%s\" ",
			dateTime, method, currentPath, httpCode, size, referrer, userAgent)

		wait := utils.RandInt(10, 1000)
		burstLines := utils.RandInt(10, 50)
		if burstMode && countBurstLines < burstLines {
			wait = 30
		} else if countBurstLines == burstLines {
			burstMode = false
			countBurstLines = 0
		} else if !burstMode {
			boolNum := utils.RandInt(0, 4)
			burstMode = boolNum == 0
		}

		fmt.Println(line)
		if burstMode {
			countBurstLines += 1
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(wait) * time.Millisecond):
		}
	}
}
