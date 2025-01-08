---
title: ProviderFailedBecauseNet_versionTimeout
description: Troubleshooting for alert ProviderFailedBecauseNet_versionTimeout
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseNet_versionTimeout

net_version timeout for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseNet_versionTimeout" %}}

{{% comment %}}

```yaml
alert: ProviderFailedBecauseNet_versionTimeout
expr: eth_rpc_status == 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because net_version timeout (instance {{ $labels.instance }})
    description: |-
        net_version timeout for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseNet_versionTimeout.md

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
