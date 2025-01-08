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
Read failures have occurred, ensure there are not too many unavailable nodes - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraClientRequestWriteFailure" %}}

{{% comment %}}

```yaml
alert: CassandraClientRequestWriteFailure
expr: increase(cassandra_client_request_failures_total{operation="write"}[1m]) > 0
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra client request write failure (instance {{ $labels.instance }})
    description: |-
        Read failures have occurred, ensure there are not too many unavailable nodes - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraClientRequestWriteFailure.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
