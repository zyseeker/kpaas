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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	desiredKernelVersion = "4.19.46"
)

// unit test of CheckKernelVersion
func TestCheckKernelVersion(t *testing.T) {
	testSample := []struct {
		comparedVersion string
		desiredVersion  string
		want            error
	}{
		{
			comparedVersion: "4.19.46",
			desiredVersion:  desiredKernelVersion,
			want:            nil,
		},
		{
			comparedVersion: "4.20",
			desiredVersion:  desiredKernelVersion,
			want:            nil,
		},
		{
			comparedVersion: "3.10.0-957.21.3.el7.x86_64",
			desiredVersion:  desiredKernelVersion,
			want:            fmt.Errorf("version too low, desired version: %v, actual version: 3.10.0-957.21.3.el7.x86_64", desiredKernelVersion),
		},
		{
			comparedVersion: "4.18.5-041805-generic",
			desiredVersion:  desiredKernelVersion,
			want:            fmt.Errorf("version too low, desired version: %v, actual version: 4.18.5-041805-generic", desiredKernelVersion),
		},
	}

	for _, eachValue := range testSample {
		assert.Equal(t, eachValue.want, CheckKernelVersion(eachValue.comparedVersion, eachValue.desiredVersion, ">"))
	}
}
