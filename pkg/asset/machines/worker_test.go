package machines

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/ignition/machine"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/rhcos"
	"github.com/openshift/installer/pkg/asset/templates/content/bootkube"
	"github.com/openshift/installer/pkg/types"
	awstypes "github.com/openshift/installer/pkg/types/aws"
)

var aroDNSWorkerMachineConfig = `apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-aro-dns
spec:
  config:
    ignition:
      config: {}
      security:
        tls: {}
      timeouts: {}
      version: 2.2.0
    networkd: {}
    passwd: {}
    storage:
      files:
      - contents:
          source: data:text/plain;charset=utf-8;base64,CnJlc29sdi1maWxlPS9ldGMvcmVzb2x2LmNvbmYuZG5zbWFzcQpzdHJpY3Qtb3JkZXIKYWRkcmVzcz0vYXBpLnRlc3QtY2x1c3Rlci50ZXN0LWRvbWFpbi8KYWRkcmVzcz0vYXBpLWludC50ZXN0LWNsdXN0ZXIudGVzdC1kb21haW4vCmFkZHJlc3M9Ly5hcHBzLnRlc3QtY2x1c3Rlci50ZXN0LWRvbWFpbi8KdXNlcj1kbnNtYXNxCmdyb3VwPWRuc21hc3EKbm8taG9zdHMKY2FjaGUtc2l6ZT0wCg==
          verification: {}
        filesystem: root
        mode: 420
        overwrite: true
        path: /etc/dnsmasq.conf
        user:
          name: root
      - contents:
          source: data:text/plain;charset=utf-8;base64,CiMhL2Jpbi9iYXNoCnNldCAtZXVvIHBpcGVmYWlsCgojIFRoaXMgYmFzaCBzY3JpcHQgaXMgYSBwYXJ0IG9mIHRoZSBBUk8gRG5zTWFzcSBjb25maWd1cmF0aW9uCiMgSXQncyBkZXBsb3llZCBhcyBwYXJ0IG9mIHRoZSA5OS1hcm8tZG5zLSogbWFjaGluZSBjb25maWcKIyBTZWUgaHR0cHM6Ly9naXRodWIuY29tL0F6dXJlL0FSTy1SUAoKIyBUaGlzIGZpbGUgY2FuIGJlIHJlcnVuIGFuZCB0aGUgZWZmZWN0IGlzIGlkZW1wb3RlbnQsIG91dHB1dCBtaWdodCBjaGFuZ2UgaWYgdGhlIERIQ1AgY29uZmlndXJhdGlvbiBjaGFuZ2VzCgpUTVBTRUxGUkVTT0xWPSQobWt0ZW1wKQpUTVBORVRSRVNPTFY9JChta3RlbXApCgplY2hvICIjIEdlbmVyYXRlZCBmb3IgZG5zbWFzcS5zZXJ2aWNlIC0gc2hvdWxkIHBvaW50IHRvIHNlbGYiID4gJFRNUFNFTEZSRVNPTFYKZWNobyAiIyBHZW5lcmF0ZWQgZm9yIGRuc21hc3Euc2VydmljZSAtIHNob3VsZCBjb250YWluIERIQ1AgY29uZmlndXJlZCBETlMiID4gJFRNUE5FVFJFU09MVgoKaWYgbm1jbGkgZGV2aWNlIHNob3cgYnItZXg7IHRoZW4KICAgIGVjaG8gIk9WTiBtb2RlIC0gYnItZXggZGV2aWNlIGV4aXN0cyIKICAgICNnZXR0aW5nIEROUyBzZWFyY2ggc3RyaW5ncwogICAgU0VBUkNIX1JBVz0kKG5tY2xpIC0tZ2V0IElQNC5ET01BSU4gZGV2aWNlIHNob3cgYnItZXgpCiAgICAjZ2V0dGluZyBETlMgc2VydmVycwogICAgTkFNRVNFUlZFUl9SQVc9JChubWNsaSAtLWdldCBJUDQuRE5TIGRldmljZSBzaG93IGJyLWV4IHwgdHIgLXMgIiB8ICIgIlxuIikKICAgIExPQ0FMX0lQU19SQVc9JChubWNsaSAtLWdldCBJUDQuQUREUkVTUyBkZXZpY2Ugc2hvdyBici1leCkKZWxzZQogICAgTkVUREVWPSQobm1jbGkgLS1nZXQgZGV2aWNlIGNvbm5lY3Rpb24gc2hvdyAtLWFjdGl2ZSB8IGhlYWQgLW4gMSkgI3RoZXJlIHNob3VsZCBiZSBvbmx5IG9uZSBhY3RpdmUgZGV2aWNlCiAgICBlY2hvICJPVlMgU0ROIG1vZGUgLSBici1leCBub3QgZm91bmQsIHVzaW5nIGRldmljZSAkTkVUREVWIgogICAgU0VBUkNIX1JBVz0kKG5tY2xpIC0tZ2V0IElQNC5ET01BSU4gZGV2aWNlIHNob3cgJE5FVERFVikKICAgIE5BTUVTRVJWRVJfUkFXPSQobm1jbGkgLS1nZXQgSVA0LkROUyBkZXZpY2Ugc2hvdyAkTkVUREVWIHwgdHIgLXMgIiB8ICIgIlxuIikKICAgIExPQ0FMX0lQU19SQVc9JChubWNsaSAtLWdldCBJUDQuQUREUkVTUyBkZXZpY2Ugc2hvdyAkTkVUREVWKQpmaQoKI3NlYXJjaCBsaW5lCmVjaG8gInNlYXJjaCAkU0VBUkNIX1JBVyIgfCB0ciAnXG4nICcgJyA+PiAkVE1QTkVUUkVTT0xWCmVjaG8gIiIgPj4gJFRNUE5FVFJFU09MVgplY2hvICJzZWFyY2ggJFNFQVJDSF9SQVciIHwgdHIgJ1xuJyAnICcgPj4gJFRNUFNFTEZSRVNPTFYKZWNobyAiIiA+PiAkVE1QU0VMRlJFU09MVgoKI25hbWVzZXJ2ZXJzIGFzIHNlcGFyYXRlIGxpbmVzCmVjaG8gIiROQU1FU0VSVkVSX1JBVyIgfCB3aGlsZSByZWFkIC1yIGxpbmUKZG8KICAgIGVjaG8gIm5hbWVzZXJ2ZXIgJGxpbmUiID4+ICRUTVBORVRSRVNPTFYKZG9uZQojIGRldmljZSBJUHMgYXJlIHJldHVybmVkIGluIGFkZHJlc3MvbWFzayBmb3JtYXQKZWNobyAiJExPQ0FMX0lQU19SQVciIHwgd2hpbGUgcmVhZCAtciBsaW5lCmRvCiAgICBlY2hvICJuYW1lc2VydmVyICRsaW5lIiB8IGN1dCAtZCcvJyAtZiAxID4+ICRUTVBTRUxGUkVTT0xWCmRvbmUKCiMgZG9uZSwgY29weWluZyBmaWxlcyB0byBkZXN0aW5hdGlvbiBsb2NhdGlvbnMgYW5kIGNsZWFuaW5nIHVwCi9iaW4vY3AgJFRNUE5FVFJFU09MViAvZXRjL3Jlc29sdi5jb25mLmRuc21hc3EKY2htb2QgMDc0NCAvZXRjL3Jlc29sdi5jb25mLmRuc21hc3EKL2Jpbi9jcCAkVE1QU0VMRlJFU09MViAvZXRjL3Jlc29sdi5jb25mCi91c3Ivc2Jpbi9yZXN0b3JlY29uIC9ldGMvcmVzb2x2LmNvbmYKL2Jpbi9ybSAkVE1QTkVUUkVTT0xWCi9iaW4vcm0gJFRNUFNFTEZSRVNPTFYK
          verification: {}
        filesystem: root
        mode: 484
        overwrite: true
        path: /usr/local/bin/aro-dnsmasq-pre.sh
        user:
          name: root
    systemd:
      units:
      - contents: |2

          [Unit]
          Description=DNS caching server.
          After=network-online.target
          Before=bootkube.service

          [Service]
          # ExecStartPre will create a copy of the customer current resolv.conf file and make it upstream DNS.
          # This file is a product of user DNS settings on the VNET. We will replace this file to point to
          # dnsmasq instance on the node. dnsmasq will inject certain dns records we need and forward rest of the queries to
          # resolv.conf.dnsmasq upstream customer dns.
          ExecStartPre=/bin/bash /usr/local/bin/aro-dnsmasq-pre.sh
          ExecStart=/usr/sbin/dnsmasq -k
          ExecStopPost=/bin/bash -c '/bin/mv /etc/resolv.conf.dnsmasq /etc/resolv.conf; /usr/sbin/restorecon /etc/resolv.conf'
          Restart=always

          [Install]
          WantedBy=multi-user.target
        enabled: true
        name: dnsmasq.service
  extensions: null
  fips: false
  kernelArguments: null
  kernelType: ""
  osImageURL: ""
`

