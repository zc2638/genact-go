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

package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/zc2638/releaser"

	"github.com/99nil/gopkg/sets"
	"github.com/spf13/cobra"

	"github.com/zc2638/genact-go/pkg/actuator"
	"github.com/zc2638/genact-go/static"
)

type Option struct {
	ListModules bool
	Modules     []string
}

func NewRoot() *cobra.Command {
	opt := &Option{}
	cmd := &cobra.Command{
		Use:          "genact",
		Short:        "A nonsense activity generator",
		Version:      releaser.Version.String(),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := static.Data()
			if err != nil {
				return err
			}
			modules := map[actuator.Interface]struct{}{
				actuator.NewDownload():  {},
				actuator.NewWeblog():    {},
				actuator.NewSimCity():   {},
				actuator.NewCargo():     {},
				actuator.NewDockerRmi(): {},
			}

			if opt.ListModules {
				for module := range modules {
					fmt.Println(module.Name())
				}
				return nil
			}
			if len(opt.Modules) > 0 {
				names := sets.NewString(opt.Modules...)
				current := make(map[actuator.Interface]struct{}, names.Len())
				for module := range modules {
					if !names.Has(module.Name()) {
						continue
					}
					current[module] = struct{}{}
				}
				modules = current
			}
			if len(modules) == 0 {
				return errors.New("module not found")
			}

			ctx := context.Background()
			for {
				for module := range modules {
					module.Execute(ctx, data)
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&opt.ListModules, "list-modules", "l", opt.ListModules, "List available modules")
	cmd.Flags().StringArrayVarP(&opt.Modules, "modules", "m", opt.Modules, "Run only these modules")
	return cmd
}
