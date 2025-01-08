---
title: CassandraCacheHitRateKeyCache
description: Troubleshooting for alert CassandraCacheHitRateKeyCache
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraCacheHitRateKeyCache

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Key cache hit rate is below 85%

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraCacheHitRateKeyCache" %}}

{{% comment %}}

```yaml
alert: CassandraCacheHitRateKeyCache
expr: cassandra_stats{name="org:apache:cassandra:metrics:cache:keycache:hitrate:value"} < .85
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra cache hit rate key cache (instance {{ $labels.instance }})
    description: |-
        Key cache hit rate is below 85%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraCacheHitRateKeyCache.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
