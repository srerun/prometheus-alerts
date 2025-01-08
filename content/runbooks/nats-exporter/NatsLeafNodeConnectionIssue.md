---
title: NatsLeafNodeConnectionIssue
description: Troubleshooting for alert NatsLeafNodeConnectionIssue
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsLeafNodeConnectionIssue

## Meaning
[//]: # "Short paragraph that explains what the alert means"
No leaf node connections have been established in the last 5 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsLeafNodeConnectionIssue" %}}

{{% comment %}}

```yaml
alert: NatsLeafNodeConnectionIssue
expr: increase(gnatsd_varz_leafnodes[5m]) == 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Nats leaf node connection issue (instance {{ $labels.instance }})
    description: |-
        No leaf node connections have been established in the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsLeafNodeConnectionIssue.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
