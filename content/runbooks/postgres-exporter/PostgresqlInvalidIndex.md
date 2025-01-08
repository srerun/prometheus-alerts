---
title: PostgresqlInvalidIndex
description: Troubleshooting for alert PostgresqlInvalidIndex
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlInvalidIndex

The table {{ $labels.relname }} has an invalid index: {{ $labels.indexrelname }}. You should execute `DROP INDEX {{ $labels.indexrelname }};`

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlInvalidIndex" %}}

{{% comment %}}

```yaml
alert: PostgresqlInvalidIndex
expr: pg_general_index_info_pg_relation_size{indexrelname=~".*ccnew.*"}
for: 6h
labels:
    severity: warning
annotations:
    summary: Postgresql invalid index (instance {{ $labels.instance }})
    description: |-
        The table {{ $labels.relname }} has an invalid index: {{ $labels.indexrelname }}. You should execute `DROP INDEX {{ $labels.indexrelname }};`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlInvalidIndex.md

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
