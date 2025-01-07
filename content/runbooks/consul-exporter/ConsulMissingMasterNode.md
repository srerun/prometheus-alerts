---
title: ConsulMissingMasterNode
description: Troubleshooting for alert ConsulMissingMasterNode
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ConsulMissingMasterNode

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Numbers of consul raft peers should be 3, in order to preserve quorum.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ConsulMissingMasterNode
expr: consul_raft_peers < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Consul missing master node (instance {{ $labels.instance }})
    description: |-
        Numbers of consul raft peers should be 3, in order to preserve quorum.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ConsulMissingMasterNode

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
