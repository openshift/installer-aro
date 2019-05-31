package api

type AROConfig struct {

	// configuration of VMs
	ImageOffer     string `json:"imageOffer,omitempty"`
	ImagePublisher string `json:"imagePublisher,omitempty"`
	ImageSKU       string `json:"imageSku,omitempty"`
	ImageVersion   string `json:"imageVersion,omitempty"`

	// SSH to system nodes allowed IP ranges
	SSHSourceAddressPrefixes []string `json:"sshSourceAddressPrefixes,omitempty"`
}
