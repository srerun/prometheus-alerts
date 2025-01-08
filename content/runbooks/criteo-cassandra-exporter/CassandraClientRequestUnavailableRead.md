---
title: CassandraClientRequestUnavailableRead
description: Troubleshooting for alert CassandraClientRequestUnavailableRead
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestUnavailableRead

Read failures have occurred because too many nodes are unavailable

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraClientRequestUnavailableRead" %}}

{{% comment %}}

```yaml
alert: CassandraClientRequestUnavailableRead
expr: changes(cassandra_stats{name="org:apache:cassandra:metrics:clientrequest:read:unavailables:count"}[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra client request unavailable read (instance {{ $labels.instance }})
    description: |-
        Read failures have occurred because too many nodes are unavailable
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraClientRequestUnavailableRead.md

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
