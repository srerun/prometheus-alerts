---
title: NomadJobQueued
description: Troubleshooting for alert NomadJobQueued
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadJobQueued

Nomad job queued

<details>
  <summary>Alert Rule</summary>

{{% rule "nomad/nomad-internal.yml" "NomadJobQueued" %}}

{{% comment %}}

```yaml
alert: NomadJobQueued
expr: nomad_nomad_job_summary_queued > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Nomad job queued (instance {{ $labels.instance }})
    description: |-
        Nomad job queued
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nomad-internal/NomadJobQueued.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "NomadJobQueued":

## Meaning

The NomadJobQueued alert is triggered when a Nomad job is queued and not running. This can indicate that the job is unable to start due to resource constraints, configuration issues, or other problems.

## Impact

The impact of this alert is that the Nomad job is not running, which can lead to:

* Delays in task execution
* Increased latency
* Potential data loss or inconsistencies
* Impact on dependent services or applications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Nomad job configuration to ensure it is correct and up-to-date.
2. Verify that the required resources (e.g. CPU, memory, network) are available for the job to run.
3. Investigate any recent changes to the Nomad cluster or job configuration that may be causing the issue.
4. Check the Nomad logs for any errors or warnings related to the job.
5. Verify that the job is not stuck in a failed or paused state.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Nomad job configuration and update it if necessary to ensure it is correct and up-to-date.
2. Verify that the required resources are available for the job to run and allocate additional resources if necessary.
3. Resolve any recent changes to the Nomad cluster or job configuration that may be causing the issue.
4. Restart the Nomad job or update the job configuration to allow it to run successfully.
5. Monitor the job to ensure it is running successfully and no further issues occur.

Note: This is a sample runbook and may need to be customized to fit the specific use case and environment.