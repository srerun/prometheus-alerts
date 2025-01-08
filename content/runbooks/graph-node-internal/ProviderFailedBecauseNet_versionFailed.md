---
title: ProviderFailedBecauseNet_versionFailed
description: Troubleshooting for alert ProviderFailedBecauseNet_versionFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseNet_versionFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Failed net_version for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseNet_versionFailed" %}}

<!-- Rule when generated

```yaml
alert: ProviderFailedBecauseNet_versionFailed
expr: eth_rpc_status == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because net_version failed (instance {{ $labels.instance }})
    description: |-
        Failed net_version for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseNet_versionFailed.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
