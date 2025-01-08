---
title: CassandraTombstoneDump
description: Troubleshooting for alert CassandraTombstoneDump
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraTombstoneDump

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cassandra tombstone dump - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraTombstoneDump" %}}

{{% comment %}}

```yaml
alert: CassandraTombstoneDump
expr: avg(cassandra_table_tombstones_scanned{quantile="0.99"}) by (instance,cassandra_cluster,keyspace) > 100
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra tombstone dump (instance {{ $labels.instance }})
    description: |-
        Cassandra tombstone dump - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraTombstoneDump.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
