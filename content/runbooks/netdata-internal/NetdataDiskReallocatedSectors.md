---
title: NetdataDiskReallocatedSectors
description: Troubleshooting for alert NetdataDiskReallocatedSectors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataDiskReallocatedSectors

Reallocated sectors on disk

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataDiskReallocatedSectors" %}}

{{% comment %}}

```yaml
alert: NetdataDiskReallocatedSectors
expr: increase(netdata_smartd_log_reallocated_sectors_count_sectors_average[1m]) > 0
for: 0m
labels:
    severity: info
annotations:
    summary: Netdata disk reallocated sectors (instance {{ $labels.instance }})
    description: |-
        Reallocated sectors on disk
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataDiskReallocatedSectors.md

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
