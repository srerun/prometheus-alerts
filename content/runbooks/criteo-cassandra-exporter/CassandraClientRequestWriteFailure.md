---
title: CassandraClientRequestWriteFailure
description: Troubleshooting for alert CassandraClientRequestWriteFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestWriteFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A lot of write failures encountered. A write failure is a non-timeout exception encountered during a write request. Examine the reason map to find to the root cause. The most common cause for this type of error is when batch sizes are too large.

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraClientRequestWriteFailure" %}}

<!-- Rule when generated

```yaml
alert: CassandraClientRequestWriteFailure
expr: increase(cassandra_stats{name="org:apache:cassandra:metrics:clientrequest:write:failures:oneminuterate"}[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra client request write failure (instance {{ $labels.instance }})
    description: |-
        A lot of write failures encountered. A write failure is a non-timeout exception encountered during a write request. Examine the reason map to find to the root cause. The most common cause for this type of error is when batch sizes are too large.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraClientRequestWriteFailure.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
