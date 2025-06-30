package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/vsphere"
)

func TestFeatureGates(t *testing.T) {
	cases := []struct {
		name          string
		installConfig *types.InstallConfig
		expected      string
	}{
		{
			name: "vSphere hosts is allowed with Feature Gates enabled",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = v1.TechPreviewNoUpgrade
				c.VSphere = validVSpherePlatform()
				c.VSphere.Hosts = []*vsphere.Host{{Role: "test"}}
				return c
			}(),
		},
		{
			name: "vSphere hosts is not allowed without Feature Gates",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.VSphere = validVSpherePlatform()
				c.VSphere.Hosts = []*vsphere.Host{{Role: "test"}}
				return c
			}(),
			expected: `^platform.vsphere.hosts: Forbidden: this field is protected by the VSphereStaticIPs feature gate which must be enabled through either the TechPreviewNoUpgrade or CustomNoUpgrade feature set$`,
		},
		{
			name: "vSphere hosts is allowed with custom Feature Gates",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = v1.CustomNoUpgrade
				c.FeatureGates = []string{"VSphereStaticIPs=true"}
				c.VSphere = validVSpherePlatform()
				c.VSphere.Hosts = []*vsphere.Host{{Role: "test"}}
				return c
			}(),
		},
		{
			name: "vSphere hosts is not allowed with custom Feature Gate disabled",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = v1.CustomNoUpgrade
				c.FeatureGates = []string{"VSphereStaticIPs=false"}
				c.VSphere = validVSpherePlatform()
				c.VSphere.Hosts = []*vsphere.Host{{Role: "test"}}
				return c
			}(),
			expected: `^platform.vsphere.hosts: Forbidden: this field is protected by the VSphereStaticIPs feature gate which must be enabled through either the TechPreviewNoUpgrade or CustomNoUpgrade feature set$`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateFeatureSet(tc.installConfig).ToAggregate()
			if tc.expected == "" {
				assert.NoError(t, err)
			} else {
				assert.Regexp(t, tc.expected, err)
			}
		})
	}
}
