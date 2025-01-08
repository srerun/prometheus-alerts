---
title: ProviderFailedBecauseGetGenesisTimeout
description: Troubleshooting for alert ProviderFailedBecauseGetGenesisTimeout
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseGetGenesisTimeout

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Timeout to get genesis for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseGetGenesisTimeout" %}}

<!-- Rule when generated

```yaml
alert: ProviderFailedBecauseGetGenesisTimeout
expr: eth_rpc_status == 4
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because get genesis timeout (instance {{ $labels.instance }})
    description: |-
        Timeout to get genesis for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseGetGenesisTimeout.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
