---
title: MongodbReplicationStatus9
description: Troubleshooting for alert MongodbReplicationStatus9
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus9

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MongoDB Replication set member is actively performing a rollback. Data is not available for reads

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus9" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus9
expr: mongodb_replset_member_state == 9
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 9 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member is actively performing a rollback. Data is not available for reads
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus9.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
