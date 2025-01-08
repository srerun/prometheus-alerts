---
title: ZfsOfflinePool
description: Troubleshooting for alert ZfsOfflinePool
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsOfflinePool

A ZFS zpool is in a unexpected state: {{ $labels.state }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/node-exporter.yml" "ZfsOfflinePool" %}}

{{% comment %}}

```yaml
alert: ZfsOfflinePool
expr: node_zfs_zpool_state{state!="online"} > 0
for: 1m
labels:
    severity: critical
annotations:
    summary: ZFS offline pool (instance {{ $labels.instance }})
    description: |-
        A ZFS zpool is in a unexpected state: {{ $labels.state }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/ZfsOfflinePool.md

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
