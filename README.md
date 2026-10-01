<!-- day 1 - README title and overview -->
# Policy-Constrained Edge Batch Admission Window Planner

An independent, Kubernetes-native **planning and analysis engine** that ranks future time windows in which a batch workload can safely be admitted to a cloud or edge cluster.

> This project is an independent planning/analysis tool inspired by concepts from Karmada, KubeEdge, Volcano and Kyverno. It does not replace or integrate with those projects.

## What it does

Given YAML forecasts, a workload and a policy, the CLI evaluates every complete sliding window. A window is feasible only when resource availability, connectivity, placement, policy and the workload's minimum-availability requirement pass for every required hour.

It **does not schedule, deploy, migrate or execute workloads**.

<!-- day 2 - architecture, concepts and stack -->
## Architecture

```text
clusters.yaml ─┐
workload.yaml ─┼─> YAML loader/validator ─> Window generator
policy.yaml ───┘                              │
                                             v
                              resource + connectivity checks
                                             │
                              placement + policy checks
                                             │
                                             v
                              deterministic explainable score
                                             │
                                             v
                              ranked admission-window report
```

## Conceptual relationship

| Inspiration | Concept used | This project |
|---|---|---|
| Karmada | multi-cluster placement, affinity, regions | evaluates candidate clusters |
| KubeEdge | edge/cloud distinction, intermittent connectivity | evaluates forecast connectivity |
| Volcano | batch resources, priority, minAvailable | models batch admission requirements |
| Kyverno | policy-as-code and validation | evaluates lightweight admission policy |

The project intentionally does not implement a scheduler, controller, policy engine, or cluster orchestrator.

## Stack

- Go 1.24+
- Kubernetes `resource.Quantity` and API machinery
- Cobra CLI
- YAML
- Prometheus client library
- Go testing + Testify-ready module

No Python, Java, database, frontend, ML, Kubernetes cluster, or cloud API is required.

<!-- day 25 - install, build and quick start -->
## Install / build

```bash
git clone https://github.com/roohitkathiresan/policy-constrained-edge-batch-admission-window-planner.git
cd policy-constrained-edge-batch-admission-window-planner
go mod tidy
go build -o planner ./cmd/planner
go test ./...

# validate inputs
./planner validate --clusters examples/clusters.yaml --workload examples/workload.yaml --policy examples/policy.yaml
go vet ./...
```

## Quick start

```bash
./planner analyze \
  --clusters examples/clusters.yaml \
  --workload examples/workload.yaml \
  --policy examples/policy.yaml
```

The CLI ranks feasible windows globally and prints human-readable rejection reasons for infeasible windows.

<!-- day 20 - algorithm section -->
## Algorithm

For a workload with `durationHours = D`, each cluster timeline is treated as an ordered sequence of hourly forecasts. Every D-hour contiguous candidate is evaluated.

For each hour the planner checks:

1. CPU availability
2. Memory availability
3. GPU availability
4. Worker availability for the `minAvailable` gang-style requirement
5. Connectivity
5. Region and cluster-type placement
6. Edge policy
7. Policy GPU limit
8. Minimum-availability requirement

Missing forecast data rejects the candidate rather than inventing a value.

<!-- day 21 - scoring section -->
## Scoring

Only feasible windows are scored:

```text
score =
    0.40 * resourceHeadroom +
    0.40 * connectivityScore +
    0.20 * placementPreference
```

All components are normalized to `[0,1]`. Resource headroom is derived from the ratio between forecast availability and the workload request, capped at 1. Connectivity is the average forecast percentage normalized to 0–1. Placement preference is `1` for the explicitly preferred cluster and `0.5` otherwise.

The score is deterministic and explainable; no machine learning or optimization framework is involved.

<!-- day 13 - example input pointer -->
## Example input

See `examples/clusters.yaml`, `examples/workload.yaml`, and `examples/policy.yaml`.

<!-- day 29 - tests section -->
## Tests

```bash
go test ./...
go test -v ./...
```

The test suite covers resource feasibility, connectivity failure, policy failure, valid windows and rejection reasons. More cases can be added as the model evolves.

<!-- day 30 - limitations and future work -->
## Limitations

- Forecasts are supplied by YAML; the project does not collect telemetry.
- The current model uses hourly slots.
- `minAvailable` is checked against the forecast `workersAvailable` value at every hour; this is an admission-time gang-style feasibility check, not pod-level gang scheduling.
- Priority is retained in the workload model but does not alter the deterministic score.
- Prometheus metric definitions are intentionally small; there is no monitoring server.
- There is no actual Kubernetes scheduling or deployment.

## Future improvements

- More expressive policy rules and reason codes
- Multi-resource preference tuning
- Configurable scoring weights
- Non-hourly forecast intervals
- Optional Kubernetes API adapters for importing observations without turning the project into a scheduler
