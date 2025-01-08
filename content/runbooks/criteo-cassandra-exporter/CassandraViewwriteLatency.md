---
title: CassandraViewwriteLatency
description: Troubleshooting for alert CassandraViewwriteLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraViewwriteLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High viewwrite latency on {{ $labels.instance }} cassandra node

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraViewwriteLatency" %}}

<!-- Rule when generated

```yaml
alert: CassandraViewwriteLatency
expr: cassandra_stats{name="org:apache:cassandra:metrics:clientrequest:viewwrite:viewwritelatency:99thpercentile",service="cas"} > 100000
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra viewwrite latency (instance {{ $labels.instance }})
    description: |-
        High viewwrite latency on {{ $labels.instance }} cassandra node
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraViewwriteLatency.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
