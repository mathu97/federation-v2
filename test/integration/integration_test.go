/*
Copyright 2018 The Federation v2 Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration

import (
	"testing"

	"github.com/kubernetes-sigs/federation-v2/test/integration/framework"
)

var FedFixture *framework.FederationFixture

// TestIntegration provides a common setup and teardown for all the integration tests.
func TestIntegration(t *testing.T) {
	tl := framework.NewIntegrationLogger(t)
	FedFixture = framework.SetUpFederationFixture(tl, 2)
	defer FedFixture.TearDown(tl)

	t.Run("Parallel-Integration-Test-Group", func(t *testing.T) {
		t.Run("TestCrud", TestCrud)
		t.Run("TestReplicaSchedulingPreference", TestReplicaSchedulingPreference)
	})
}
