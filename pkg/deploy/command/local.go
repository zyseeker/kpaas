// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type LocalShellCommand struct {
	cmd  string
	args []string
}

func NewLocalShellCommand(cmd string, args ...string) *LocalShellCommand {
	return &LocalShellCommand{
		cmd:  cmd,
		args: args,
	}
}

func (c *LocalShellCommand) Execute() (stderr, stdout []byte, err error) {

	cmd := exec.Command(c.cmd, c.args...)
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	err = cmd.Start()

	if stderr, err = ioutil.ReadAll(errReader); err != nil {
		return
	}

	if stdout, err = ioutil.ReadAll(outReader); err != nil {
		return
	}

	return
}

func (c *LocalShellCommand) GetCommand() string {

	return fmt.Sprint(c.cmd, strings.Join(c.args, " "))
}

func (c *LocalShellCommand) Exists() (isExist bool, err error) {

	cmd := exec.Command(getCommandExistShell(c.cmd))
	if err = cmd.Run(); err != nil {
		return
	}

	if cmd.ProcessState.ExitCode() == 0 {
		isExist = true
	}

	return
}
