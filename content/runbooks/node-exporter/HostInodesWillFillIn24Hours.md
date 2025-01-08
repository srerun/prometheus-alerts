---
title: HostInodesWillFillIn24Hours
description: Troubleshooting for alert HostInodesWillFillIn24Hours
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostInodesWillFillIn24Hours

Filesystem is predicted to run out of inodes within the next 24 hours at current write rate

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostInodesWillFillIn24Hours" %}}

{{% comment %}}

```yaml
alert: HostInodesWillFillIn24Hours
expr: (node_filesystem_files_free{fstype!="msdosfs"} / node_filesystem_files{fstype!="msdosfs"} * 100 < 10 and predict_linear(node_filesystem_files_free{fstype!="msdosfs"}[1h], 24 * 3600) < 0 and ON (instance, device, mountpoint) node_filesystem_readonly{fstype!="msdosfs"} == 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host inodes will fill in 24 hours (instance {{ $labels.instance }})
    description: |-
        Filesystem is predicted to run out of inodes within the next 24 hours at current write rate
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostInodesWillFillIn24Hours.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
