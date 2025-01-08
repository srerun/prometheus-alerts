---
title: HostRaidArrayGotInactive
description: Troubleshooting for alert HostRaidArrayGotInactive
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostRaidArrayGotInactive

RAID array {{ $labels.device }} is in a degraded state due to one or more disk failures. The number of spare drives is insufficient to fix the issue automatically.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostRaidArrayGotInactive" %}}

{{% comment %}}

```yaml
alert: HostRaidArrayGotInactive
expr: (node_md_state{state="inactive"} > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: critical
annotations:
    summary: Host RAID array got inactive (instance {{ $labels.instance }})
    description: |-
        RAID array {{ $labels.device }} is in a degraded state due to one or more disk failures. The number of spare drives is insufficient to fix the issue automatically.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostRaidArrayGotInactive.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostRaidArrayGotInactive` alert rule:

## Meaning

The `HostRaidArrayGotInactive` alert is triggered when the RAID array on a host becomes inactive due to one or more disk failures. This means that the RAID array is no longer functioning properly and data may be at risk.

## Impact

The impact of an inactive RAID array is high, as it can lead to data loss or corruption. Additionally, if the failed disk(s) are not replaced promptly, the entire RAID array may become unavailable, causing significant disruptions to the system or application.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus alert label `device` to identify the specific RAID array that is affected.
2. Log in to the affected host and check the RAID array status using the `mdadm` command or a similar tool.
3. Identify the failed disk(s) and determine the cause of the failure (e.g., hardware fault, software issue, etc.).
4. Check the system logs for any error messages related to the RAID array or disk failures.

## Mitigation

To mitigate the issue, follow these steps:

1. Replace the failed disk(s) with new ones to ensure the RAID array has sufficient redundancy.
2. Rebuild the RAID array using the `mdadm` command or a similar tool.
3. Monitor the RAID array status to ensure it is functioning properly and data is being written correctly.
4. Consider implementing additional monitoring and alerting for disk failures to prevent similar issues in the future.

Remember to also investigate the root cause of the disk failure to prevent similar issues from occurring in the future.