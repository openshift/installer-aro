package api

import (
	"crypto/rsa"

	v1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

// Config holds the cluster config structure
type Config struct {
	// InstallerVersion defines installer version
	InstallerVersion string `json:"pluginVersion,omitempty"`

	// ComponentLogLevel specifies the log levels for the various openshift components
	ComponentLogLevel ComponentLogLevel `json:"componentLogLevel,omitempty"`

	// PlatformConfig defines platform specific Configuration
	PlatformConfig `json:"platformConfig"`

	SSHKey *rsa.PrivateKey `json:"sshKey,omitempty"`

	// Manifest defines all cluster manifests configurations
	Manifests ManifestConfig `json:"manifest,omitempty"`

	// Ignition defines all ignition configurations
	Ignition IgnitionConfig `json:"ignition,omitempty"`

	// Images defines all images configurations
	Images ImageConfig `json:"images,omitempty"`

	// kubeconfigs
	AdminKubeconfig  *v1.Config `json:"adminKubeconfig,omitempty"`
	MasterKubeconfig *v1.Config `json:"masterKubeconfig,omitempty"`
}

type PlatformConfig struct {
	Azure AzureConfig
	ARO   AROConfig
	//AWS       AWSConfig
	//OpenStack OpenStackConfig
}

// ComponentLogLevel represents the log levels for the various components of a
// cluster
type ComponentLogLevel struct {
	APIServer         *int `json:"apiServer,omitempty"`
	ControllerManager *int `json:"controllerManager,omitempty"`
	Node              *int `json:"node,omitempty"`
}

// ImageConfig contains all images for the pods
type ImageConfig struct {
	// Format of the pull spec that is going to be
	// used in the cluster.
	Format string `json:"format,omitempty"`

	APIServer string `json:"apiServer,omitempty"`

	// ImagePullSecret defines the secret used to pull from the private registries, used system-wide
	ImagePullSecret []byte `json:"imagePullSecret,omitempty"`
}

// IgnitionConfig contains all ignition configs
type IgnitionConfig struct {
	Bootstrap []byte `json:"bootstrap,omitempty"`
	Master    []byte `json:"master,omitempty"`
	Worker    []byte `json:"worker,omitempty"`
}

// ManifestConfig contains all manifests for the different components
type ManifestConfig struct {
	Ingress       Manifest
	NetworkCRD    Manifest
	NetworkConfig Manifest
}

type Priority int

const (
	Boot Priority = iota
	Install
	Startup
	PostStart
)

type Manifest struct {
	FileName string
	Priority Priority
	Name     string
	Data     []byte
}
