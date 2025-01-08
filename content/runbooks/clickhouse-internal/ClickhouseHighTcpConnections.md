---
title: ClickhouseHighTcpConnections
description: Troubleshooting for alert ClickhouseHighTcpConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseHighTcpConnections

High number of TCP connections, indicating heavy client or inter-cluster communication.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseHighTcpConnections" %}}

{{% comment %}}

```yaml
alert: ClickhouseHighTcpConnections
expr: ClickHouseMetrics_TCPConnection > 400
for: 5m
labels:
    severity: warning
annotations:
    summary: ClickHouse High TCP Connections (instance {{ $labels.instance }})
    description: |-
        High number of TCP connections, indicating heavy client or inter-cluster communication.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseHighTcpConnections.md

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
