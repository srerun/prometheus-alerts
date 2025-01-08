---
title: HighNumberOfSnapshots
description: Troubleshooting for alert HighNumberOfSnapshots
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighNumberOfSnapshots

High snapshots number on {{ $labels.instance }}: {{ $value }}

<details>
  <summary>Alert Rule</summary>

{{% rule "vmware/pryorda-vmware-exporter.yml" "HighNumberOfSnapshots" %}}

{{% comment %}}

```yaml
alert: HighNumberOfSnapshots
expr: vmware_vm_snapshots > 3
for: 30m
labels:
    severity: warning
annotations:
    summary: High Number of Snapshots (instance {{ $labels.instance }})
    description: |-
        High snapshots number on {{ $labels.instance }}: {{ $value }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pryorda-vmware-exporter/HighNumberOfSnapshots.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "HighNumberOfSnapshots":

## Meaning

The "HighNumberOfSnapshots" alert is triggered when the number of snapshots for a VMware virtual machine exceeds 3. This alert is raised to warn of a potential issue with snapshot management, which can lead to performance degradation, storage consumption, and even virtual machine corruption.

## Impact

A high number of snapshots can cause:

* Performance degradation: Excessive snapshots can slow down virtual machine performance, leading to slower application response times and decreased productivity.
* Storage consumption: Snapshots consume storage space, which can lead to storage capacity issues and increased storage costs.
* Virtual machine corruption: A large number of snapshots can increase the risk of virtual machine corruption, making it difficult or impossible to recover the virtual machine in case of a failure.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected virtual machine: Check the alert label "instance" to determine which virtual machine is experiencing the high number of snapshots.
2. Check snapshot count: Verify the number of snapshots for the affected virtual machine using the vSphere Web Client or vSphere API.
3. Investigate snapshot creation: Determine the reason for the high number of snapshots. Is it due to an automated process, a user error, or a configuration issue?
4. Review snapshot retention: Check the snapshot retention policy to ensure it is set correctly and not causing the accumulation of snapshots.

## Mitigation

To resolve the issue, follow these steps:

1. Remove unnecessary snapshots: Delete snapshots that are no longer needed, following the organization's snapshot retention policy.
2. Adjust snapshot creation: Modify the process or configuration that is causing the high number of snapshots to prevent further accumulation.
3. Implement snapshot management: Establish a regular snapshot management process to maintain a healthy snapshot count and prevent future occurrences of this issue.
4. Monitor snapshot count: Continuously monitor the snapshot count for the affected virtual machine to ensure the issue is resolved and does not reoccur.

By following these steps, you should be able to resolve the high number of snapshots and prevent future occurrences of this issue.