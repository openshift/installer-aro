package api

type AzureConfig struct {

	// SSH to system nodes allowed IP ranges
	SSHSourceAddressPrefixes []string `json:"sshSourceAddressPrefixes,omitempty"`
}
