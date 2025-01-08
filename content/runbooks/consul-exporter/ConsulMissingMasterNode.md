---
title: ConsulMissingMasterNode
description: Troubleshooting for alert ConsulMissingMasterNode
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ConsulMissingMasterNode

Numbers of consul raft peers should be 3, in order to preserve quorum.

<details>
  <summary>Alert Rule</summary>

{{% rule "consul/consul-exporter.yml" "ConsulMissingMasterNode" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/consul-exporter/ConsulMissingMasterNode.md

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
