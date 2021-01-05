// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"encoding/json"

	"github.com/openshift/installer/pkg/types/baremetal"
)

type config struct {
	LibvirtURI              string `json:"libvirt_uri"`
	BootstrapProvisioningIP string `json:"bootstrap_provisioning_ip"`
	BootstrapOSImage        string `json:"bootstrap_os_image"`
	ExternalBridge          string `json:"external_bridge"`
	ProvisioningBridge      string `json:"provisioning_bridge"`

	// Data required for control plane deployment - several maps per host, because of terraform's limitations
	Hosts         []map[string]interface{} `json:"hosts"`
	RootDevices   []map[string]interface{} `json:"root_devices"`
	Properties    []map[string]interface{} `json:"properties"`
	DriverInfos   []map[string]interface{} `json:"driver_infos"`
	InstanceInfos []map[string]interface{} `json:"instance_infos"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, bootstrapProvisioningIP, bootstrapOSImage, externalBridge, provisioningBridge string, platformHosts []*baremetal.Host, image string) ([]byte, error) {
	return json.MarshalIndent("", "", "  ")
}
