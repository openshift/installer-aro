//go:build aro

package targets

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/cluster"
	"github.com/openshift/installer/pkg/asset/ignition/machine"
	"github.com/openshift/installer/pkg/asset/kubeconfig"
	"github.com/openshift/installer/pkg/asset/password"
	"github.com/openshift/installer/pkg/asset/tls"
)

var (
	// Cluster are the cluster targeted assets. Don't include Terraform variables.
	Cluster = []asset.WritableAsset{
		&cluster.Metadata{},
		&machine.MasterIgnitionCustomizations{},
		&machine.WorkerIgnitionCustomizations{},
		&kubeconfig.AdminClient{},
		&password.KubeadminPassword{},
		&tls.JournalCertKey{},
	}
)
