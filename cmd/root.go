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
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	cmd := NewRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), err.Error())
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "semv",
		Short:   "To be used for working with semvers",
		Version: "0.0.1",
	}
	c.AddCommand(NewIncrementCmd())
	c.AddCommand(NewEqualsCmd())
	c.AddCommand(NewLicenseCmd())
	return c
}
