---
title: NomadBlockedEvaluation
description: Troubleshooting for alert NomadBlockedEvaluation
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadBlockedEvaluation

Nomad blocked evaluation

<details>
  <summary>Alert Rule</summary>

{{% rule "nomad/nomad-internal.yml" "NomadBlockedEvaluation" %}}

{{% comment %}}

```yaml
alert: NomadBlockedEvaluation
expr: nomad_nomad_blocked_evals_total_blocked > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Nomad blocked evaluation (instance {{ $labels.instance }})
    description: |-
        Nomad blocked evaluation
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nomad-internal/NomadBlockedEvaluation.md

```

{{% /comment %}}

</details>


Here is a runbook for the NomadBlockedEvaluation alert:

## Meaning

The NomadBlockedEvaluation alert indicates that Nomad, a clustering and orchestration tool, has blocked an evaluation. This means that Nomad has temporarily stopped scheduling new tasks or deployments due to internal inconsistencies or errors. This can lead to delays or failures in task execution, and may impact the overall performance and reliability of the system.

## Impact

The impact of this alert is moderate to high, as it can cause delays or failures in task execution, leading to:

* Delays in deployment of new services or applications
* Inconsistent or failed deployments
* Increased latency or errors in service requests
* Potential data loss or corruption

## Diagnosis

To diagnose the root cause of the NomadBlockedEvaluation alert, follow these steps:

1. Check the Nomad logs for errors or warnings related to the evaluation blocking.
2. Verify that the Nomad cluster is healthy and all nodes are online.
3. Check the Nomad configuration files for any inconsistencies or errors.
4. Run the `nomad node status` command to verify the status of the nodes in the cluster.
5. Run the `nomad eval status` command to verify the status of the evaluations in the cluster.

## Mitigation

To mitigate the NomadBlockedEvaluation alert, follow these steps:

1. Identify and fix the underlying cause of the evaluation blocking, based on the diagnosis steps above.
2. Restart the Nomad agent on the affected node(s) to clear the blocked evaluation.
3. Verify that the Nomad cluster is healthy and all nodes are online.
4. Verify that the Nomad configuration files are consistent and error-free.
5. Implement additional logging and monitoring to detect and prevent similar issues in the future.
6. Consider implementing automated rollbacks or failovers to minimize the impact of blocked evaluations on the system.