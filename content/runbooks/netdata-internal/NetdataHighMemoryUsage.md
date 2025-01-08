---
title: NetdataHighMemoryUsage
description: Troubleshooting for alert NetdataHighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataHighMemoryUsage

Netdata high memory usage (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataHighMemoryUsage" %}}

{{% comment %}}

```yaml
alert: NetdataHighMemoryUsage
expr: 100 / netdata_system_ram_MiB_average * netdata_system_ram_MiB_average{dimension=~"free|cached"} < 20
for: 5m
labels:
    severity: warning
annotations:
    summary: Netdata high memory usage (instance {{ $labels.instance }})
    description: |-
        Netdata high memory usage (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataHighMemoryUsage.md

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
