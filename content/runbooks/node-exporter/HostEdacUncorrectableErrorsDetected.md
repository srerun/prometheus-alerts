---
title: HostEdacUncorrectableErrorsDetected
description: Troubleshooting for alert HostEdacUncorrectableErrorsDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostEdacUncorrectableErrorsDetected

Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} uncorrectable memory errors reported by EDAC in the last 5 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostEdacUncorrectableErrorsDetected" %}}

{{% comment %}}

```yaml
alert: HostEdacUncorrectableErrorsDetected
expr: (node_edac_uncorrectable_errors_total > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host EDAC Uncorrectable Errors detected (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} uncorrectable memory errors reported by EDAC in the last 5 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostEdacUncorrectableErrorsDetected.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `HostEdacUncorrectableErrorsDetected`:

## Meaning

This alert is triggered when the Node Exporter metric `node_edac_uncorrectable_errors_total` reports one or more uncorrectable memory errors on a host. EDAC (Error Detection and Correction) is a Linux kernel module that detects and reports memory errors. Uncorrectable errors are those that cannot be recovered from and may indicate a hardware issue.

## Impact

The impact of this alert is potentially high, as uncorrectable memory errors can lead to data corruption, system crashes, or even data loss. If left unaddressed, these errors can cause system instability, affecting the overall reliability and availability of the host.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus metrics to identify the specific host and instance affected.
2. Review system logs for error messages related to EDAC and memory errors.
3. Verify that the EDAC kernel module is properly configured and loaded.
4. Run diagnostic tools, such as `edac-util` or `memtester`, to further investigate memory issues.
5. Check system hardware for signs of wear or damage, such as overheating, electrical issues, or physical damage to the RAM.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately shut down the affected host to prevent further data corruption or system crashes.
2. Replace any faulty RAM modules or other hardware components that may be causing the errors.
3. Perform a thorough system memory test to identify and isolate any issues.
4. Verify that the EDAC kernel module is properly configured and loaded to detect and report memory errors.
5. Consider implementing additional memory redundancy or error correction mechanisms to increase system reliability.

Remember to adjust the mitigation steps based on the specific system configuration, hardware, and requirements.