func TestWorkerGenerate(t *testing.T) {
	cases := []struct {
		name                  string
		key                   string
		hyperthreading        types.HyperthreadingMode
		expectedMachineConfig []string
	}{
		{
			name:                  "no key hyperthreading enabled",
			hyperthreading:        types.HyperthreadingEnabled,
			expectedMachineConfig: []string{aroDNSWorkerMachineConfig},
		},
		{
			name:           "key present hyperthreading enabled",
			key:            "ssh-rsa: dummy-key",
			hyperthreading: types.HyperthreadingEnabled,
			expectedMachineConfig: []string{`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-ssh
spec:
  config:
    ignition:
      version: 3.2.0
    passwd:
      users:
      - name: core
        sshAuthorizedKeys:
        - 'ssh-rsa: dummy-key'
  extensions: null
  fips: false
  kernelArguments: null
  kernelType: ""
  osImageURL: ""
`, aroDNSWorkerMachineConfig},
		},
		{
			name:           "no key hyperthreading disabled",
			hyperthreading: types.HyperthreadingDisabled,
			expectedMachineConfig: []string{`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-disable-hyperthreading
spec:
  config:
    ignition:
      version: 3.2.0
  extensions: null
  fips: false
  kernelArguments:
  - nosmt
  kernelType: ""
  osImageURL: ""
`, aroDNSWorkerMachineConfig},
		},
		{
			name:           "key present hyperthreading disabled",
			key:            "ssh-rsa: dummy-key",
			hyperthreading: types.HyperthreadingDisabled,
			expectedMachineConfig: []string{`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-disable-hyperthreading
spec:
  config:
    ignition:
      version: 3.2.0
  extensions: null
  fips: false
  kernelArguments:
  - nosmt
  kernelType: ""
  osImageURL: ""
`, `apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-ssh
spec:
  config:
    ignition:
      version: 3.2.0
    passwd:
      users:
      - name: core
        sshAuthorizedKeys:
        - 'ssh-rsa: dummy-key'
  extensions: null
  fips: false
  kernelArguments: null
  kernelType: ""
  osImageURL: ""
`, aroDNSWorkerMachineConfig},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			parents := asset.Parents{}
			parents.Add(
				&installconfig.ClusterID{
					UUID:    "test-uuid",
					InfraID: "test-infra-id",
				},
				installconfig.MakeAsset(
					&types.InstallConfig{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test-cluster",
						},
						SSHKey:     tc.key,
						BaseDomain: "test-domain",
						Platform: types.Platform{
							AWS: &awstypes.Platform{
								Region: "us-east-1",
							},
						},
						Compute: []types.MachinePool{
							{
								Replicas:       pointer.Int64Ptr(1),
								Hyperthreading: tc.hyperthreading,
								Platform: types.MachinePoolPlatform{
									AWS: &awstypes.MachinePool{
										Zones:        []string{"us-east-1a"},
										InstanceType: "m5.large",
									},
								},
							},
						},
					}),
				(*rhcos.Image)(pointer.StringPtr("test-image")),
				(*rhcos.Release)(pointer.StringPtr("412.86.202208101040-0")),
				&machine.Worker{
					File: &asset.File{
						Filename: "worker-ignition",
						Data:     []byte("test-ignition"),
					},
				},
				&bootkube.ARODNSConfig{},
			)
			worker := &Worker{}
			if err := worker.Generate(parents); err != nil {
				t.Fatalf("failed to generate worker machines: %v", err)
			}
			expectedLen := len(tc.expectedMachineConfig)
			if assert.Equal(t, expectedLen, len(worker.MachineConfigFiles)) {
				for i := 0; i < expectedLen; i++ {
					assert.Equal(t, tc.expectedMachineConfig[i], string(worker.MachineConfigFiles[i].Data), "unexepcted machine config contents")
				}
			} else {
				assert.Equal(t, 0, len(worker.MachineConfigFiles), "expected no machine config files")
			}
		})
	}
}

