---
title: HostDiskWillFillIn24Hours
description: Troubleshooting for alert HostDiskWillFillIn24Hours
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostDiskWillFillIn24Hours

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Filesystem is predicted to run out of space within the next 24 hours at current write rate

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostDiskWillFillIn24Hours" %}}

{{% comment %}}

```yaml
alert: HostDiskWillFillIn24Hours
expr: ((node_filesystem_avail_bytes * 100) / node_filesystem_size_bytes < 10 and ON (instance, device, mountpoint) predict_linear(node_filesystem_avail_bytes{fstype!~"tmpfs"}[1h], 24 * 3600) < 0 and ON (instance, device, mountpoint) node_filesystem_readonly == 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host disk will fill in 24 hours (instance {{ $labels.instance }})
    description: |-
        Filesystem is predicted to run out of space within the next 24 hours at current write rate
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostDiskWillFillIn24Hours.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"

The impact of a disk filling up can range from applications failing to the host panicing and shutting down depending on which disk is filling and the operating system.

## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"

The alert will provide the host and filesystem that is in danger of filling up.  For Windows machines, login and check the storage manager for the filesystem indicated.

On Linux machines, login and run
```bash
df -h
```
this will show the utilization of all mounted filesystems.

To determine where the majority of the utilization is, change into the mounted directory and run:

```bash
sudo df -h *
```
This will list out each directory and the amount of space it is using.  



## Mitigation
[//]: # "The steps necessary to resolve the alert"

There are two options to mitigate the issue.  

1. Add more space to the filesystem, eg.
   - grow the disk on a vm
   - add an additional drive and span them on the filesystem
2. Free up space on the affected drive
   - using the output of the `df -h` command, locate the directories using the most space
   - check the files in those directories to determine what can be deleted
