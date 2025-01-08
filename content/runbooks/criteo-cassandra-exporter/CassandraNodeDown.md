---
title: CassandraNodeDown
description: Troubleshooting for alert CassandraNodeDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraNodeDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cassandra node down

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraNodeDown" %}}

{{% comment %}}

```yaml
alert: CassandraNodeDown
expr: sum(cassandra_stats{name="org:apache:cassandra:net:failuredetector:downendpointcount"}) by (service,group,cluster,env) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra node down (instance {{ $labels.instance }})
    description: |-
        Cassandra node down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraNodeDown.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
