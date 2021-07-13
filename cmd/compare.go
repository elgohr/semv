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
	"strings"
)

func NewEqualsCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "compare",
		Short: "check a semver for equality",
		Long:  "-1 = lower, 0 = equal, 1 = higher",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			v1, err := semver.NewVersion(strings.TrimLeft(args[0], "v"))
			if err != nil {
				return err
			}
			v2, err := semver.NewVersion(strings.TrimLeft(args[1], "v"))
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(cmd.OutOrStdout(), v1.Compare(*v2))
			return err
		},
	}
	return c
}
