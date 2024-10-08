// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformation

// Exports for use in tests only.
var (
	ResourceStack            = resourceStack
	ResourceStackSet         = resourceStackSet
	ResourceStackSetInstance = resourceStackSetInstance
	ResourceStackInstances   = resourceStackInstances
	ResourceType             = resourceType

	FindStackInstanceByFourPartKey          = findStackInstanceByFourPartKey
	FindStackInstanceSummariesByFourPartKey = findStackInstanceSummariesByFourPartKey
	FindStackSetByName                      = findStackSetByName
	FindTypeByARN                           = findTypeByARN
	FindStackInstancesByNameCallAs          = findStackInstancesByNameCallAs
	StackSetInstanceResourceIDPartCount     = stackSetInstanceResourceIDPartCount
	StackInstancesResourceIDPartCount       = stackInstancesResourceIDPartCount
	TypeVersionARNToTypeARNAndVersionID     = typeVersionARNToTypeARNAndVersionID
)
