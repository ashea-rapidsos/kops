/*
Copyright 2016 The Kubernetes Authors.

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

package components

import (
	"fmt"
	"github.com/blang/semver"
	"k8s.io/kops/pkg/apis/kops"
)

// OptionsContext is the context object for options builders
type OptionsContext struct {
	Cluster *kops.Cluster
}

// KubernetesVersion parses the semver version of kubernetes, from the cluster spec
func (c *OptionsContext) KubernetesVersion() (*semver.Version, error) {
	kubernetesVersion := c.Cluster.Spec.KubernetesVersion

	if kubernetesVersion == "" {
		return nil, fmt.Errorf("KubernetesVersion is required")
	}

	sv, err := kops.ParseKubernetesVersion(kubernetesVersion)
	if err != nil {
		return nil, fmt.Errorf("unable to determine kubernetes version from %q", kubernetesVersion)
	}

	return sv, nil
}
