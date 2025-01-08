---
title: CassandraNodeIsUnavailable
description: Troubleshooting for alert CassandraNodeIsUnavailable
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraNodeIsUnavailable

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cassandra Node is unavailable - {{ $labels.cassandra_cluster }} {{ $labels.exported_endpoint }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraNodeIsUnavailable" %}}

<!-- Rule when generated

```yaml
alert: CassandraNodeIsUnavailable
expr: sum(cassandra_endpoint_active) by (cassandra_cluster,instance,exported_endpoint) < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra Node is unavailable (instance {{ $labels.instance }})
    description: |-
        Cassandra Node is unavailable - {{ $labels.cassandra_cluster }} {{ $labels.exported_endpoint }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraNodeIsUnavailable.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
