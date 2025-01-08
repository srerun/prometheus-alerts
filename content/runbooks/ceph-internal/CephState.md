---
title: CephState
description: Troubleshooting for alert CephState
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephState

Ceph instance unhealthy

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephState" %}}

{{% comment %}}

```yaml
alert: CephState
expr: ceph_health_status != 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph State (instance {{ $labels.instance }})
    description: |-
        Ceph instance unhealthy
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephState.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `CephState`:

## Meaning

The `CephState` alert is triggered when the Ceph cluster's health status is not healthy (i.e., `ceph_health_status != 0`). This indicates that there is an issue with the Ceph cluster that needs to be addressed promptly to prevent data loss or unavailability.

## Impact

The impact of this alert is critical, as an unhealthy Ceph cluster can lead to:

* Data loss or corruption
* Unavailability of storage resources
* Disruption to dependent services and applications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Ceph cluster's status using `ceph -s` command
2. Review the Ceph log files for errors or warnings
3. Verify that all Ceph nodes are online and reachable
4. Check for any ongoing maintenance or upgrade activities that may be causing the issue
5. Review the `LABELS` and `VALUE` provided in the alert to identify the specific instance and error code

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address the root cause of the health issue (e.g., fix any hardware or software issues, resolve network connectivity problems, etc.)
2. Restart any failed Ceph services or nodes
3. Run `ceph heal` command to initiate the healing process
4. Verify that the Ceph cluster's health status has returned to normal
5. If the issue persists, consider escalating to a senior engineer or Ceph expert for further assistance.

Remember to update the runbook with any additional steps or procedures specific to your environment.