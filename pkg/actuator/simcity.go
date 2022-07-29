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

func NewSimCity() Interface {
	return &simCity{}
}

type simCity struct{}

func (a *simCity) Name() string {
	return "simCity"
}

func (a *simCity) Execute(ctx context.Context, manifest map[string][]string) {
	data := manifest["simcity"]
	if len(data) == 0 {
		return
	}

	spinners := []string{"/", "-", "\\", "|"}
	maxSpinnerLoops := 20

	currentSet := getCurrentSet(data, -1, 0)
	var simcity string
	for i := 0; i < 500; i++ {
		spinnerLoops := utils.RandInt(1, maxSpinnerLoops)

		preSimcity := simcity
		for v := range currentSet {
			// Don't choose the same message twice in a row
			if v == preSimcity {
				continue
			}
			simcity = v
			break
		}

		// Choose a status/resolution per "task"
		resolution := utils.RandInputStr([]string{"FAIL", "WARN", "ERROR", "SUCCESS", "DEBUG", "OK"})

		// Prepare and color the messages
		unchecked := "[ ] "
		checked := "[o] "

		// Keep track of when the message is first printed
		first := true
		for si := 0; si < spinnerLoops; si++ {
			for _, spinner := range spinners {
				_ = utils.ClearPrint()
				// Output a message, with a checkbox in front and spinner behind
				msg := fmt.Sprintf("%s... %s", simcity, spinner)

				// on first print, text appears letter by letter
				if first {
					fmt.Print(unchecked)
					utils.VerbatimPrint(ctx, msg, 15)
					first = false
				} else {
					fmt.Print(unchecked)
					fmt.Print(msg)
				}

				// Wait a bit, then erase the line
				time.Sleep(50 * time.Millisecond)
				fmt.Print("\r")

				// Don't wait until finished, exit both loops if that is requested
				select {
				case <-ctx.Done():
					return
				default:
				}
			}
		}

		// Select the color
		switch resolution {
		case "FAIL", "ABORTED":
		case "SUCCESS", "OK":
		case "WARN":
		default:
		}

		// End of loop, the line has been removed, conclude the status
		utils.VerbatimPrint(ctx, checked, 10)
		fmt.Printf("%s... %s\n", simcity, resolution)
	}
}
