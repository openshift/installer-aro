//go:build tools
// +build tools

// Official workaround to track tool dependencies with go modules:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/daixiang0/gci" // dependency of hack/go-fmt.sh
	// Needed for mock generation
	_ "go.uber.org/mock/mockgen/model"
	// dependency of generating CRD for install-config
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
