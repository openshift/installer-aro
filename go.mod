module github.com/openshift/installer

go 1.13

require (
	github.com/Azure/azure-sdk-for-go v43.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.10.0
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.1
	github.com/Azure/go-autorest/autorest/to v0.3.1-0.20191028180845-3492b2aff503
	github.com/Azure/go-autorest/autorest/validation v0.2.1-0.20191028180845-3492b2aff503 // indirect
	github.com/Netflix/go-expect v0.0.0-20190729225929-0e00d9168667 // indirect
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/containers/image v3.0.2+incompatible
	github.com/coreos/go-systemd v0.0.0 // indirect
	github.com/coreos/ignition v0.35.0
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-logr/logr v0.2.1 // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.3.1
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/h2non/filetype v1.0.12
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hinshun/vt10x v0.0.0-20180809195222-d55458df857c // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/onsi/gomega v1.10.1 // indirect
	github.com/openshift/api v3.9.1-0.20191111211345-a27ff30ebf09+incompatible
	github.com/openshift/cluster-api v0.0.0-20191129101638-b09907ac6668
	github.com/openshift/machine-config-operator v4.2.0-alpha.0.0.20190917115525-033375cbe820+incompatible
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.10.1
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/ulikunitz/xz v0.5.6
	github.com/vincent-petithory/dataurl v0.0.0-20191104211930-d1553a71de50
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200428211428-0c9eba77bc32 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.9-0.20200217094205-6773bdf39b7f
	k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog/v2 v2.3.0 // indirect
	k8s.io/utils v0.0.0-20200729134348-d5654de09c73 // indirect
	sigs.k8s.io/cluster-api-provider-azure v0.0.0-00010101000000-000000000000
	sigs.k8s.io/controller-runtime v0.6.2 // indirect
	sigs.k8s.io/controller-tools v0.3.0
)

replace (
	bitbucket.org/ww/goautoneg => github.com/munnerz/goautoneg v0.0.0-20120707110453-a547fc61f48d // 404 on bitbucket.org/ww/goautoneg
	github.com/Azure/go-autorest => github.com/tombuildsstuff/go-autorest v14.0.1-0.20200416184303-d4e299a3c04a+incompatible
	github.com/Azure/go-autorest/autorest => github.com/tombuildsstuff/go-autorest/autorest v0.10.1-0.20200416184303-d4e299a3c04a
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/tombuildsstuff/go-autorest/autorest/azure/auth v0.4.3-0.20200416184303-d4e299a3c04a
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0 // Pin non-versioned import to v22.0.0
	github.com/go-log/log => github.com/go-log/log v0.1.1-0.20181211034820-a514cf01a3eb // Pinned by MCO
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0
	github.com/openshift/api => github.com/openshift/api v0.0.0-20200413201024-c6e8c9b6eb9a // Pin API
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20200116152001-92a2713fa240 // Pin client-go
	github.com/openshift/cluster-api-provider-baremetal => github.com/openshift/cluster-api-provider-baremetal v0.0.0-20200911144710-1cf0189fc640
	github.com/openshift/machine-api-operator => github.com/openshift/machine-api-operator v0.2.1-0.20201009151430-0af747bec740
	github.com/openshift/machine-config-operator => github.com/openshift/machine-config-operator v0.0.1-0.20201023000951-3b284da53fa0
	google.golang.org/api => google.golang.org/api v0.13.0 // Pin to version required by tf-provider-google
	k8s.io/api => k8s.io/api v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/apiserver => k8s.io/apiserver v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/client-go => k8s.io/client-go v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/code-generator => k8s.io/code-generator v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/component-base => k8s.io/component-base v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/cri-api => k8s.io/cri-api v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kubectl => k8s.io/kubectl v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kubelet => k8s.io/kubelet v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/kubernetes => k8s.io/kubernetes v1.17.1 // Replaced by MCO/CRI-O
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/metrics => k8s.io/metrics v0.17.1 // Replaced by MCO/CRI-O
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.1 // Replaced by MCO/CRI-O
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20200120114645-8a9592f1f87b // Pin OpenShift fork
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.5.3
	sigs.k8s.io/controller-tools => github.com/abhinavdahiya/controller-tools v0.3.1-0.20200430222905-6fdf2d5fc069 // Using fork for sigs.k8s.io/controller-tools#427
)
