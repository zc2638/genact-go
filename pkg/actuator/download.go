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
	"os"
	"time"

	"github.com/schollz/progressbar/v3"

	"github.com/zc2638/genact-go/pkg/utils"
)

func NewDownload() Interface {
	return &download{}
}

type download struct{}

func (a *download) Name() string {
	return "download"
}

func (a *download) Execute(ctx context.Context, manifest map[string][]string) {
	data := manifest["packages"]
	dataLen := len(data)
	if dataLen == 0 {
		return
	}

	extension := extensions[utils.RandInt(0, len(extensions)-1)]
	fileNum := utils.RandInt(3, 10)

	for i := 0; i < fileNum; i++ {
		// File size in bytes.
		fileBytes := utils.RandInt64(30_000_000, 300_000_000)
		name := data[utils.RandInt(0, dataLen-1)]
		fileName := name + "." + extension

		waitMills := utils.RandInt64(500, 5000)

		bar := progressbar.NewOptions64(
			fileBytes,
			progressbar.OptionSetDescription(fileName),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(10),
			progressbar.OptionThrottle(time.Millisecond*50),
			progressbar.OptionShowCount(),
			progressbar.OptionOnCompletion(func() { fmt.Fprint(os.Stderr, "\n") }),
			progressbar.OptionSpinnerType(14),
			progressbar.OptionFullWidth(),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerPadding: " ",
				SaucerHead:    ">",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)
		bar.RenderBlank()

		everyWaitBytes := fileBytes / waitMills * 10
		for bi := 0; bi < int(waitMills/10); bi++ {
			bar.Add64(everyWaitBytes)
			time.Sleep(time.Millisecond * 10)
		}
		bar.Finish()
	}
}
