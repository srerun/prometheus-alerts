---
title: MongodbReplicationLag
description: Troubleshooting for alert MongodbReplicationLag
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationLag

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Mongodb replication lag is more than 10s

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationLag" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationLag
expr: avg(mongodb_replset_member_optime_date{state="PRIMARY"}) - avg(mongodb_replset_member_optime_date{state="SECONDARY"}) > 10
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication lag (instance {{ $labels.instance }})
    description: |-
        Mongodb replication lag is more than 10s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationLag.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
