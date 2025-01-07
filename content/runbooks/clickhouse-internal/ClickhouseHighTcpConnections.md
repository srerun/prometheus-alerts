---
title: ClickhouseHighTcpConnections
description: Troubleshooting for alert ClickhouseHighTcpConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseHighTcpConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High number of TCP connections, indicating heavy client or inter-cluster communication.

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ClickhouseHighTcpConnections

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
