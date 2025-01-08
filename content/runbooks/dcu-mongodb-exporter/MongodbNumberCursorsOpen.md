---
title: MongodbNumberCursorsOpen
description: Troubleshooting for alert MongodbNumberCursorsOpen
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbNumberCursorsOpen

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many cursors opened by MongoDB for clients (> 10k)

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbNumberCursorsOpen" %}}

{{% comment %}}

```yaml
alert: MongodbNumberCursorsOpen
expr: mongodb_metrics_cursor_open{state="total_open"} > 10000
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB number cursors open (instance {{ $labels.instance }})
    description: |-
        Too many cursors opened by MongoDB for clients (> 10k)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbNumberCursorsOpen.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
