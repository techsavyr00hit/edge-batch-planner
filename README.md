# Policy-Constrained Edge Batch Admission Window Planner

An independent, Kubernetes-native **planning and analysis engine** that ranks future time windows in which a batch workload can safely be admitted to a cloud or edge cluster.

> This project is an independent planning/analysis tool inspired by concepts from Karmada, KubeEdge, Volcano and Kyverno. It does not replace or integrate with those projects.

## What it does

Given YAML forecasts, a workload and a policy, the CLI evaluates every complete sliding window. A window is feasible only when resource availability, connectivity, placement, policy and the workload's minimum-availability requirement pass for every required hour.

It **does not schedule, deploy, migrate or execute workloads**.
