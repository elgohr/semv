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

package cmd_test

import (
	"bytes"
	"github.com/elgohr/semv/cmd"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewIncrementCmd(t *testing.T) {
	for _, scenario := range []struct {
		when   string
		given  []string
		expect func(t *testing.T, err error, stdOut string, stdErr string)
	}{
		{
			when:  "bumping patch version",
			given: []string{"--patch", "0.0.1"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:  "bumping minor version",
			given: []string{"--minor", "0.0.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.1.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:  "bumping major version",
			given: []string{"--major", "0.2.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "1.0.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:  "bumping patch version with v",
			given: []string{"--patch", "v0.0.1"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:  "bumping minor version with v",
			given: []string{"--minor", "v0.0.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v0.1.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:  "bumping major version with v",
			given: []string{"--major", "v0.2.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v1.0.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:  "version is invalid",
			given: []string{"--major", "1.0.0.0"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.Error(t, err)
				require.True(t, strings.HasPrefix(stdOut, "Usage"), stdOut)
				require.Equal(t, "Error: strconv.ParseInt: parsing \"0.0\": invalid syntax\n", stdErr)
			},
		},
		{
			when:  "no flag was given",
			given: []string{"0.2.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.2.5", stdOut)
				require.Empty(t, stdErr)
			},
		},
	} {
		t.Run(scenario.when, func(t *testing.T) {
			stdOut := bytes.NewBufferString("")
			stdErr := bytes.NewBufferString("")
			c := cmd.NewIncrementCmd()
			c.SetArgs(scenario.given)
			c.SetOut(stdOut)
			c.SetErr(stdErr)
			scenario.expect(t, c.Execute(), stdOut.String(), stdErr.String())
		})
	}
}
