package machine

import (
	"fmt"
	"net/url"

	ignutil "github.com/coreos/ignition/v2/config/util"
	igntypes "github.com/coreos/ignition/v2/config/v3_1/types"
	"github.com/vincent-petithory/dataurl"

	"github.com/openshift/installer/pkg/types"
)

// pointerIgnitionConfig generates a config which references the remote config
// served by the machine config server.
func pointerIgnitionConfig(installConfig *types.InstallConfig, rootCA []byte, role string) *igntypes.Config {
	var ignitionHost string
	// Default platform independent ignitionHost
	ignitionHost = fmt.Sprintf("api-int.%s:22623", installConfig.ClusterDomain())
	// Update ignitionHost as necessary for platform

	return &igntypes.Config{
		Ignition: igntypes.Ignition{
			Version: igntypes.MaxVersion.String(),
			Config: igntypes.IgnitionConfig{
				Merge: []igntypes.Resource{{
					Source: ignutil.StrToPtr(func() *url.URL {
						return &url.URL{
							Scheme: "https",
							Host:   ignitionHost,
							Path:   fmt.Sprintf("/config/%s", role),
						}
					}().String()),
				}},
			},
			Security: igntypes.Security{
				TLS: igntypes.TLS{
					CertificateAuthorities: []igntypes.Resource{{
						Source: ignutil.StrToPtr(dataurl.EncodeBytes(rootCA)),
					}},
				},
			},
		},
	}
}
