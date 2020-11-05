#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
set -x

for platform in aws baremetal gcp libvirt openstack ovirt vsphere ; do \
	rm -rf pkg/asset/cluster/$platform/ \
			pkg/asset/installconfig/$platform/ \
			pkg/asset/machines/$platform/ \
			pkg/asset/manifests/$platform/ \
			pkg/asset/ignition/bootstrap/$platform/ \
			pkg/asset/cluster/$platform/ \
			pkg/asset/types/$platform/ \
			pkg/asset/quota/$platform/ \
			pkg/destroy/$platform/ \
			pkg/tfvars/$platform/ \
			pkg/quota/$platform/ \
			pkg/types/$platform/ \
			pkg/types/installconfig_$platform.go \
			; \
done
rm -rf cmd/openshift-install/ \
		pkg/destroy/ \
		pkg/terraform/

find pkg/ -name '*_test.go' -delete

declare -a cleanups=(
	"s/^((?<tab>\t*)(case|if platform ==) ((types)?(aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere|none)(types)?\.Name(, )?)*(?<st> {)?(:)?\n(^\g{tab}\t+.*\n|\n)+(?:(?(<st>)(\g{tab}})|)))//gm"
	's/^(?:(?<tab>\t*)case (?:(?:types)?(?:aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere|none)(?:types)?\.Name(?:, )?)*(?<az>azure(?:types)?\.Name)(?:(, )?(?:types)?(?:aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere|none)(?:types)?\.Name(?:, )?)*):/$1case $2:/gm'
	's/^(?:(?<tab>\t*)(?<case>case|if) [a-zA-Z]+(?:\.Config)?(?:\.Platform)?\.(?:AWS|Ovirt|BareMetal|Libvirt|OpenStack|GCP|VSphere) != nil \|\| (?<az>.*)(?<st> {|:))/$1$2 $3$4/gm'
	's/^(?:(?<tab>\t*)case (?:(?:types)?(?:aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere|none)(?:types)?\.Name(?:, )?)*)://gm'
	"s/^\t*switch \S+\.Platform\.Name\(\) \{\s*\}//gm"
	"s/^\t+(aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere|none)(provider|api)+\..*\n//gm"
	"s/^.*\*(aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere|none)\.Metadata.*\n//gm"
	"s/^((?<tab>\t*)(case|if) [a-zA-Z]+(\.Config)?(\.Platform)?\.(AWS|Ovirt|BareMetal|Libvirt|OpenStack|GCP|VSphere) != nil( {|:)(^\g{tab}\t+.*(\n)?|\n)*(\g{tab}})?)//gm"
	"s/^(func (default(AWS|Ovirt|BareMetal|Libvirt|OpenStack|GCP|VSphere)|awsDefault|ConvertBaremetal\(|validateGCPMachinePool\().*{(^\t+.*(\n)?|\n)*(}(\n)+))//gm"
	"s/^\t*(?:types)?(?:aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere)(?:types)?\.Name(,|:\s+.*)//gm"
	's/^(\t+[a-zA-Z]+(\.Config)?(\.Platform)?\.(AWS|Ovirt|BareMetal|Libvirt|OpenStack|GCP|VSphere) =.*\n)//gm'
	"s/^\t(AWS|Ovirt|BareMetal|Baremetal|Libvirt|OpenStack|GCP|VSphere)\s+\*(?:aws|libvirt|openstack|gcp|baremetal|ovirt|vsphere)\.(Platform|TemplateData|MachinePool).*$//gm"
)

for val in "${cleanups[@]}"; do
  find pkg/ -name "*.go" -exec perl -i -0777 -pe "$val" {} \;
done

declare -a modcleanups=(
	"s/^(\t*(github\.com\/(terraform-providers|metal3-io|vmware|hashicorp)\/.*|sigs\.k8s\.io\/cluster-api-provider-(aws|openstack)) =>.*\n)//gm"
)

for val in "${modcleanups[@]}"; do
  perl -i -0777 -pe "$val" go.mod
done

goimports -w pkg/

go mod tidy
go mod vendor

go vet ./...

go mod tidy
go mod vendor
