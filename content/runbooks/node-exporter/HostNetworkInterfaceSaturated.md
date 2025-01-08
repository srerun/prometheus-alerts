---
title: HostNetworkInterfaceSaturated
description: Troubleshooting for alert HostNetworkInterfaceSaturated
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkInterfaceSaturated

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The network interface "{{ $labels.device }}" on "{{ $labels.instance }}" is getting overloaded.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkInterfaceSaturated" %}}

<!-- Rule when generated

```yaml
alert: HostNetworkInterfaceSaturated
expr: ((rate(node_network_receive_bytes_total{device!~"^tap.*|^vnet.*|^veth.*|^tun.*"}[1m]) + rate(node_network_transmit_bytes_total{device!~"^tap.*|^vnet.*|^veth.*|^tun.*"}[1m])) / node_network_speed_bytes{device!~"^tap.*|^vnet.*|^veth.*|^tun.*"} > 0.8 < 10000) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 1m
labels:
    severity: warning
annotations:
    summary: Host Network Interface Saturated (instance {{ $labels.instance }})
    description: |-
        The network interface "{{ $labels.device }}" on "{{ $labels.instance }}" is getting overloaded.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkInterfaceSaturated.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
