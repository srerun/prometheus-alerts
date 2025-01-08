---
title: CassandraConnectionTimeoutsTotal
description: Troubleshooting for alert CassandraConnectionTimeoutsTotal
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraConnectionTimeoutsTotal

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some connection between nodes are ending in timeout

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraConnectionTimeoutsTotal" %}}

<!-- Rule when generated

```yaml
alert: CassandraConnectionTimeoutsTotal
expr: rate(cassandra_stats{name="org:apache:cassandra:metrics:connection:totaltimeouts:count"}[1m]) > 5
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra connection timeouts total (instance {{ $labels.instance }})
    description: |-
        Some connection between nodes are ending in timeout
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraConnectionTimeoutsTotal.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
