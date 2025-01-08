---
title: ZfsCollectorFailed
description: Troubleshooting for alert ZfsCollectorFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsCollectorFailed

ZFS collector for {{ $labels.instance }} has failed to collect information

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/zfs_exporter.yml" "ZfsCollectorFailed" %}}

{{% comment %}}

```yaml
alert: ZfsCollectorFailed
expr: zfs_scrape_collector_success != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: ZFS collector failed (instance {{ $labels.instance }})
    description: |-
        ZFS collector for {{ $labels.instance }} has failed to collect information
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/zfs_exporter/ZfsCollectorFailed.md

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