func TestComputeIsNotModified(t *testing.T) {
	parents := asset.Parents{}
	installConfig := installconfig.MakeAsset(
		&types.InstallConfig{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-cluster",
			},
			SSHKey:     "ssh-rsa: dummy-key",
			BaseDomain: "test-domain",
			Platform: types.Platform{
				AWS: &awstypes.Platform{
					Region: "us-east-1",
					DefaultMachinePlatform: &awstypes.MachinePool{
						InstanceType: "TEST_INSTANCE_TYPE",
					},
				},
			},
			Compute: []types.MachinePool{
				{
					Replicas:       pointer.Int64Ptr(1),
					Hyperthreading: types.HyperthreadingDisabled,
					Platform: types.MachinePoolPlatform{
						AWS: &awstypes.MachinePool{
							Zones:        []string{"us-east-1a"},
							InstanceType: "",
						},
					},
				},
			},
		})

	parents.Add(
		&installconfig.ClusterID{
			UUID:    "test-uuid",
			InfraID: "test-infra-id",
		},
		installConfig,
		(*rhcos.Image)(pointer.StringPtr("test-image")),
		(*rhcos.Release)(pointer.StringPtr("412.86.202208101040-0")),
		&machine.Worker{
			File: &asset.File{
				Filename: "worker-ignition",
				Data:     []byte("test-ignition"),
			},
		},
		&bootkube.ARODNSConfig{},
	)
	worker := &Worker{}
	if err := worker.Generate(parents); err != nil {
		t.Fatalf("failed to generate master machines: %v", err)
	}

	if installConfig.Config.Compute[0].Platform.AWS.Type != "" {
		t.Fatalf("compute in the install config has been modified")
	}
}

