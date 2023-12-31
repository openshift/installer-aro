// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RouteTables_Route_Spec. Use v1api20201101.RouteTables_Route_Spec instead
type RouteTables_Route_Spec_ARM struct {
	Name       string                     `json:"name,omitempty"`
	Properties *RoutePropertiesFormat_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RouteTables_Route_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (route RouteTables_Route_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (route *RouteTables_Route_Spec_ARM) GetName() string {
	return route.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables/routes"
func (route *RouteTables_Route_Spec_ARM) GetType() string {
	return "Microsoft.Network/routeTables/routes"
}

// Deprecated version of RoutePropertiesFormat. Use v1api20201101.RoutePropertiesFormat instead
type RoutePropertiesFormat_ARM struct {
	AddressPrefix    *string           `json:"addressPrefix,omitempty"`
	HasBgpOverride   *bool             `json:"hasBgpOverride,omitempty"`
	NextHopIpAddress *string           `json:"nextHopIpAddress,omitempty"`
	NextHopType      *RouteNextHopType `json:"nextHopType,omitempty"`
}
