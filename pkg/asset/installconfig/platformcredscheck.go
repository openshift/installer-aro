package installconfig

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/types/azure"
)

// PlatformCredsCheck is an asset that checks the platform credentials, asks for them or errors out if invalid
// the cluster.
type PlatformCredsCheck struct {
}

var _ asset.Asset = (*PlatformCredsCheck)(nil)

// Dependencies returns the dependencies for PlatformCredsCheck
func (a *PlatformCredsCheck) Dependencies() []asset.Asset {
	return []asset.Asset{
		&InstallConfig{},
	}
}

// Generate queries for input from the user.
func (a *PlatformCredsCheck) Generate(dependencies asset.Parents) error {
	ic := &InstallConfig{}
	dependencies.Get(ic)

	var err error
	platform := ic.Config.Platform.Name()
	switch platform {
	case azure.Name:
		_, err = ic.Azure.Session()
		if err != nil {
			return errors.Wrap(err, "creating Azure session")
		}
	default:
		err = fmt.Errorf("unknown platform type %q", platform)
	}

	return err
}

// Name returns the human-friendly name of the asset.
func (a *PlatformCredsCheck) Name() string {
	return "Platform Credentials Check"
}
