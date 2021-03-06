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

package check

import (
	"bytes"

	"github.com/kpaas-io/kpaas/pkg/deploy/command"
	"github.com/kpaas-io/kpaas/pkg/deploy/machine"
	"github.com/kpaas-io/kpaas/pkg/deploy/operation"
	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
)

type CheckDockerOperation struct {
	shellCmd *command.ShellCommand
}

func (ckops *CheckDockerOperation) RunCommands(config *pb.NodeCheckConfig, logChan chan<- *bytes.Buffer) (stdOut, stdErr []byte, err error) {

	itemBuffer := &bytes.Buffer{}

	m, err := machine.NewMachine(config.Node)
	if err != nil {
		return nil, nil, err
	}

	defer m.Close()

	// construct command for check docker
	ckops.shellCmd = command.NewShellCommand(m, "docker", "version | grep -C1 'Client' | grep -w 'Version:' | awk '{print $2}'").
		WithDescription("检查机器 docker 版本是否满足最低要求").
		WithExecuteLogWriter(itemBuffer)

	// run commands
	stdOut, stdErr, err = ckops.shellCmd.Execute()

	// write buffer to channel
	logChan <- itemBuffer

	return
}

// check docker version if version larger or equal than standard version
func CheckDockerVersion(dockerVersion string, standardVersion string, comparedSymbol string) error {
	err := operation.CheckVersion(dockerVersion, standardVersion, comparedSymbol)
	if err != nil {
		return err
	}
	return nil
}
