---
title: CassandraConnectionTimeoutsTotal
description: Troubleshooting for alert CassandraConnectionTimeoutsTotal
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraConnectionTimeoutsTotal

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some connection between nodes are ending in timeout - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraConnectionTimeoutsTotal
expr: avg(cassandra_client_request_timeouts_total) by (cassandra_cluster,instance) > 5
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra connection timeouts total (instance {{ $labels.instance }})
    description: |-
        Some connection between nodes are ending in timeout - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraConnectionTimeoutsTotal

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
