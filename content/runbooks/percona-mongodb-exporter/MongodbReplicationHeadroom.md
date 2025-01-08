---
title: MongodbReplicationHeadroom
description: Troubleshooting for alert MongodbReplicationHeadroom
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationHeadroom

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MongoDB replication headroom is <= 0

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbReplicationHeadroom" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationHeadroom
expr: sum(avg(mongodb_mongod_replset_oplog_head_timestamp - mongodb_mongod_replset_oplog_tail_timestamp)) - sum(avg(mongodb_rs_members_optimeDate{member_state="PRIMARY"} - on (set) group_right mongodb_rs_members_optimeDate{member_state="SECONDARY"})) <= 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication headroom (instance {{ $labels.instance }})
    description: |-
        MongoDB replication headroom is <= 0
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbReplicationHeadroom.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
