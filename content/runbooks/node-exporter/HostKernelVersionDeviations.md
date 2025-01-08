---
title: HostKernelVersionDeviations
description: Troubleshooting for alert HostKernelVersionDeviations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostKernelVersionDeviations

Different kernel versions are running

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostKernelVersionDeviations" %}}

{{% comment %}}

```yaml
alert: HostKernelVersionDeviations
expr: (count(sum(label_replace(node_uname_info, "kernel", "$1", "release", "([0-9]+.[0-9]+.[0-9]+).*")) by (kernel)) > 1) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 6h
labels:
    severity: warning
annotations:
    summary: Host kernel version deviations (instance {{ $labels.instance }})
    description: |-
        Different kernel versions are running
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostKernelVersionDeviations.md

```

{{% /comment %}}

</details>


Here is the runbook for the `HostKernelVersionDeviations` alert rule:

## Meaning

This alert is triggered when there are multiple kernel versions running on hosts in the cluster. This can indicate a potential security risk or compatibility issues between hosts.

## Impact

The impact of this issue can be:

* Security vulnerabilities: Running different kernel versions can expose the cluster to known security vulnerabilities.
* Inconsistent behavior: Different kernel versions can lead to inconsistent behavior between hosts, making troubleshooting and maintenance more challenging.
* Performance issues: Mismatched kernel versions can cause performance issues, affecting the overall cluster performance.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the alert description for the instance and kernel versions affected.
2. Verify the kernel versions running on each host using `node_uname_info` metrics.
3. Identify the hosts with inconsistent kernel versions.
4. Check the kernel version installed on each host using `uname -r` command.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the recommended kernel version for the cluster.
2. Update the kernel version on each host to the recommended version.
3. Verify that all hosts are running the same kernel version using `node_uname_info` metrics.
4. Implement a process to ensure kernel version consistency across the cluster, such as automating kernel updates or implementing a kernel version management policy.

Additionally, consider implementing a monitoring and alerting system to detect kernel version deviations in the future.