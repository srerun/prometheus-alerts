---
title: MongodbTooManyConnections
description: Troubleshooting for alert MongodbTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbTooManyConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many connections (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbTooManyConnections" %}}

<!-- Rule when generated

```yaml
alert: MongodbTooManyConnections
expr: avg by(instance) (rate(mongodb_ss_connections{conn_type="current"}[1m])) / avg by(instance) (sum (mongodb_ss_connections) by (instance)) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB too many connections (instance {{ $labels.instance }})
    description: |-
        Too many connections (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbTooManyConnections.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
