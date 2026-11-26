package config

import (
	"fmt"
	v1 "github.com/roohitkathiresan/policy-constrained-edge-batch-admission-window-planner/api/v1alpha1"
	"os"
	"sigs.k8s.io/yaml"
)

type ClusterFile struct {
	Clusters []v1.ClusterForecast `yaml:"clusters"`
}
type WorkloadFile struct {
	Workload v1.BatchWorkload `yaml:"workload"`
}
type PolicyFile struct {
	Policy v1.AdmissionPolicy `yaml:"policy"`
}

func Load(path string, dst any) error {
	b, e := os.ReadFile(path)
	if e != nil {
		return e
	}
	if e = yaml.Unmarshal(b, dst); e != nil {
		return fmt.Errorf("parse %s: %w", path, e)
	}
	return nil
}

func Validate(c ClusterFile, w WorkloadFile, p PolicyFile) error {
	if len(c.Clusters) == 0 {
		return fmt.Errorf("no clusters supplied")
	}
	if w.Workload.Name == "" {
		return fmt.Errorf("workload name is required")
	}
	if w.Workload.DurationHours < 1 {
		return fmt.Errorf("durationHours must be positive")
	}
	if w.Workload.MinAvailable < 1 {
		return fmt.Errorf("minAvailable must be positive")
	}
	if p.Policy.MinimumConnectivity < 0 || p.Policy.MinimumConnectivity > 100 {
		return fmt.Errorf("minimumConnectivity must be 0..100")
	}
	return nil
}