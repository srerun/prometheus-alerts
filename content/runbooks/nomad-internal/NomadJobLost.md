---
title: NomadJobLost
description: Troubleshooting for alert NomadJobLost
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadJobLost

Nomad job lost

<details>
  <summary>Alert Rule</summary>

{{% rule "nomad/nomad-internal.yml" "NomadJobLost" %}}

{{% comment %}}

```yaml
alert: NomadJobLost
expr: nomad_nomad_job_summary_lost > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Nomad job lost (instance {{ $labels.instance }})
    description: |-
        Nomad job lost
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nomad-internal/NomadJobLost.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "NomadJobLost":

## Meaning

The NomadJobLost alert is triggered when a Nomad job is lost, meaning it is no longer running or has disappeared from the Nomad cluster. This alert is critical because it can indicate a problem with the job itself, the Nomad cluster, or the underlying infrastructure.

## Impact

The impact of a lost Nomad job can be significant, depending on the importance of the job and the services it provides. If the job is critical to the operation of the system, its loss can lead to:

* Service disruptions or outages
* Data loss or corruption
* Increased latency or errors
* Security vulnerabilities

## Diagnosis

To diagnose the cause of the lost Nomad job, follow these steps:

1. Check the Nomad job logs for errors or warnings that may indicate the reason for the job's disappearance.
2. Verify that the Nomad cluster is healthy and all nodes are online.
3. Check the system logs for any errors or issues that may be related to the job's loss.
4. Verify that the job's configuration and deployment files are correct and up-to-date.
5. Check for any recent changes to the Nomad cluster or underlying infrastructure that may have caused the job to fail.

## Mitigation

To mitigate the effects of a lost Nomad job, follow these steps:

1. Identify the root cause of the job's disappearance and address it immediately.
2. Restart the Nomad job if possible, or redeploy it with the correct configuration.
3. Verify that the job is running correctly and is healthy.
4. Implement monitoring and logging to prevent similar issues in the future.
5. Consider implementing redundancy or failover mechanisms to minimize the impact of job losses.

Remember to update the runbook with specific steps and procedures relevant to your organization's Nomad cluster and infrastructure.