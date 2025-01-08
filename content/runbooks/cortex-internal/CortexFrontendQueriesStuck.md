---
title: CortexFrontendQueriesStuck
description: Troubleshooting for alert CortexFrontendQueriesStuck
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexFrontendQueriesStuck

There are queued up queries in query-frontend.

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/cortex-internal.yml" "CortexFrontendQueriesStuck" %}}

{{% comment %}}

```yaml
alert: CortexFrontendQueriesStuck
expr: sum by (job) (cortex_query_frontend_queue_length) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Cortex frontend queries stuck (instance {{ $labels.instance }})
    description: |-
        There are queued up queries in query-frontend.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexFrontendQueriesStuck.md

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