func TestDefaultAWSMachinePoolPlatform(t *testing.T) {
	type testCase struct {
		name                string
		poolName            string
		expectedMachinePool awstypes.MachinePool
		assert              func(tc *testCase)
	}
	cases := []testCase{
		{
			name:     "default EBS type for compute pool",
			poolName: types.MachinePoolComputeRoleName,
			expectedMachinePool: awstypes.MachinePool{
				EC2RootVolume: awstypes.EC2RootVolume{
					Type: awstypes.VolumeTypeGp3,
					Size: decimalRootVolumeSize,
				},
			},
			assert: func(tc *testCase) {
				mp := defaultAWSMachinePoolPlatform(tc.poolName)
				want := tc.expectedMachinePool.EC2RootVolume.Type
				got := mp.EC2RootVolume.Type
				assert.Equal(t, want, got, "unexepcted EBS type")
			},
		},
		{
			name:     "default EBS type for edge pool",
			poolName: types.MachinePoolEdgeRoleName,
			expectedMachinePool: awstypes.MachinePool{
				EC2RootVolume: awstypes.EC2RootVolume{
					Type: awstypes.VolumeTypeGp2,
					Size: decimalRootVolumeSize,
				},
			},
			assert: func(tc *testCase) {
				mp := defaultAWSMachinePoolPlatform(tc.poolName)
				want := tc.expectedMachinePool.EC2RootVolume.Type
				got := mp.EC2RootVolume.Type
				assert.Equal(t, want, got, "unexepcted EBS type")
			},
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			tc.assert(&tc)
		})
	}
}
