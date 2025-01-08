---
title: HostEdacCorrectableErrorsDetected
description: Troubleshooting for alert HostEdacCorrectableErrorsDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostEdacCorrectableErrorsDetected

Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} correctable memory errors reported by EDAC in the last 5 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostEdacCorrectableErrorsDetected" %}}

{{% comment %}}

```yaml
alert: HostEdacCorrectableErrorsDetected
expr: (increase(node_edac_correctable_errors_total[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: info
annotations:
    summary: Host EDAC Correctable Errors detected (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} correctable memory errors reported by EDAC in the last 5 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostEdacCorrectableErrorsDetected.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "HostEdacCorrectableErrorsDetected":

## Meaning

This alert is triggered when a host has detected correctable memory errors reported by EDAC (Error Detection and Correction) in the last 5 minutes. EDAC is a Linux kernel module that detects and corrects memory errors. Correctable errors are errors that can be recovered from without system failure, but may indicate a hardware issue.

## Impact

Correctable memory errors can impact system performance and reliability. If left unchecked, these errors can lead to system crashes, data corruption, or other unexpected behavior. This alert allows system administrators to take proactive measures to investigate and resolve the issue before it becomes more severe.

## Diagnosis

To diagnose the issue, follow these steps:

* Check the system logs for any error messages related to EDAC or memory errors.
* Verify the system's memory configuration and ensure that it is correctly installed and seated.
* Run a memory test using tools like `memtest86+` to identify any faulty memory modules.
* Check the system's event logs for any other errors or warnings that may be related to the memory errors.
* If the issue persists, consider replacing the faulty memory module or seeking assistance from the system manufacturer or a qualified technical expert.

## Mitigation

To mitigate the issue, follow these steps:

* Restart the system to clear any temporary errors.
* If the issue persists, schedule a maintenance window to replace the faulty memory module.
* Consider implementing additional monitoring and logging to detect memory errors earlier.
* Review system logs and event logs regularly to identify any other potential issues before they become more severe.