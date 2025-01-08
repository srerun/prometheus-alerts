---
title: ClickhouseHighNetworkTraffic
description: Troubleshooting for alert ClickhouseHighNetworkTraffic
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseHighNetworkTraffic

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Network traffic is unusually high, may affect cluster performance.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseHighNetworkTraffic" %}}

{{% comment %}}

```yaml
alert: ClickhouseHighNetworkTraffic
expr: ClickHouseMetrics_NetworkSend > 250 or ClickHouseMetrics_NetworkReceive > 250
for: 5m
labels:
    severity: warning
annotations:
    summary: ClickHouse High Network Traffic (instance {{ $labels.instance }})
    description: |-
        Network traffic is unusually high, may affect cluster performance.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseHighNetworkTraffic.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
