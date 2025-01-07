---
title: HostDiskWillFillIn24Hours
description: Troubleshooting for alert HostDiskWillFillIn24Hours
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostDiskWillFillIn24Hours

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Filesystem is predicted to run out of space within the next 24 hours at current write rate

<details>
  <summary>Alert Rule</summary>

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
    runbook: http://wiki.ringsq.io/runbook/HostDiskWillFillIn24Hours

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
