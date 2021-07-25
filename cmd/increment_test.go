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
	"io"
	"strings"
	"testing"
)

func TestNewIncrementCmd(t *testing.T) {
	for _, scenario := range []struct {
		when       string
		givenArgs  []string
		givenStdIn io.Reader
		expect     func(t *testing.T, err error, stdOut string, stdErr string)
	}{
		{
			when:      "bumping patch version from arg",
			givenArgs: []string{"--patch", "0.0.1"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "bumping minor version from arg",
			givenArgs: []string{"--minor", "0.0.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.1.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "bumping major version from arg",
			givenArgs: []string{"--major", "0.2.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "1.0.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "bumping patch version with v from arg",
			givenArgs: []string{"--patch", "v0.0.1"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "bumping minor version with v from arg",
			givenArgs: []string{"--minor", "v0.0.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v0.1.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "bumping major version with v from arg",
			givenArgs: []string{"--major", "v0.2.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v1.0.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "version is invalid from args",
			givenArgs: []string{"--major", "1.0.0.0"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.Error(t, err)
				require.True(t, strings.HasPrefix(stdOut, "Usage"), stdOut)
				require.Equal(t, "Error: strconv.ParseInt: parsing \"0.0\": invalid syntax\n", stdErr)
			},
		},
		{
			when:       "bumping patch version from stdin",
			givenArgs:  []string{"--patch"},
			givenStdIn: strings.NewReader("0.0.1"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "bumping minor version from stdin",
			givenArgs:  []string{"--minor"},
			givenStdIn: strings.NewReader("0.0.5"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.1.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "bumping major version from stdin",
			givenArgs:  []string{"--major"},
			givenStdIn: strings.NewReader("0.2.5"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "1.0.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "bumping patch version with v from stdin",
			givenArgs:  []string{"--patch"},
			givenStdIn: strings.NewReader("v0.0.1"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "bumping minor version with v from stdin",
			givenArgs:  []string{"--minor"},
			givenStdIn: strings.NewReader("v0.0.5"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v0.1.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "bumping major version with v from stdin",
			givenArgs:  []string{"--major"},
			givenStdIn: strings.NewReader("v0.2.5"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "v1.0.0", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "version is invalid from stdin",
			givenArgs:  []string{"--major"},
			givenStdIn: strings.NewReader("1.0.0.0"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.Error(t, err)
				require.True(t, strings.HasPrefix(stdOut, "Usage"), stdOut)
				require.Equal(t, "Error: strconv.ParseInt: parsing \"0.0\": invalid syntax\n", stdErr)
			},
		},
		{
			when:      "bumping patch version from stdin with linebreak",
			givenArgs: []string{"--patch"},
			givenStdIn: strings.NewReader("0.0.1\n"),
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.0.2", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:      "no flag was given",
			givenArgs: []string{"0.2.5"},
			expect: func(t *testing.T, err error, stdOut string, stdErr string) {
				require.NoError(t, err)
				require.Equal(t, "0.2.5", stdOut)
				require.Empty(t, stdErr)
			},
		},
		{
			when:       "just stdin was given",
			givenArgs:  []string{},
			givenStdIn: strings.NewReader("0.2.5"),
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

			if scenario.givenStdIn != nil {
				c.SetIn(scenario.givenStdIn)
			}
			c.SetArgs(scenario.givenArgs)
			c.SetOut(stdOut)
			c.SetErr(stdErr)
			scenario.expect(t, c.Execute(), stdOut.String(), stdErr.String())
		})
	}
}
