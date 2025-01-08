---
title: MongodbReplicationStatus10
description: Troubleshooting for alert MongodbReplicationStatus10
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus10

MongoDB Replication set member was once in a replica set but was subsequently removed

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus10" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus10
expr: mongodb_replset_member_state == 10
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 10 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member was once in a replica set but was subsequently removed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus10.md

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
