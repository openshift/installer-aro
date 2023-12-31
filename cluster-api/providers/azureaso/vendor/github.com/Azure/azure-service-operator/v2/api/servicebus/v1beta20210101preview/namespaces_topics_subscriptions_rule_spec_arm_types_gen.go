// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespaces_Topics_Subscriptions_Rule_Spec. Use v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec instead
type Namespaces_Topics_Subscriptions_Rule_Spec_ARM struct {
	Name       string              `json:"name,omitempty"`
	Properties *Ruleproperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_Topics_Subscriptions_Rule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (rule Namespaces_Topics_Subscriptions_Rule_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rule *Namespaces_Topics_Subscriptions_Rule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics/subscriptions/rules"
func (rule *Namespaces_Topics_Subscriptions_Rule_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics/subscriptions/rules"
}

// Deprecated version of Ruleproperties. Use v1api20210101preview.Ruleproperties instead
type Ruleproperties_ARM struct {
	Action            *Action_ARM            `json:"action,omitempty"`
	CorrelationFilter *CorrelationFilter_ARM `json:"correlationFilter,omitempty"`
	FilterType        *FilterType            `json:"filterType,omitempty"`
	SqlFilter         *SqlFilter_ARM         `json:"sqlFilter,omitempty"`
}

// Deprecated version of Action. Use v1api20210101preview.Action instead
type Action_ARM struct {
	CompatibilityLevel    *int    `json:"compatibilityLevel,omitempty"`
	RequiresPreprocessing *bool   `json:"requiresPreprocessing,omitempty"`
	SqlExpression         *string `json:"sqlExpression,omitempty"`
}

// Deprecated version of CorrelationFilter. Use v1api20210101preview.CorrelationFilter instead
type CorrelationFilter_ARM struct {
	ContentType           *string           `json:"contentType,omitempty"`
	CorrelationId         *string           `json:"correlationId,omitempty"`
	Label                 *string           `json:"label,omitempty"`
	MessageId             *string           `json:"messageId,omitempty"`
	Properties            map[string]string `json:"properties,omitempty"`
	ReplyTo               *string           `json:"replyTo,omitempty"`
	ReplyToSessionId      *string           `json:"replyToSessionId,omitempty"`
	RequiresPreprocessing *bool             `json:"requiresPreprocessing,omitempty"`
	SessionId             *string           `json:"sessionId,omitempty"`
	To                    *string           `json:"to,omitempty"`
}

// Deprecated version of SqlFilter. Use v1api20210101preview.SqlFilter instead
type SqlFilter_ARM struct {
	CompatibilityLevel    *int    `json:"compatibilityLevel,omitempty"`
	RequiresPreprocessing *bool   `json:"requiresPreprocessing,omitempty"`
	SqlExpression         *string `json:"sqlExpression,omitempty"`
}
