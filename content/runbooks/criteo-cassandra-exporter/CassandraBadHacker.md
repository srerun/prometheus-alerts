---
title: CassandraBadHacker
description: Troubleshooting for alert CassandraBadHacker
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraBadHacker

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Increase of Cassandra authentication failures

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraBadHacker" %}}

{{% comment %}}

```yaml
alert: CassandraBadHacker
expr: rate(cassandra_stats{name="org:apache:cassandra:metrics:client:authfailure:count"}[1m]) > 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra bad hacker (instance {{ $labels.instance }})
    description: |-
        Increase of Cassandra authentication failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraBadHacker.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
