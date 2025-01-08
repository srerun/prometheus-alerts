---
title: HostUnusualDiskWriteRate
description: Troubleshooting for alert HostUnusualDiskWriteRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskWriteRate

Disk is probably writing too much data (> 50 MB/s)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualDiskWriteRate" %}}

{{% comment %}}

```yaml
alert: HostUnusualDiskWriteRate
expr: (sum by (instance) (rate(node_disk_written_bytes_total[2m])) / 1024 / 1024 > 50) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host unusual disk write rate (instance {{ $labels.instance }})
    description: |-
        Disk is probably writing too much data (> 50 MB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualDiskWriteRate.md

```

{{% /comment %}}

</details>


Here is the runbook for the `HostUnusualDiskWriteRate` alert rule:

## Meaning
The `HostUnusualDiskWriteRate` alert is triggered when a host's disk write rate exceeds 50 MB/s over a 2-minute period. This alert indicates that a host is writing data to disk at an unusually high rate, which may be a sign of abnormal system behavior or a potential issue.

## Impact
If left unaddressed, an unusual disk write rate can lead to:

* Disk I/O bottlenecks, causing slow system performance and potentially affecting application responsiveness
* Increased wear and tear on disk hardware, potentially reducing its lifespan
* Potential data loss or corruption if the underlying issue is not addressed

## Diagnosis
To diagnose the root cause of the unusual disk write rate, follow these steps:

1. **Investigate the host**: Check the host's system logs, disk usage, and running processes to identify any unusual activity or resource-intensive processes.
2. **Verify disk usage**: Use tools like `df` or `du` to check disk usage and identify which files or directories are contributing to the high write rate.
3. **Check for disk errors**: Run disk diagnostic tools like `smartctl` or `fsck` to identify any disk errors or corruption.
4. **Review system configuration**: Check system configuration files, such as `sysctl` and `fstab`, to ensure that they are correctly configured and not contributing to the high write rate.

## Mitigation
To mitigate the effects of an unusual disk write rate, follow these steps:

1. **Identify and terminate resource-intensive processes**: Use tools like `top` or `htop` to identify and terminate any resource-intensive processes that may be contributing to the high write rate.
2. **Reduce disk usage**: Free up disk space by deleting unnecessary files or relocating data to a different storage location.
3. **Adjust system configuration**: Adjust system configuration files to optimize disk performance and reduce the write rate.
4. **Monitor disk performance**: Closely monitor disk performance and adjust mitigation strategies as needed to prevent further issues.

Remember to investigate and address the root cause of the unusual disk write rate to prevent similar issues from occurring in the future.