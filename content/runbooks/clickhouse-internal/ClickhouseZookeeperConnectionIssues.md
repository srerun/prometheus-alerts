---
title: ClickhouseZookeeperConnectionIssues
description: Troubleshooting for alert ClickhouseZookeeperConnectionIssues
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseZookeeperConnectionIssues

ClickHouse is experiencing issues with ZooKeeper connections, which may affect cluster state and coordination.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseZookeeperConnectionIssues" %}}

{{% comment %}}

```yaml
alert: ClickhouseZookeeperConnectionIssues
expr: avg(ClickHouseMetrics_ZooKeeperSession) != 1
for: 3m
labels:
    severity: warning
annotations:
    summary: ClickHouse ZooKeeper Connection Issues (instance {{ $labels.instance }})
    description: |-
        ClickHouse is experiencing issues with ZooKeeper connections, which may affect cluster state and coordination.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseZookeeperConnectionIssues.md

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
