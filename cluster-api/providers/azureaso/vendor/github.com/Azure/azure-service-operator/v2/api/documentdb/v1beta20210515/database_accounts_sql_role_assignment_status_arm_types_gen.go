// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

// Deprecated version of DatabaseAccounts_SqlRoleAssignment_STATUS. Use v1api20210515.DatabaseAccounts_SqlRoleAssignment_STATUS instead
type DatabaseAccounts_SqlRoleAssignment_STATUS_ARM struct {
	Id         *string                               `json:"id,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	Properties *SqlRoleAssignmentResource_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                               `json:"type,omitempty"`
}

// Deprecated version of SqlRoleAssignmentResource_STATUS. Use v1api20210515.SqlRoleAssignmentResource_STATUS instead
type SqlRoleAssignmentResource_STATUS_ARM struct {
	PrincipalId      *string `json:"principalId,omitempty"`
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`
	Scope            *string `json:"scope,omitempty"`
}
