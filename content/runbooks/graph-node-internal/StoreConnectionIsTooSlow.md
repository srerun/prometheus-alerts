---
title: StoreConnectionIsTooSlow
description: Troubleshooting for alert StoreConnectionIsTooSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# StoreConnectionIsTooSlow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Store connection is too slow to `{{$labels.pool}}` pool, `{{$labels.shard}}` shard in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "StoreConnectionIsTooSlow" %}}

<!-- Rule when generated

```yaml
alert: StoreConnectionIsTooSlow
expr: store_connection_wait_time_ms > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Store connection is too slow (instance {{ $labels.instance }})
    description: |-
        Store connection is too slow to `{{$labels.pool}}` pool, `{{$labels.shard}}` shard in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/StoreConnectionIsTooSlow.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
