---
title: HostOutOfInodes
description: Troubleshooting for alert HostOutOfInodes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOutOfInodes

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Disk is almost running out of available inodes (< 10% left)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostOutOfInodes
expr: (node_filesystem_files_free{fstype!="msdosfs"} / node_filesystem_files{fstype!="msdosfs"} * 100 < 10 and ON (instance, device, mountpoint) node_filesystem_readonly == 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host out of inodes (instance {{ $labels.instance }})
    description: |-
        Disk is almost running out of available inodes (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostOutOfInodes

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
