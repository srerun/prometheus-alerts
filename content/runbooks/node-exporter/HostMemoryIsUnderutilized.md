---
title: HostMemoryIsUnderutilized
description: Troubleshooting for alert HostMemoryIsUnderutilized
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostMemoryIsUnderutilized

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Node memory is < 20% for 1 week. Consider reducing memory space. (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostMemoryIsUnderutilized" %}}

{{% comment %}}

```yaml
alert: HostMemoryIsUnderutilized
expr: (100 - (avg_over_time(node_memory_MemAvailable_bytes[30m]) / node_memory_MemTotal_bytes * 100) < 20) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 1w
labels:
    severity: info
annotations:
    summary: Host Memory is underutilized (instance {{ $labels.instance }})
    description: |-
        Node memory is < 20% for 1 week. Consider reducing memory space. (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostMemoryIsUnderutilized.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
