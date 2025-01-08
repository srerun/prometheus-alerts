---
title: ProviderFailedBecauseGetGenesisFailed
description: Troubleshooting for alert ProviderFailedBecauseGetGenesisFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseGetGenesisFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Failed to get genesis for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseGetGenesisFailed" %}}

<!-- Rule when generated

```yaml
alert: ProviderFailedBecauseGetGenesisFailed
expr: eth_rpc_status == 2
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because get genesis failed (instance {{ $labels.instance }})
    description: |-
        Failed to get genesis for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseGetGenesisFailed.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
