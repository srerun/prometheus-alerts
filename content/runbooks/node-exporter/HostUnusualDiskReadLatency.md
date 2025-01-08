---
title: HostUnusualDiskReadLatency
description: Troubleshooting for alert HostUnusualDiskReadLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskReadLatency

Disk latency is growing (read operations > 100ms)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualDiskReadLatency" %}}

{{% comment %}}

```yaml
alert: HostUnusualDiskReadLatency
expr: (rate(node_disk_read_time_seconds_total[1m]) / rate(node_disk_reads_completed_total[1m]) > 0.1 and rate(node_disk_reads_completed_total[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host unusual disk read latency (instance {{ $labels.instance }})
    description: |-
        Disk latency is growing (read operations > 100ms)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualDiskReadLatency.md

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
