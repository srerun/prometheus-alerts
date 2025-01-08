---
title: MongodbReplicationStatus6
description: Troubleshooting for alert MongodbReplicationStatus6
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus6

MongoDB Replication set member as seen from another member of the set, is not yet known

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus6" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus6
expr: mongodb_replset_member_state == 6
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 6 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member as seen from another member of the set, is not yet known
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus6.md

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
