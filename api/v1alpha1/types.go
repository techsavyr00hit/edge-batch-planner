// day 3 - workload API types

// Package v1alpha1 contains lightweight Kubernetes-native API models for the planner.
package v1alpha1

import "k8s.io/apimachinery/pkg/api/resource"

type BatchWorkload struct {
	Name          string          `yaml:"name"`
	Resources     ResourceRequest `yaml:"resources"`
	DurationHours int             `yaml:"durationHours"`
	MinAvailable  int             `yaml:"minAvailable"`
	Priority      string          `yaml:"priority"`
	Placement     Placement       `yaml:"placement"`
	Connectivity  ConnectivityReq `yaml:"connectivity"`
}

type ResourceRequest struct {
	CPU    resource.Quantity `yaml:"cpu"`
	Memory resource.Quantity `yaml:"memory"`
	GPU    resource.Quantity `yaml:"gpu"`
}

type Placement struct {
	AllowedRegions      []string `yaml:"allowedRegions"`
	AllowedClusterTypes []string `yaml:"allowedClusterTypes"`
	PreferredCluster    string   `yaml:"preferredCluster"`
}

type ConnectivityReq struct {
	Minimum int `yaml:"minimum"`
}

// day 4 - cluster forecast API types
type ClusterForecast struct {
	Name      string          `yaml:"name"`
	Type      string          `yaml:"type"`
	Region    string          `yaml:"region"`
	Resources ResourceRequest `yaml:"resources"`
	Timeline  []TimeSlot      `yaml:"timeline"`
}

type TimeSlot struct {
	Hour             int               `yaml:"hour"`
	CPUAvailable     resource.Quantity `yaml:"cpuAvailable"`
	MemoryAvail      resource.Quantity `yaml:"memoryAvailable"`
	GPUAvailable     resource.Quantity `yaml:"gpuAvailable"`
	Connectivity     int               `yaml:"connectivity"`
	WorkersAvailable int               `yaml:"workersAvailable"`
}

// day 5 - admission policy API type
type AdmissionPolicy struct {
	Name                string            `yaml:"name"`
	AllowedRegions      []string          `yaml:"allowedRegions"`
	AllowedClusterTypes []string          `yaml:"allowedClusterTypes"`
	MinimumConnectivity int               `yaml:"minimumConnectivity"`
	AllowEdge           bool              `yaml:"allowEdge"`
	MaximumGPU          resource.Quantity `yaml:"maximumGPU"`
}
