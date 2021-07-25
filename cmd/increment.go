// Copyright © 2021 Lars Gohr <lars@gohr.digital>
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

package cmd

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func NewIncrementCmd() *cobra.Command {
	changePatch := false
	changeMinor := false
	changeMajor := false

	c := &cobra.Command{
		Use:   "increment",
		Short: "Increment a semver",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var inputVersion string
			if len(args) > 0 {
				inputVersion = args[0]
			} else {
				in, err := ioutil.ReadAll(cmd.InOrStdin())
				if err != nil {
					return err
				}
				inputVersion = strings.TrimRight(string(in), "\n")
			}

			hasPrefix := false
			if strings.HasPrefix(inputVersion, "v") {
				hasPrefix = true
				inputVersion = strings.TrimLeft(inputVersion, "v")
			}
			v, err := semver.NewVersion(inputVersion)
			if err != nil {
				return err
			}

			if changePatch {
				v.BumpPatch()
			} else if changeMinor {
				v.BumpMinor()
			} else if changeMajor {
				v.BumpMajor()
			}

			if hasPrefix {
				_, err = fmt.Fprint(cmd.OutOrStdout(), "v"+v.String())
			} else {
				_, err = fmt.Fprint(cmd.OutOrStdout(), v.String())
			}

			return err
		},
	}

	c.PersistentFlags().BoolVarP(&changePatch, "patch", "p", false, "change patch version")
	c.PersistentFlags().BoolVarP(&changeMinor, "minor", "l", false, "change minor version")
	c.PersistentFlags().BoolVarP(&changeMajor, "major", "m", false, "change major version")

	return c
}
