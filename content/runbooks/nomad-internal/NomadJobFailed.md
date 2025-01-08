---
title: NomadJobFailed
description: Troubleshooting for alert NomadJobFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadJobFailed

Nomad job failed

<details>
  <summary>Alert Rule</summary>

{{% rule "nomad/nomad-internal.yml" "NomadJobFailed" %}}

{{% comment %}}

```yaml
alert: NomadJobFailed
expr: nomad_nomad_job_summary_failed > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Nomad job failed (instance {{ $labels.instance }})
    description: |-
        Nomad job failed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nomad-internal/NomadJobFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "NomadJobFailed":

## Meaning

This alert indicates that a Nomad job has failed. Nomad is a distributed job scheduler and runner that manages and runs tasks and services. A failed job can impact the overall health and availability of the system.

## Impact

The impact of a failed Nomad job can be significant, depending on the job's purpose and the services it provides. Some possible consequences include:

* Disruption to critical services
* Loss of data or processing
* Increased latency or errors
* Inability to scale or deploy new services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Nomad UI or CLI to identify the failed job and its associated tasks.
2. Review the job's configuration and dependencies to identify potential causes of the failure.
3. Check the Nomad agent logs for errors or warnings related to the failed job.
4. Verify that the job's dependencies, such as databases or file systems, are available and functional.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address the root cause of the failure, such as a misconfigured job or a dependency issue.
2. Restart the failed job or task, if possible.
3. Check the job's configuration and dependencies to ensure they are correct and functional.
4. Consider implementing retries or fallbacks for failed jobs to minimize downtime and data loss.
5. Monitor the job's status and performance to ensure it is running correctly and make any necessary adjustments.