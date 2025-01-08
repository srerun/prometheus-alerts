---
title: HostNetworkTransmitErrors
description: Troubleshooting for alert HostNetworkTransmitErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkTransmitErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Host {{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} transmit errors in the last two minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkTransmitErrors" %}}

<!-- Rule when generated

```yaml
alert: HostNetworkTransmitErrors
expr: (rate(node_network_transmit_errs_total[2m]) / rate(node_network_transmit_packets_total[2m]) > 0.01) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host Network Transmit Errors (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} transmit errors in the last two minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkTransmitErrors.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
