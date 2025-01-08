---
title: MongodbReplicaMemberUnhealthy
description: Troubleshooting for alert MongodbReplicaMemberUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicaMemberUnhealthy

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MongoDB replica member is not healthy

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbReplicaMemberUnhealthy" %}}

<!-- Rule when generated

```yaml
alert: MongodbReplicaMemberUnhealthy
expr: mongodb_rs_members_health == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Mongodb replica member unhealthy (instance {{ $labels.instance }})
    description: |-
        MongoDB replica member is not healthy
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbReplicaMemberUnhealthy.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
