package model

import v1 "github.com/roohitkathiresan/policy-constrained-edge-batch-admission-window-planner/api/v1alpha1" // day 8

type ReasonCode string

const (
	ResourceCPUInsufficient    ReasonCode = "RESOURCE_CPU_INSUFFICIENT"
	ResourceMemoryInsufficient ReasonCode = "RESOURCE_MEMORY_INSUFFICIENT"
	GPUInsufficient            ReasonCode = "GPU_INSUFFICIENT"
	ConnectivityTooLow         ReasonCode = "CONNECTIVITY_TOO_LOW"
	RegionNotAllowed           ReasonCode = "REGION_NOT_ALLOWED"
	ClusterTypeNotAllowed      ReasonCode = "CLUSTER_TYPE_NOT_ALLOWED"
	EdgeNotAllowed             ReasonCode = "EDGE_NOT_ALLOWED"
	PolicyGPULimit             ReasonCode = "POLICY_GPU_LIMIT"
	MissingTimelineData        ReasonCode = "MISSING_TIMELINE_DATA"
	MinAvailableNotSatisfied   ReasonCode = "MIN_AVAILABLE_NOT_SATISFIED"
)