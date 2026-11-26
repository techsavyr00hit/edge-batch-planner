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

type WindowResult struct {
	Cluster    string
	Start, End int
	Feasible   bool
	Score      float64
	Reasons    []Reason
	Components ScoreComponents
}
type Reason struct {
	Code    ReasonCode
	Message string
}
type ScoreComponents struct{ ResourceHeadroom, Connectivity, PlacementPreference float64 }
type Analysis struct {
	Workload string
	Windows  []WindowResult
}

// day 8 - timeline slot helper
func SlotMap(c v1.ClusterForecast) map[int]v1.TimeSlot {
	m := make(map[int]v1.TimeSlot, len(c.Timeline))
	for _, s := range c.Timeline {
		m[s.Hour] = s
	}
	return m
}