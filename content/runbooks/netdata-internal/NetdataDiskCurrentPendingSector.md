---
title: NetdataDiskCurrentPendingSector
description: Troubleshooting for alert NetdataDiskCurrentPendingSector
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataDiskCurrentPendingSector

Disk current pending sector

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataDiskCurrentPendingSector" %}}

{{% comment %}}

```yaml
alert: NetdataDiskCurrentPendingSector
expr: netdata_smartd_log_current_pending_sector_count_sectors_average > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Netdata disk current pending sector (instance {{ $labels.instance }})
    description: |-
        Disk current pending sector
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataDiskCurrentPendingSector.md

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
