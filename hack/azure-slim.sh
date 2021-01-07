#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
set -x

# These packages are unused by ARO-RP and contain all of the direct usages of
# terraform
rm -rf cmd/openshift-install/ \
       pkg/destroy/ \
       pkg/terraform/

# tests mostly test AWS functionality
find pkg/ -name '*_test.go' -delete

# We do a mod tidy + vendor here since some carry patches remove required deps
# which cause problems building/updating to 4.6 pins
go mod tidy
go mod vendor

for x in aws azure openstack gcp libvirt ovirt; do
	go mod edit -replace sigs.k8s.io/cluster-api-provider-$x=github.com/openshift/cluster-api-provider-$x@$(curl https://proxy.golang.org/github.com/openshift/cluster-api-provider-$x/@v/release-4.6.info | jq -r ."Version")
done

for x in vendor/github.com/openshift/*; do
	go mod edit -replace ${x##vendor/}=${x##vendor/}@$(curl https://proxy.golang.org/${x##vendor/}/@v/release-4.6.info | jq -r ."Version")
done
go mod edit -replace github.com/openshift/client-go=github.com/openshift/client-go@$(curl https://proxy.golang.org/github.com/openshift/client-go/@v/release-4.6.info | jq -r ."Version")

for x in vendor/k8s.io/*; do
	if [[ "$x" = "vendor/k8s.io/utils" || "$x" = "vendor/k8s.io/klog" ]]; then
		continue
	elif [[ "$x" = "vendor/k8s.io/kube-openapi" ]]; then
		go mod edit -replace ${x##vendor/}=${x##vendor/}@$(curl https://proxy.golang.org/k8s.io/kube-openapi/@v/release-1.19.info | jq -r ."Version")
	else
		go mod edit -replace ${x##vendor/}=${x##vendor/}@v0.19.4
	fi
done

go mod tidy
go mod vendor

# Ensure the latest versions of the openshift dependencies work, and that nothing we need is gone
goimports -w pkg/
go vet ./...
