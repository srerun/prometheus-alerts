---
title: ClickhouseNoAvailableReplicas
description: Troubleshooting for alert ClickhouseNoAvailableReplicas
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseNoAvailableReplicas

## Meaning
[//]: # "Short paragraph that explains what the alert means"
No available replicas in ClickHouse.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseNoAvailableReplicas" %}}

{{% comment %}}

```yaml
alert: ClickhouseNoAvailableReplicas
expr: ClickHouseErrorMetric_NO_AVAILABLE_REPLICA == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: ClickHouse No Available Replicas (instance {{ $labels.instance }})
    description: |-
        No available replicas in ClickHouse.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseNoAvailableReplicas.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
