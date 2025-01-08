---
title: PostgresqlUnusedReplicationSlot
description: Troubleshooting for alert PostgresqlUnusedReplicationSlot
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlUnusedReplicationSlot

Unused Replication Slots

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlUnusedReplicationSlot" %}}

{{% comment %}}

```yaml
alert: PostgresqlUnusedReplicationSlot
expr: pg_replication_slots_active == 0
for: 1m
labels:
    severity: warning
annotations:
    summary: Postgresql unused replication slot (instance {{ $labels.instance }})
    description: |-
        Unused Replication Slots
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlUnusedReplicationSlot.md

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
