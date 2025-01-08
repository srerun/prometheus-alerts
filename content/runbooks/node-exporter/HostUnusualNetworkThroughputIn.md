---
title: HostUnusualNetworkThroughputIn
description: Troubleshooting for alert HostUnusualNetworkThroughputIn
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualNetworkThroughputIn

Host network interfaces are probably receiving too much data (> 100 MB/s)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualNetworkThroughputIn" %}}

{{% comment %}}

```yaml
alert: HostUnusualNetworkThroughputIn
expr: (sum by (instance) (rate(node_network_receive_bytes_total[2m])) / 1024 / 1024 > 100) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host unusual network throughput in (instance {{ $labels.instance }})
    description: |-
        Host network interfaces are probably receiving too much data (> 100 MB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualNetworkThroughputIn.md

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
