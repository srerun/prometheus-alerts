---
title: HostOomKillDetected
description: Troubleshooting for alert HostOomKillDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOomKillDetected

OOM kill detected

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostOomKillDetected" %}}

{{% comment %}}

```yaml
alert: HostOomKillDetected
expr: (increase(node_vmstat_oom_kill[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host OOM kill detected (instance {{ $labels.instance }})
    description: |-
        OOM kill detected
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostOomKillDetected.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostOomKillDetected` alert rule:

## Meaning

The `HostOomKillDetected` alert is triggered when a host running out of memory (OOM) kills a process to free up memory. This alert indicates that the host is experiencing memory pressure and is taking drastic measures to recover. This can lead to unexpected behavior, data loss, and performance degradation.

## Impact

The impact of this alert can be significant, as it can lead to:

* Unpredictable behavior of applications and services running on the host
* Data loss or corruption due to processes being killed abruptly
* Performance degradation and slowdowns due to memory constraints
* Potential security risks if sensitive data is lost or exposed

## Diagnosis

To diagnose the cause of the OOM kill, follow these steps:

* Check the system logs for error messages related to memory allocation and deallocation
* Verify that the host's memory usage is within expected limits
* Investigate recent changes to the system, such as new deployments or updates
* Review the output of the `node_vmstat_oom_kill` metric to determine the frequency and severity of OOM kills
* Use tools like `top` or `htop` to identify processes consuming excessive memory

## Mitigation

To mitigate the effects of an OOM kill, follow these steps:

* immediately investigate and address the root cause of the memory pressure
* Consider adding more memory to the host or optimizing memory usage
* Implement memory monitoring and alerting to detect early signs of memory pressure
* Consider implementing automatic restarts or failovers for critical services
* Review and optimize system configuration and resource allocation to prevent future OOM kills

Additionally, consider implementing long-term solutions to prevent OOM kills, such as:

* Implementing memory-efficient coding practices
* Optimizing database and cache configurations
* Using memory-optimized storage solutions
* Implementing load balancing and horizontal scaling to distribute workload and reduce memory pressure.