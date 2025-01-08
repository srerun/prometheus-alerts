---
title: MongodbCursorsTimeouts
description: Troubleshooting for alert MongodbCursorsTimeouts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbCursorsTimeouts

Too many cursors are timing out

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbCursorsTimeouts" %}}

{{% comment %}}

```yaml
alert: MongodbCursorsTimeouts
expr: increase(mongodb_metrics_cursor_timed_out_total[1m]) > 100
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB cursors timeouts (instance {{ $labels.instance }})
    description: |-
        Too many cursors are timing out
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbCursorsTimeouts.md

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
