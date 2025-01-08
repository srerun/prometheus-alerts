---
title: HostNetworkBondDegraded
description: Troubleshooting for alert HostNetworkBondDegraded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkBondDegraded

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Bond "{{ $labels.device }}" degraded on "{{ $labels.instance }}".

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkBondDegraded" %}}

<!-- Rule when generated

```yaml
alert: HostNetworkBondDegraded
expr: ((node_bonding_active - node_bonding_slaves) != 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host Network Bond Degraded (instance {{ $labels.instance }})
    description: |-
        Bond "{{ $labels.device }}" degraded on "{{ $labels.instance }}".
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkBondDegraded.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
