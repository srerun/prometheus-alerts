---
title: CassandraClientRequestReadFailure
description: Troubleshooting for alert CassandraClientRequestReadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestReadFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Read failures have occurred, ensure there are not too many unavailable nodes - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraClientRequestReadFailure
expr: increase(cassandra_client_request_failures_total{operation="read"}[1m]) > 0
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra client request read failure (instance {{ $labels.instance }})
    description: |-
        Read failures have occurred, ensure there are not too many unavailable nodes - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraClientRequestReadFailure

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
