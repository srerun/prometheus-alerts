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

Too much tombstones scanned in queries

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraTombstoneDump" %}}

{{% comment %}}

```yaml
alert: CassandraTombstoneDump
expr: cassandra_stats{name="org:apache:cassandra:metrics:table:tombstonescannedhistogram:99thpercentile"} > 1000
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra tombstone dump (instance {{ $labels.instance }})
    description: |-
        Too much tombstones scanned in queries
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraTombstoneDump.md

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
