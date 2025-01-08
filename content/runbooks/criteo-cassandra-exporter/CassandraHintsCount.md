---
title: CassandraHintsCount
description: Troubleshooting for alert CassandraHintsCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraHintsCount

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cassandra hints count has changed on {{ $labels.instance }} some nodes may go down

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraHintsCount" %}}

{{% comment %}}

```yaml
alert: CassandraHintsCount
expr: changes(cassandra_stats{name="org:apache:cassandra:metrics:storage:totalhints:count"}[1m]) > 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra hints count (instance {{ $labels.instance }})
    description: |-
        Cassandra hints count has changed on {{ $labels.instance }} some nodes may go down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraHintsCount.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
