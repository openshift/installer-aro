package quota

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/machines"
	"github.com/openshift/installer/pkg/diagnostics"
	"github.com/openshift/installer/pkg/quota"
)

// PlatformQuotaCheck is an asset that validates the install-config platform for
// any resource requirements based on the quotas available.
type PlatformQuotaCheck struct {
}

var _ asset.Asset = (*PlatformQuotaCheck)(nil)

// Dependencies returns the dependencies for PlatformQuotaCheck
func (a *PlatformQuotaCheck) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
		&machines.Master{},
		&machines.Worker{},
	}
}

// Generate queries for input from the user.
func (a *PlatformQuotaCheck) Generate(dependencies asset.Parents) error {
	ic := &installconfig.InstallConfig{}
	mastersAsset := &machines.Master{}
	workersAsset := &machines.Worker{}
	dependencies.Get(ic, mastersAsset, workersAsset)

	return nil
}

// Name returns the human-friendly name of the asset.
func (a *PlatformQuotaCheck) Name() string {
	return "Platform Quota Check"
}

// summarizeFailingReport summarizes a report when there are failing constraints.
func summarizeFailingReport(reports []quota.ConstraintReport) error {
	var notavailable []string
	var unknown []string
	for _, report := range reports {
		switch report.Result {
		case quota.NotAvailable:
			notavailable = append(notavailable, fmt.Sprintf("%s is not available in %s because %s", report.For.Name, report.For.Region, report.Message))
		case quota.Unknown:
			unknown = append(unknown, report.For.Name)
		default:
			continue
		}
	}

	if len(notavailable) == 0 && len(unknown) > 0 {
		// all quotas are missing information so warn and skip
		logrus.Warnf("Failed to find information on quotas %s", strings.Join(unknown, ", "))
		return nil
	}

	msg := strings.Join(notavailable, ", ")
	if len(unknown) > 0 {
		msg = fmt.Sprintf("%s, and could not find information on %s", msg, strings.Join(unknown, ", "))
	}
	return &diagnostics.Err{Reason: "MissingQuota", Message: msg}
}

// summarizeReport summarizes a report when there are availble.
func summarizeReport(reports []quota.ConstraintReport) {
	var low []string
	for _, report := range reports {
		switch report.Result {
		case quota.AvailableButLow:
			low = append(low, fmt.Sprintf("%s (%s)", report.For.Name, report.For.Region))
		default:
			continue
		}
	}
	if len(low) > 0 {
		logrus.Warnf("Following quotas %s are available but will be completely used pretty soon.", strings.Join(low, ", "))
	}
}
