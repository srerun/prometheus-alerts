---
title: CassandraClientRequestUnavailableWrite
description: Troubleshooting for alert CassandraClientRequestUnavailableWrite
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestUnavailableWrite

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Write failures have occurred because too many nodes are unavailable

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraClientRequestUnavailableWrite
expr: changes(cassandra_stats{name="org:apache:cassandra:metrics:clientrequest:write:unavailables:count"}[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra client request unavailable write (instance {{ $labels.instance }})
    description: |-
        Write failures have occurred because too many nodes are unavailable
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraClientRequestUnavailableWrite

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
