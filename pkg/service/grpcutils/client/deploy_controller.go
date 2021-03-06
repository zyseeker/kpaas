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

package client

import (
	"github.com/kpaas-io/kpaas/pkg/deploy/protos"
	"github.com/kpaas-io/kpaas/pkg/service/grpcutils/connection"
)

var (
	deployControllerClient protos.DeployContollerClient
)

func GetDeployController() protos.DeployContollerClient {

	if deployControllerClient != nil {
		return deployControllerClient
	}

	conn := connection.GetDeployControllerConnection()

	deployControllerClient = protos.NewDeployContollerClient(conn)
	return deployControllerClient
}

func SetDeployController(client protos.DeployContollerClient) {

	deployControllerClient = client
}
