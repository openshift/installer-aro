// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceIPSet,
			TypeName: "aws_wafv2_ip_set",
		},
		{
			Factory:  DataSourceRegexPatternSet,
			TypeName: "aws_wafv2_regex_pattern_set",
		},
		{
			Factory:  DataSourceRuleGroup,
			TypeName: "aws_wafv2_rule_group",
		},
		{
			Factory:  DataSourceWebACL,
			TypeName: "aws_wafv2_web_acl",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceIPSet,
			TypeName: "aws_wafv2_ip_set",
			Name:     "IP Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRegexPatternSet,
			TypeName: "aws_wafv2_regex_pattern_set",
			Name:     "Regex Pattern Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRuleGroup,
			TypeName: "aws_wafv2_rule_group",
			Name:     "Rule Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceWebACL,
			TypeName: "aws_wafv2_web_acl",
			Name:     "Web ACL",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceWebACLAssociation,
			TypeName: "aws_wafv2_web_acl_association",
		},
		{
			Factory:  ResourceWebACLLoggingConfiguration,
			TypeName: "aws_wafv2_web_acl_logging_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.WAFV2
}

var ServicePackage = &servicePackage{}
